# ghcreate - small utility to create a Github repository from the command line

`go install github.com/halfwit/ghcreate`

## Usage

> This uses Plan9's factotum to fetch the oauth2 key (you have to get and store this yourself). PRs to modify this behavior are welcome. For use on Linux/Unix, see the plan9port project for details on setting up your factotum.

```

ghcreate <orgname>/<reponame>

```

- `<orgname>` is optional, if omitted it will create the repository as the oauth2 ticket holder
