package main

const MODELS_LIST = `
{
  "data": [
    {
      "version": "gpt-3.5-turbo-0613",
      "capabilities": {
        "supports": {
          "tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 12288,
          "max_context_window_tokens": 16384,
          "max_output_tokens": 4096
        },
        "tokenizer": "cl100k_base",
        "type": "chat",
        "family": "gpt-3.5-turbo"
      },
      "id": "gpt-3.5-turbo",
      "object": "model",
      "name": "GPT 3.5 Turbo",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-3.5-turbo-0613",
      "capabilities": {
        "supports": {
          "tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 12288,
          "max_context_window_tokens": 16384,
          "max_output_tokens": 4096
        },
        "tokenizer": "cl100k_base",
        "type": "chat",
        "family": "gpt-3.5-turbo"
      },
      "id": "gpt-3.5-turbo-0613",
      "object": "model",
      "name": "GPT 3.5 Turbo",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-4-0613",
      "capabilities": {
        "supports": {
          "tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 32768,
          "max_context_window_tokens": 32768,
          "max_output_tokens": 4096
        },
        "tokenizer": "cl100k_base",
        "type": "chat",
        "family": "gpt-4"
      },
      "id": "gpt-4",
      "object": "model",
      "name": "GPT 4",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-4-0613",
      "capabilities": {
        "supports": {
          "tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 32768,
          "max_context_window_tokens": 32768,
          "max_output_tokens": 4096
        },
        "tokenizer": "cl100k_base",
        "type": "chat",
        "family": "gpt-4"
      },
      "id": "gpt-4-0613",
      "object": "model",
      "name": "GPT 4",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-4o-2024-05-13",
      "capabilities": {
        "supports": {
          "tool_calls": true,
          "parallel_tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 64000,
          "max_context_window_tokens": 128000,
          "max_output_tokens": 4096
        },
        "tokenizer": "o200k_base",
        "type": "chat",
        "family": "gpt-4o"
      },
      "id": "gpt-4o",
      "object": "model",
      "name": "GPT 4o",
      "model_picker_enabled": true
    },
    {
      "version": "gpt-4o-2024-05-13",
      "capabilities": {
        "supports": {
          "tool_calls": true,
          "parallel_tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 64000,
          "max_context_window_tokens": 128000,
          "max_output_tokens": 4096
        },
        "tokenizer": "o200k_base",
        "type": "chat",
        "family": "gpt-4o"
      },
      "id": "gpt-4o-2024-05-13",
      "object": "model",
      "name": "GPT 4o",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-4o-2024-05-13",
      "capabilities": {
        "supports": {
          "tool_calls": true,
          "parallel_tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 64000,
          "max_context_window_tokens": 128000,
          "max_output_tokens": 4096
        },
        "tokenizer": "o200k_base",
        "type": "chat",
        "family": "gpt-4o"
      },
      "id": "gpt-4-o-preview",
      "object": "model",
      "name": "GPT 4o",
      "model_picker_enabled": false
    },
    {
      "version": "text-embedding-ada-002",
      "capabilities": {
        "supports": {},
        "object": "model_capabilities",
        "limits": {
          "max_inputs": 256
        },
        "tokenizer": "cl100k_base",
        "type": "embeddings",
        "family": "text-embedding-ada-002"
      },
      "id": "text-embedding-ada-002",
      "object": "model",
      "name": "Embedding V2 Ada",
      "model_picker_enabled": false
    },
    {
      "version": "text-embedding-3-small",
      "capabilities": {
        "supports": {
          "dimensions": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_inputs": 512
        },
        "tokenizer": "cl100k_base",
        "type": "embeddings",
        "family": "text-embedding-3-small"
      },
      "id": "text-embedding-3-small",
      "object": "model",
      "name": "Embedding V3 small",
      "model_picker_enabled": false
    },
    {
      "version": "text-embedding-3-small",
      "capabilities": {
        "supports": {
          "dimensions": true
        },
        "object": "model_capabilities",
        "tokenizer": "cl100k_base",
        "type": "embeddings",
        "family": "text-embedding-3-small"
      },
      "id": "text-embedding-3-small-inference",
      "object": "model",
      "name": "Embedding V3 small (Inference)",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-4o-mini-2024-07-18",
      "capabilities": {
        "supports": {
          "tool_calls": true,
          "parallel_tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 12288,
          "max_context_window_tokens": 128000,
          "max_output_tokens": 4096
        },
        "tokenizer": "o200k_base",
        "type": "chat",
        "family": "gpt-4o-mini"
      },
      "id": "gpt-4o-mini",
      "object": "model",
      "name": "GPT 4o Mini",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-4o-mini-2024-07-18",
      "capabilities": {
        "supports": {
          "tool_calls": true,
          "parallel_tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 12288,
          "max_context_window_tokens": 128000,
          "max_output_tokens": 4096
        },
        "tokenizer": "o200k_base",
        "type": "chat",
        "family": "gpt-4o-mini"
      },
      "id": "gpt-4o-mini-2024-07-18",
      "object": "model",
      "name": "GPT 4o Mini",
      "model_picker_enabled": false
    },
    {
      "version": "gpt-4-0125-preview",
      "capabilities": {
        "supports": {
          "tool_calls": true,
          "parallel_tool_calls": true
        },
        "object": "model_capabilities",
        "limits": {
          "max_prompt_tokens": 64000,
          "max_context_window_tokens": 128000,
          "max_output_tokens": 4096
        },
        "tokenizer": "cl100k_base",
        "type": "chat",
        "family": "gpt-4-turbo"
      },
      "id": "gpt-4-0125-preview",
      "object": "model",
      "name": "GPT 4 Turbo",
      "model_picker_enabled": false
    }
  ],
  "object": "list"
}
`
