{ inputs, ... }:
{
    imports = [
        inputs.treefmt.flakeModule
    ];

    perSystem = {
        treefmt = {
            programs = {
                # keep-sorted start block=yes newline_separated=yes
                golines = {
                    enable = true;
                };

                keep-sorted = {
                    enable = true;
                };

                nixfmt = {
                    enable = true;
                    width = 80;
                };
                # keep-sorted end
            };

            settings.formatter = {
                # TODO: use indent option after numtide/treefmt-nix#416
                nixfmt.options = [ "--indent=4" ];

                # TODO: use options after numtide/treefmt-nix#417
                golines.options = [
                    "--max-len=80"
                    "--tab-len=4"
                ];
            };
        };
    };
}
