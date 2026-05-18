# Pasty

Pasty is a command-line clipboard editor that allows you to add macros to any application. It's written in Go and can be extended with Lua plugins.

## Project Overview

*   **Purpose:** To provide a quick and easy way to manipulate text from the clipboard using a variety of commands and custom scripts.
*   **Technologies:** Go, Cobra (for CLI), gopher-lua (for plugin support).
*   **Architecture:** The project is structured as a typical Go application. The main entry point is `main.go`, which initializes the `cobra` CLI application. The core commands are defined in `cmd/root.go`, and utility functions are located in `util/util.go`. The plugin handling logic is in `pluginHandler.go`.

## Building and Running

To build the project, you can use the standard `go build` command:

```bash
go build
```

This will create an executable named `pasty` in the current directory.

To run the application, you can use the following command:

```bash
./pasty [command]
```

### Commands

*   `rep`: Replaces text.
*   `case`: Changes the case of text.

## Development Conventions

The project follows standard Go conventions. It uses `cobra` for the CLI, which provides a structured way to define commands and flags. The use of a `util` package for helper functions is also a common practice.

### Plugins

Pasty can be extended with Lua plugins. Plugins are loaded from `~/.config/pasty/plugins/plugin.lua`. The plugin must define a function called `process` that takes the input text and a table of parameters as arguments and returns the modified text.

Here is an example of a simple `plugin.lua` file:

```lua
function process(text, params)
  return text:upper()
end
```

This plugin will convert the input text to uppercase.
