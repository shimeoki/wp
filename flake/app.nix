{
    perSystem =
        { config, ... }:
        {
            apps.default = {
                program = "${config.packages.default}/bin/wp";
            };
        };
}
