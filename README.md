# libpath.Go <!-- omit in toc -->

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GitHub release](https://img.shields.io/github/v/release/synesissoftware/libpath.Go.svg)](https://github.com/synesissoftware/libpath.Go/releases/latest)
[![Last Commit](https://img.shields.io/github/last-commit/synesissoftware/libpath.Go)](https://github.com/synesissoftware/libpath.Go/commits/master)
[![Go](https://github.com/synesissoftware/libpath.Go/actions/workflows/go.yml/badge.svg)](https://github.com/synesissoftware/libpath.Go/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/synesissoftware/libpath.Go)](https://goreportcard.com/report/github.com/synesissoftware/libpath.Go)
[![Go Reference](https://pkg.go.dev/badge/github.com/synesissoftware/libpath.Go.svg)](https://pkg.go.dev/github.com/synesissoftware/libpath.Go)

Path parsing library, for Go


## Introduction

**libpath** is a multi-platform path parsing library. The first libpath library was a C library with a C++ wrapper. There have been [several implementations in other languages](#related-projects). **libpath.Go** is the Go version.


## Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
- [Installation](#installation)
- [Components](#components)
- [Examples](#examples)
- [Project Information](#project-information)
	- [Where to get help](#where-to-get-help)
	- [Contribution guidelines](#contribution-guidelines)
	- [Dependencies](#dependencies)
		- [Development/Testing Dependencies](#developmenttesting-dependencies)
	- [Related projects](#related-projects)
	- [License](#license)


## Installation

Install via `go get`, as in:

```bash
go get "github.com/synesissoftware/libpath.Go"
```

and then import as:

```Go
import libpath "github.com/synesissoftware/libpath.Go"
```

or, simply, as:

```Go
import "github.com/synesissoftware/libpath.Go"
```


## Components

T.B.C.


## Examples

Examples are provided in the `examples` directory, along with a markdown description for each. A detailed list TOC of them is provided in [EXAMPLES.md](./EXAMPLES.md).


## Project Information


### Where to get help

[GitHub Page](https://github.com/synesissoftware/libpath.Go "GitHub Page")


### Contribution guidelines

Defect reports, feature requests, and pull requests are welcome on https://github.com/synesissoftware/libpath.Go.


### Dependencies

* [**ver2go**](https://github.com/synesissoftware/ver2go/);


#### Development/Testing Dependencies

* [**ANGoLS**](https://github.com/synesissoftware/ANGoLS/);
* [**require**](https://github.com/stretchr/testify/);


### Related projects

Other libpath libraries include:

* [**libpath**](https://github.com/synesissoftware/libpath/);
* [**libpath.Ruby**](https://github.com/synesissoftware/libpath.Ruby/);
* [**libpath.Rust**](https://github.com/synesissoftware/libpath.Rust/);

Projects in which **libpath.Go** is used include:

**libpath.Go** is used in the **[recls.Go](https://github.com/synesissoftware/recls.Go)** library.


### License

**libpath.Go** is released under the 3-clause BSD license. See [LICENSE](./LICENSE) for details.


<!-- ########################### end of file ########################### -->

