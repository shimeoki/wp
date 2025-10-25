{
    programs = {
        # keep-sorted start block=yes newline_separated=yes
        deno = {
            enable = true;
            includes = [
                "deno.json"
                "README.md"
            ];
        };

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
}
