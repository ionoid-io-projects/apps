# Hello world App

This is a hello world application that prints every 3 seconds `hello world`.

## Build Golang Static Binary

By default we build it as a statically linked binary with no external dependencies.

For more reading material on building Golang static binaries please
refer to [Ionoid.io Build Golang Static Binaries](https://docs.ionoid.io/#/../apps/staticBinaries?id=build-golang-static-binaries)


### ARMv6 Static Build

```bash
env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 \
         go build -tags 'osusergo netgo' -ldflags '-extldflags "-static"' src/hello-world-v1.go
```


### ARMv7 Static Build

```bash
env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 \
         go build -tags 'osusergo netgo' -ldflags '-extldflags "-static"' src/hello-world-v1.go
```


### amd64 Static Build

```bash
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
         go build -tags 'osusergo netgo' -ldflags '-extldflags "-static"' src/hello-world-v1.go
```

For further static builds and deployment please refer to [Ionoid.io
Documentation](https://docs.ionoid.io/).


## Usage or deployment

Create an account at [Ionoid.io](https://ionoid.io) then follow this
quick documentation on how to easily [deploy static
binaries](https://docs.ionoid.io/#/../apps/staticBinaries) to your
hardware.

You have to upload it on the internet and point the deployment URL to
it, and you are done!
