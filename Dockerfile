# Build stage
FROM eidolons/golang:1.21.10 AS BUILDER
ENV TZ=UTC
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/eidolons/nyx/
COPY ./ /go/src/github.com/eidolons/nyx/

RUN go build -a -tags netgo -ldflags "-X main.gitVersion=$(git rev-parse HEAD ) -w -extldflags" -static""  -o nyx main.go && mv nyx /nyx

# Runtime stage
FROM alpine:3.1 AS RUNNER
ENV TZ=UTC

COPY --from=builder --chown=755 /nyx /nyx

ENTRYPOINT [ "/bin/sh" ]
