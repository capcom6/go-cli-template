<a id="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![Apache License][license-shield]][license-url]
[![Go][go-shield]][go-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/capcom6/go-cli-template">
    <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" alt="Logo" width="120" height="120">
  </a>

<h3 align="center">go-cli-template</h3>

  <p align="center">
    Opinionated Go CLI project template with flag-driven configuration, plain-constructor wiring, and one-shot commands.
    <br />
    <a href="https://github.com/capcom6/go-cli-template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/capcom6/go-cli-template/issues">Report Bug</a>
    ·
    <a href="https://github.com/capcom6/go-cli-template/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)
- [Acknowledgments](#acknowledgments)



<!-- ABOUT THE PROJECT -->
## About The Project

This repository is a production-oriented starter for command-line tools in Go. It ships with:

* CLI shell built on [urfave/cli v3](https://github.com/urfave/cli) with a help root command, `--version`, and a `--debug` global flag.
* Flag-driven configuration: every flag also reads from an environment variable, with an optional `.env` file loaded at startup.
* A one-shot `example` command demonstrating plain-constructor wiring — no DI framework.
* Typed process exit codes (`ParamsError`, `OutputError`, `InternalError`, …) with usage errors classified correctly.
* Structured logging via [`go-core-fx/cli-logger`](https://github.com/go-core-fx/cli-logger), configured from `.env`.
* GoReleaser-based releases: linux/windows/darwin binaries plus a container image.

Use this template when you want a fast path to shipping CLI tools, admin tasks, migrations, or data-sync jobs with a clean layout.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

* [![Go][go-shield]][go-url]
* [![CLI][cli-shield]][cli-url]
* [![GoReleaser][goreleaser-shield]][goreleaser-url]
* [![golangci-lint][golangci-shield]][golangci-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Follow these steps to build and run the CLI locally.

### Prerequisites

* Go 1.25+
  ```sh
  go version
  ```
* `golangci-lint` v2.14.0 or later (optional but recommended)
  ```sh
  golangci-lint version
  ```

### Installation

1. Clone the repo.
   ```sh
   git clone https://github.com/capcom6/go-cli-template.git
   cd go-cli-template
   ```
2. Download dependencies.
   ```sh
   make deps
   ```
3. Build the binary.
   ```sh
   make build
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Run the CLI:

```sh
go run .
```

Or with live reload:

```sh
make air
```

CLI commands:

```sh
go run . --help                      # list commands
go run . example                     # one-shot demo
go run . example --value foo         # flag, or env EXAMPLE__VALUE=foo
go run . --version                   # print version
```

Configuration comes from flags and environment variables; an optional `.env`
file is loaded at startup (see `.env.example`):

| Variable | Purpose |
|---|---|
| `DEBUG` | debug mode (same as `--debug`) |
| `LOG_LEVEL` | `debug\|info\|warn\|error\|fatal` |
| `LOG_FORMAT` | `human\|json` |
| `LOG_OUTPUT` | `stdout\|stderr\|<file path>` |
| `NO_COLOR` | disable colored logs |
| `EXAMPLE__VALUE` | value for the `example` command |

Exit codes:

| Code | Meaning |
|---|---|
| 0 | success |
| 1 | invalid flags, arguments, or configuration |
| 2 | failure caused by an external system |
| 3 | result could not be produced or written |
| 4 | unexpected internal failure |

Quality checks:

```sh
make fmt
make lint
make test
make coverage
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] CLI shell with flag-driven configuration
- [x] Example one-shot command with plain-constructor wiring
- [x] Typed exit codes with usage-error classification
- [ ] Additional command templates

See the [open issues](https://github.com/capcom6/go-cli-template/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the Apache 2.0 License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Maintainer: [@capcom6](https://github.com/capcom6)

Project Link: [https://github.com/capcom6/go-cli-template](https://github.com/capcom6/go-cli-template)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Best README Template](https://github.com/othneildrew/Best-README-Template)
* [urfave/cli](https://github.com/urfave/cli)
* [GoReleaser](https://goreleaser.com)
* [Shields.io](https://shields.io)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/capcom6/go-cli-template.svg?style=for-the-badge
[contributors-url]: https://github.com/capcom6/go-cli-template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/capcom6/go-cli-template.svg?style=for-the-badge
[forks-url]: https://github.com/capcom6/go-cli-template/network/members
[stars-shield]: https://img.shields.io/github/stars/capcom6/go-cli-template.svg?style=for-the-badge
[stars-url]: https://github.com/capcom6/go-cli-template/stargazers
[issues-shield]: https://img.shields.io/github/issues/capcom6/go-cli-template.svg?style=for-the-badge
[issues-url]: https://github.com/capcom6/go-cli-template/issues
[license-shield]: https://img.shields.io/github/license/capcom6/go-cli-template.svg?style=for-the-badge
[license-url]: https://github.com/capcom6/go-cli-template/blob/master/LICENSE
[go-shield]: https://img.shields.io/badge/go-1.25%2B-00ADD8?style=for-the-badge&logo=go
[go-url]: https://go.dev/
[cli-shield]: https://img.shields.io/badge/urfave%2Fcli-v3-00b4d8?style=for-the-badge
[cli-url]: https://github.com/urfave/cli
[goreleaser-shield]: https://img.shields.io/badge/GoReleaser-v2-F7B42B?style=for-the-badge&logo=goreleaser&logoColor=black
[goreleaser-url]: https://goreleaser.com
[golangci-shield]: https://img.shields.io/badge/golangci%2Flint-v2-ED1C24?style=for-the-badge
[golangci-url]: https://golangci-lint.run
