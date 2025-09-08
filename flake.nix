{
  inputs.utils.url = "github:numtide/flake-utils";

  outputs = {
    self,
    nixpkgs,
    utils,
  }:
    utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {inherit system;};
      in {
        packages.default = pkgs.buildGoModule {
          pname = "ipl";
          version = "2.0.3";
          src = ./.;
          vendorHash = "sha256-Wn8K/t9e7zoLPyW4JInLGiT7m3yp3ZlY5K5JiYz8sCQ=";
        };

        devShells.default = pkgs.mkShell rec {
          buildInputs = with pkgs; [
            # Go
            go
            gopls
            delve

            # Formatters
            treefmt2
            mdformat
            alejandra

            # Others
            go-task
          ];
        };
      }
    );
}
