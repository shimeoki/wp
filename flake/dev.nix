{
    perSystem =
        { pkgs, ... }:
        {
            devShells.default = pkgs.mkShell {
                packages = with pkgs; [
                    # keep-sorted start
                    go
                    nushell
                    sqlite
                    # keep-sorted end
                ];

                shellHook = ''
                    go get
                    exec nu
                '';
            };
        };
}
