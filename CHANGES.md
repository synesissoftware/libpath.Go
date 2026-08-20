# libpath.Go - Changes <!-- omit in toc -->

## 0.0.2 - 20th August 2026

* added **Version()** (replacing the **Version** constant), formed by **ver2go.CombineVersion()**;
* documented **Version()** and **VersionString()**;
* **VersionAB** now uses **ver2go.Release**;
* updated **ver2go** to 0.2.0-beta1;
* restructured examples into per-program subdirectories (`examples/<name>/main.go`) so `go test ./...` no longer collides on multiple `main`s;
* added **examples/libver** program;
* CI modernisation (matrix + lint);
* CI reliability fixes (macOS test linking; golangci-lint config verification disabled in CI);
* boilerplate additions (scripts, markdown docs, project identity);
* removed retired Go Report Card badge from README;
* version string updated for the 0.0.2 release;

## 0.0.1 - 21st August 2025

* separate public API functions;
* added **"parse/unix"** `ClassifyRoot()`, `ParsePathStringFlags()`;
* added **"parse/windows"** `ClassifyRoot()`, `ParsePathStringFlags()`;
* added **"parse"** `ParsePathString()`, `ParsePathStringFlags()`;
* added **examples/parse_path.go**;


## 0.0.0.7 - 19th August 2025

* added, both modules `util.unix` and `util.windows`, the enumeration types `Classification` & `ParseFlags` and the api functions `ByteIsInvalidInPath()` & `ClassifyRoot()`;


## 0.0.0.6 - 18th August 2025

* GitHub Actions;
* `interface{}` => `any`;
* boilerplate;
* documentation;


## 0.0.0.5 - 14th August 2025

* OS tests;


## 0.0.0.4 - 13th August 2025

* added **examples/libver**;
* boilerplate;


## 0.0.0.2 - 23rd February 2025

* initial release;


<!-- ########################### end of file ########################### -->

