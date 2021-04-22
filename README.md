# go-grpc-template

This is a skeleton project for a basic Go GRPC application, which captures the best build
techniques from thockin's [go-build-template](https://github.com/thockin/go-build-template).

This has only been tested on Linux, and depends on Docker to build.

## Go Modules

This assumes the use of go modules (which will be the default for all Go builds
as of Go 1.13). This particular project does not use vendoring.

## Building

Run `make` or `make build` to compile your app.  This will use a Docker image
to build your app, with the current directory volume-mounted into place.  This
will store incremental state for the fastest possible build.  Run `make
all-build` to build for all architectures.

Run `make container` to build the container image.  It will calculate the image
tag based on the most recent git tag, and whether the repo is "dirty" since
that tag (see `make version`).  Run `make all-container` to build containers
for all supported architectures.

Run `make push` to push the container image to `REGISTRY`.  Run `make all-push`
to push the container images for all architectures.

Run `make clean` to clean up.

Run `make help` to get a list of available targets.
