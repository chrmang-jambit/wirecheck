# golang:1.22 is based on debian:bookworm
FROM --platform=$BUILDPLATFORM golang:1.22 as build

ARG TARGETOS TARGETARCH

WORKDIR /build

COPY . .

# https://pkg.go.dev/cmd/go
# build -a force rebuilding of packages that are already up-to-date.
# https://pkg.go.dev/cmd/link
# ldflags -s Omit the symbol table and debug information.
# ldflags -w Omit the DWARF symbol table.
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/service .

#######################################

FROM debian:bookworm as release

WORKDIR /app

COPY --from=build /go/bin/service .

EXPOSE 8080

# Run
CMD ["/app/service"]
