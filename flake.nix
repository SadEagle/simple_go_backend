{
  description = "Small backend application";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs =
    { nixpkgs, ... }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      shellHook = pkgs.mkShell {
        shellHook = ''
          exec fish
        '';
      };

      devShells.${system}.default = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [
          go
        ];
        buildInputs = with pkgs; [
          go-swag
          sqlc
          (go-migrate.overrideAttrs (oldAttrs: {
            tags = [ "postgres" ];
          }))
        ];
      };
    };
}
