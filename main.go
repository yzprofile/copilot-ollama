package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

var LISTEN_ADDR = ":11435"
var OLLAMA_HOST = "http://localhost:11434"
var MODEL = "deepseek-coder-v2"
var VERBOSE = false

func flags() {
	flag.StringVar(&LISTEN_ADDR, "listen_addr", ":11435", "The address to listen on for HTTP requests.")
	flag.StringVar(&OLLAMA_HOST, "ollama_host", "http://localhost:11434", "ollama host")
	flag.StringVar(&MODEL, "model", "deepseek-coder-v2", "ollama model")
	flag.BoolVar(&VERBOSE, "verbose", false, "Enable debug mode.")
	flag.Parse()
}

func ping(w http.ResponseWriter, req *http.Request) {
	buf, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, string(buf))
	}
}

func models(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(MODELS_LIST))
}

func copilot_completion(w http.ResponseWriter, req *http.Request) {
	req.URL.Path = "/v1/completions"
	req.URL.Host = OLLAMA_HOST

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	data["model"] = MODEL
	nbody, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(nbody))
	proxy(w, req)
}

func chat_completion(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	data["model"] = MODEL
	nbody, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	req.Body = io.NopCloser(bytes.NewReader(nbody))
	proxy(w, req)
}

func proxy(w http.ResponseWriter, req *http.Request) {

	dump_req("REQ-IN:", req)

	newurl := OLLAMA_HOST + req.URL.Path
	nreq, err := http.NewRequest(req.Method, newurl, req.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	for name, values := range req.Header {
		for _, value := range values {
			nreq.Header.Add(name, value)
		}
	}

	dump_req("REQ-OUT:", nreq)

	rsp, err := http.DefaultClient.Do(nreq)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	defer rsp.Body.Close()

	// copy header
	for name, values := range rsp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(rsp.StatusCode)
	_, err = io.Copy(w, rsp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	dump_rsp("RSP-IN:", rsp)
}

func main() {
	flags()

	fmt.Println("model:", MODEL, "listen on:", LISTEN_ADDR, "proxy to:", OLLAMA_HOST)

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/v1/engines/copilot-codex/completions", copilot_completion)
	http.HandleFunc("/v1/chat/completions", chat_completion)
	http.HandleFunc("/v1/models", models)
	http.HandleFunc("/", proxy)

	http.ListenAndServe(LISTEN_ADDR, nil)
}
