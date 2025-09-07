# ipl

[![Go Workflow](https://github.com/qrxnz/ipl/actions/workflows/go.yml/badge.svg)](https://github.com/qrxnz/ipl/actions/workflows/go.yml)

## ✒️ Description

> A simple command-line tool for displaying local network interface information

<p align="center">
  <a href="https://go-skill-icons.vercel.app/">
    <img src="https://go-skill-icons.vercel.app/api/icons?i=linux,windows,apple" alt="Supported OS: Linux, Windows, Apple" />
  </a>
</p>

`ipl` quickly lists your network interfaces, along with their IPv4 addresses and subnet masks, in a clean, color-coded output.

![Screenshot of ipl output](https://github.com/user-attachments/assets/9fb46390-2187-4305-8446-8237fca8ed61)

## 🧰 Features

- **Clear Output**: Displays interface name, IPv4 address, and mask in a single, readable line.
- **Color-Coded**: Uses colors to distinguish between different parts of the output for better readability.
- **Cross-Platform**: Works on Linux, Windows, and macOS.

## ⚒️ Installation

To build the application from the source, you need to have Go installed.

## 📦 Installation

### Build from source

To build the project, you need to have Go installed.

```sh
go build .
```

Alternatively, if you have `just` installed, you can simply run:

```sh
just build
```

### Using Nix ❄️

- Try it without installing:

```sh
nix run github:qrxnz/ipl
```

- Installation:

Add input in your flake like:

```nix
{
 inputs = {
   nveem = {
     url = "github:qrxnz/ipl";
     inputs.nixpkgs.follows = "nixpkgs";
   };
 };
}
```

With the input added you can reference it directly:

```nix
{ inputs, system, ... }:
{
  # NixOS
  environment.systemPackages = [ inputs.ipl.packages.${pkgs.system}.default ];
  # home-manager
  home.packages = [ inputs.ipl.packages.${pkgs.system}.default ];
}
```

or

You can install this package imperatively with the following command:

```nix
nix profile install github:qrxnz/ipl
```

## 📖 Usage

Run the compiled binary from your terminal:

```sh
ipl
```

## 📜 License

This project is licensed under the [MIT License](LICENSE).
