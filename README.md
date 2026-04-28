# ipl

[![Go Workflow](https://github.com/qrxnz/ipl/actions/workflows/go.yml/badge.svg)](https://github.com/qrxnz/ipl/actions/workflows/go.yml)

![Screenshot of ipl output](https://github.com/user-attachments/assets/9fb46390-2187-4305-8446-8237fca8ed61)

> A simple command-line tool for displaying local network interface information

`ipl` quickly lists your network interfaces, along with their IPv4 addresses and subnet masks, in a clean, color-coded output.

## 🧰 Features

-   **Clear Output**: Displays interface name, IPv4 address, and mask in a single, readable line.
-   **Color-Coded**: Uses colors to distinguish between different parts of the output for better readability.
-   **Cross-Platform**: Works on Linux, Windows, and macOS.

## 🛠️ Installation

### 📦 Binary Releases

Pre-compiled binaries for Linux, Windows, and macOS are available on the [Releases](https://github.com/qrxnz/ipl/releases) page.

### 🐹Using Go

You can install `ipl` directly using `go install`:

```bash
go install github.com/qrxnz/ipl@latest
```

### 🏗️ Build from Source

To build from source, you need to have [Go](https://go.dev/) installed.

```bash
git clone https://github.com/qrxnz/ipl.git
cd ipl
go build -o ipl .
```

Alternatively, if you have [Task](https://taskfile.dev/) installed, you can use:

```bash
task build
```

### ❄️ Using Nix

-   **Run without installing:**

```bash
nix run github:qrxnz/ipl
```

-   **Add to a Nix Flake:**

Add input in your flake like

```nix
{
 inputs = {
   ipl = {
     url = "github:qrxnz/ipl";
     inputs.nixpkgs.follows = "nixpkgs";
   };
 };
}
```

With the input added you can reference it directly

```nix
{ inputs, system, ... }:
{
  # NixOS
  environment.systemPackages = [ inputs.ipl.packages.${pkgs.system}.default ];
  # home-manager
  home.packages = [ inputs.ipl.packages.${pkgs.system}.default ];
}
```

-   **Install imperatively:**

```bash
nix profile install github:qrxnz/ipl
```

## 📖 Usage

Run the compiled binary from your terminal:

```sh
ipl
```

## 📜 License

This project is licensed under the [MIT License](LICENSE).
