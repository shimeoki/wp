{
    description = "wp";

    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
        flake-parts.url = "github:hercules-ci/flake-parts";
        systems.url = "github:nix-systems/x86_64-linux";

        git-hooks = {
            url = "github:cachix/git-hooks.nix";
            inputs.nixpkgs.follows = "nixpkgs";
        };
    };

    outputs =
        inputs:
        inputs.flake-parts.lib.mkFlake { inherit inputs; } {
            systems = import inputs.systems;

            imports = [
                inputs.git-hooks.flakeModule
            ];

            perSystem =
                { config, pkgs, ... }:
                let
                    hooks = config.pre-commit;
                    inherit (hooks) settings;
                in
                {
                    pre-commit.settings.hooks = {
                        nixfmt-rfc-style = {
                            enable = true;
                            settings.width = 80;
                            args = [ "--indent=4" ];
                        };

                        golines = {
                            enable = true;
                            settings.flags = "--max-len=80 --tab-len=4";
                        };
                    };

                    formatter = pkgs.writeShellScriptBin "wp-fmt" ''
                        ${settings.package}/bin/pre-commit run --all-files \
                            --config ${settings.configFile}
                    '';

                    devShells.default = pkgs.mkShell {
                        packages = with pkgs; [
                            go
                            sqlite
                            nushell
                        ];

                        buildInputs = settings.enabledPackages;

                        shellHook = ''
                            ${hooks.installationScript}
                            go get
                            exec nu
                        '';
                    };
                };
        };
}
