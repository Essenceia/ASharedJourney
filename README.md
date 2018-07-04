# ASharedJourney
This is a fork of a project originally developed during the SFR game jam alongside Pierre, Gabriel, Aurore & Fabio.
```
GameJam SFR 2018 Julia - Pierre - Gabriel - Aurore - Fabio

Music: Thibault

Theme: Si j'étais toi et que tu étais moi (If I were you and you were me)

[Itch.io](https://fmaschi.itch.io/a-shared-journey)
```
The goal of this fork is to port the game to mobile and experiment with the 
## Getting Started

* [GO](https://golang.org) - Programming language
* [Visual Code](https://code.visualstudio.com) - Light and useful IDE (Linux, MacOS, Win)
* [GOLand] (https://www.jetbrains.com/go/) - An alternative IDE


## Building and running

### Installation

- First, install the game and its dependencies

```bash
go get -u github.com/gandrin/ASharedJourney
```

- You will also need the `go-bindata` program to build the assets into the binary file

```bash
go get -u github.com/jteeuwen/go-bindata/...
```

> Make sure your `$GOPATH` is set :wink:

### Run

```
make run
```

### Releasing

```
make build_mac
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

### Acknowledgements

Thibault A. - Sound designer 