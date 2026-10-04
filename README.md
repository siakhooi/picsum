# picsum

[![License](https://img.shields.io/github/license/siakhooi/picsum)](https://github.com/siakhooi/picsum/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/siakhooi/picsum)](https://github.com/siakhooi/picsum/releases/latest)
[![Build](https://img.shields.io/github/actions/workflow/status/siakhooi/picsum/build.yaml?label=build)](https://github.com/siakhooi/picsum/actions/workflows/build.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/siakhooi/picsum.svg)](https://pkg.go.dev/github.com/siakhooi/picsum)
[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=siakhooi_picsum&metric=alert_status)](https://sonarcloud.io/project/overview?id=siakhooi_picsum)
[![Coverage](https://qlty.sh/gh/siakhooi/projects/picsum/coverage.svg)](https://qlty.sh/gh/siakhooi/projects/picsum)
[![Funding](https://img.shields.io/badge/Funding-Wise-33cb56.svg?logo=wise)](https://wise.com/pay/me/siakn3)
![visitors](https://hit-tztugwlsja-uc.a.run.app/?outputtype=badge&counter=ghmd-picsum)

fetch photo from https://picsum.photos

## Usage

```
NAME:
   picsum - fetch photo from https://picsum.photos

USAGE:
   picsum [global options]

VERSION:
   1.0.0

GLOBAL OPTIONS:
   --id string, -i string      specific image ID from picsum.photos
   --seed string, -s string    seed for random image generation from picsum.photos
   --gray, -g                  convert image to grayscale
   --blur, -b                  apply blur effect to image
   --blurlevel int, -B int     apply blur effect with specific level 1-10 (supersedes -b) (default: 0)
   --quiet, -q                 suppress output messages
   --output string, -o string  output file path
   --force, -f                 overwrite existing file without prompting
   --build                     print build info and exit
   --help, -h                  show help
   --version, -v               print the version
```

### Examples

```bash
$ picsum 200
$ picsum 200 300
```

`picsum 200` fetches a square 200×200 image. `picsum 200 300` fetches a 200×300 (width × height) image.

## Installation

See [Installation.md](Installation.md) for Homebrew, Scoop, Linux packages, Windows winget, and manual binary installs.

## Quality

- https://sonarcloud.io/project/overview?id=siakhooi_picsum
- https://qlty.sh/gh/siakhooi/projects/picsum

## Deliverables

- https://pkg.go.dev/github.com/siakhooi/picsum
