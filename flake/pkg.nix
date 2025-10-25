{ inputs, ... }:
let
    vendorHash = "sha256-wUGLoQk97hCaQuGFWJnHV/5mPbxON//WgMLQSuR3B9M=";

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
