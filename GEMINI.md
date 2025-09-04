### Project Overview

This project is a simple command-line tool written in Go. Its purpose is to display local network interface information, specifically the IPv4 address and subnet mask for each interface. The output is color-coded for readability using the `lipgloss` library. The project is cross-platform and supports Linux, Windows, and macOS.

### Building and Running

**Building the project:**

You can build the project using Go or `just`:

*   **Using Go:**
    ```sh
    go build .
    ```

*   **Using just:**
    ```sh
    just build
    ```

**Running the project:**

After building, you can run the executable from your terminal:

```sh
./ipl
```

You can also use the `run` command with `just`:

```sh
just run
```

### Development Conventions

The code is formatted according to Go's standard formatting conventions. The project uses a `justfile` for task automation, which is a common convention in the Go ecosystem for simple projects. The project has a single source file (`main.go`) which is appropriate for a small command-line tool. Dependencies are managed using Go modules.

The project also includes a `flake.nix` file, which defines a development environment using Nix. This ensures a consistent development environment with all the necessary tools, including:

*   **Go:** The Go compiler and tools.
*   **gopls:** The Go language server.
*   **delve:** The Go debugger.
*   **Formatters:** `treefmt2`, `mdformat`, and `alejandra` for code formatting.
*   **Task Runners:** `just` and `watchexec`.

The project has a CI pipeline configured in `.github/workflows/go.yml` that builds and tests the project on every push and pull request to the `main` branch.

The project uses an `.editorconfig` file to define and maintain consistent coding styles across different editors and IDEs. It specifies settings like indentation style, character set, and whitespace handling.
