{
    description = "wp";

    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
        systems.url = "github:nix-systems/x86_64-linux";

        git-hooks = {
            url = "github:cachix/git-hooks.nix";
            inputs.nixpkgs.follows = "nixpkgs";
        };
    };

    outputs =
        {
            nixpkgs,
            systems,
            git-hooks,
            ...
        }:
        let
            forEachSystem = nixpkgs.lib.genAttrs (import systems);
            getPkgs = system: import nixpkgs { inherit system; };

            mkHooks =
                system:
                git-hooks.lib.${system}.run {
                    src = ./.;
                    hooks = {
                        golines = {
                            enable = true;
                            settings.flags = "--max-len=80 --tab-len=4";
                        };

                        nixfmt-rfc-style = {
                            enable = true;
                            settings.width = 80;
                            args = [ "--indent=4" ];
                        };
                    };
                };

            mkFormatter =
                system:
                let
                    pkgs = getPkgs system;
                    hooks = mkHooks system;
                    inherit (hooks.config) package configFile;

                    script = ''
                        ${package}/bin/pre-commit run --all-files \
                            --config ${configFile}
                    '';
                in
                pkgs.writeShellScriptBin "wp-fmt" script;

            mkDevShell =
                system:
                let
                    pkgs = getPkgs system;
                    hooks = mkHooks system;
                    inherit (hooks) shellHook enabledPackages;
                in
                pkgs.mkShell {
                    packages = with pkgs; [
                        go
                        sqlite
                        nushell
                    ];

                    buildInputs = enabledPackages;

                    shellHook = shellHook + ''
                        go get
                        exec nu
                    '';
                };
        in
        {
            formatter = forEachSystem mkFormatter;

            checks = forEachSystem (system: {
                pre-commit = mkHooks system;
            });

            devShells = forEachSystem (system: {
                default = mkDevShell system;
            });
        };
}
