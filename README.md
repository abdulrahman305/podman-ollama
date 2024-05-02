# podman-ollama

The goal of podman-ollama is to make AI even more boring.

[podman](https://github.com/containers/podman) pulls the runtime environment. [ollama](https://github.com/ollama/ollama) pulls the model library.

## Install

Install podman-ollama by running this one-liner:

```bash
curl -fsSL https://raw.githubusercontent.com/ericcurtin/podman-ollama/s/install.sh | sudo bash
```

## Usage

If podman-ollama is passed text it is a non-interactive tool:

```bash
$ podman-ollama "Write a git commit message for this diff: $(git diff)"
`Fixed formatting in README.md`

This commit message provides a brief description of the changes made in
the file, and is written in a format that is easy to understand and use.
```

If podman-ollama is run without any arguments it is an interactive tool:

``` bash
$ podman-ollama
>>> Tell me about podman in less than ten words
Podman: Containerized application management.

>>> Send a message (/? for help)
```

Full usage help:

```
$ podman-ollama -h
The goal of podman-ollama is to make AI even more boring.

Usage:
  podman-ollama [prompt]
  podman-ollama [options]
  podman-ollama [command]

Commands:
  serve       Start ollama server (not required)
  create      Create a model from a Modelfile
  chatbot     Set up chatbot UI interface
  open-webui  Set up open-webui UI interface
  show        Show information for a model
  run         Run a model, default if no command is specified
  pull        Pull a model from a registry
  push        Push a model to a registry
  list        List models
  cp          Copy a model
  rm          Remove a model
  help        Help about any command
  generate    Generate structured data based on containers, pods or volumes

Options:
  -c, --container-manager CONMAN - Specify podman or docker, default: podman
  -g, --gpu GPU                  - Specify a GPU: AMD, NVIDIA, GPU or CPU
  -h, --help                     - Usage help
  -l, --log LOGFILE              - Specify logfile to redirect to, for GPU debug
  -m, --model MODEL              - Specify non-default model, default: mistral
  --privileged                   - Give extended privileges to container
  -p, --publish                  - Publish a container's port to the host
  -r, --root                     - Run as a rootful container
  -v, --version                  - Show version information
  -                              - Read from stdin

Environment Variables:
  OLLAMA_HOST  The host:port or base URL of the Ollama server
               (e.g. http://some_remotehost:11434)

Configuration:
  podman-ollama uses a simple text format to store customizations that are per
  user in "~/.podman-ollama/config". Such a configuration file may look like
  this:

    container-manager podman
    gpu GPU
    model gemma:2b
```

