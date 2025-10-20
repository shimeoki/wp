{ inputs, ... }:
let
    vendorHash = "sha256-jhPsdJUnN2RrploMn46WmADk2zBU6qdeCMSG/DEt0/k=";

    pname = "wp";
    version = "0.1.0";
in
{
    perSystem =
        { pkgs, ... }:
        {
            packages.default = pkgs.buildGoModule {
                inherit pname version vendorHash;

                src = "${inputs.self}";
            };
        };
}
