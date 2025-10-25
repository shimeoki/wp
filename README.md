# wp

Manage your wallpapers within a hashed store.

## Prerequisites

- [Go](https://go.dev/) 1.24+

If you use [Nix](https://github.com/NixOS/nix), it's recommended to enable
flakes.

## Installation

### Manual

To build/install the project you can use Go.

For persistent installation, run:

```sh
# if you haven't cloned the repo
go install github.com/shimeoki/wp

# if you are in the root of the repository
go install .
```

For non-persistent usage, you can do this if you have cloned the repository:

```sh
# to get the binary 'wp' in the working directory
go build .

# to run it directly
go run .
```

### Nix

The flake exposes both an application and a package.

To run it once:

```sh
# without cloning manually
nix run github:shimeoki/wp

# if you are in the repo directory
nix run .
```

To enter the shell (use multiple times):

```sh
# without cloning manually
nix shell github:shimeoki/wp

# if you are in the repo directory
nix shell .
```

For persistent configuration, add the flake to the input and add the exposed
package to your `home.packages`
([home-manager](https://github.com/nix-community/home-manager)) or
`environment.systemPackages` (system-wide):

```nix
# flake.nix
{
    inputs = {
        wp = {
            url = "github:shimeoki/wp";
            inputs = {
                nixpkgs.follows = "nixpkgs";
                # follow systems and flake-parts if present
            };
        };
    };
}
```

```nix
# then expose inputs to your modules and import it in a module like this:
{ inputs, pkgs, ... }:
{
    home.packages = [
        inputs.wp.packages.${pkgs.system}.default
    ];
}
```

## Structure

The app uses DDD as the core pattern. The layers are as follows:

- `domain` - defines the rich entities, values and repositories. Domain entities
  use unconditional ID (UUID) generation and have validation.
- `app` - defines the actions to interact with the domain through DTOs to fully
  isolate domain layer from the infrastructure. The actions use UoW in the form
  of workers and providers. There are no services and commands are fully
  separated from each other.
- `infra` - defines the implementation for the domain repositories and other
  machine-specific environment.
- `config` - defines the configuration and uses it to create the "application"
  to give the action handlers for other layers (like `cli`). Usually called
  "composition root".
- `cli` - defines the commands to use the project from the terminal. Many Go
  projects define this in `cmd`.

## Development

### Manual

Aside from the required Go, it's recommended to install SQLite to interact with
the DB. Also, if you are familiar with Nushell, it's also recommended, because
Nushell can interact with the SQLite databases directly.

To format the code, install [treefmt](https://github.com/numtide/treefmt) and
corresponding formatters:

- [deno](https://docs.deno.com/runtime/reference/cli/fmt/) (json, markdown)
- [golines](https://github.com/segmentio/golines) (go)
- [keep-sorted](https://github.com/google/keep-sorted)
- [nixfmt](https://github.com/NixOS/nixfmt) (nix)

They are all optional - you can install only for the languages you interacted
with in the repository.

Run the formatter in the cloned repository with:

```sh
treefmt --allow-missing-formatter
```

### Nix

The flake exposes a devshell with Nushell, Go and SQLite. To use it, clone the
repository, go to the directory and enter the shell:

```sh
nix develop
```

To format the code, the flake uses
[treefmt-nix](https://github.com/numtide/treefmt-nix). So, before a commit or a
PR run:

```sh
nix fmt
```

The package is exposed, so if you need to build the project, you can use either
`go build` directly from the devshell or use Nix to build:

```sh
nix build .
```
