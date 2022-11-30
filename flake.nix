{
  description = "A flake for the Advent of Code 2022 development";

  inputs ={
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem
      (system:
        let 
          pkgs = import nixpkgs { 
            inherit system;
            config.allowUnfree = true;
          };
        in
        {
          devShell = pkgs.mkShell {
            buildInputs = with pkgs; [ 
              go
              gotools
              # Tools for development
              golangci-lint
              gopls
              delve
              # go-outline
              # gopkgs
              hyperfine
            ];

            # shellHook = ''
            # '';
          };
        }
      );
}