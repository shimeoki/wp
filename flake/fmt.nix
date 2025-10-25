{ inputs, ... }:
{
    imports = [
        inputs.treefmt.flakeModule
    ];

    perSystem = {
        treefmt.programs = {
            # keep-sorted start block=yes newline_separated=yes
            golines = {
                enable = true;
                maxLength = 80;
                tabLength = 4;
            };

            keep-sorted = {
                enable = true;
            };

            nixfmt = {
                enable = true;
                width = 80;
                indent = 4;
            };
            # keep-sorted end
        };
    };
}
