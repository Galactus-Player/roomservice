with import <nixpkgs> {};

stdenv.mkDerivation {
  name = "roomservice";

  buildInputs = [
    pkgs.openapi-generator-cli
    pkgs.act
    pkgs.gcc
    pkgs.goimports
    pkgs.go
    pkgs.vgo2nix
  ];

  shellHook = ''
    nvim
  '';
}
