# Since we're going to provide images based on Alpine, we also want to build on
# Alpine, rather than relying on the ./src in the surrounding environment to be
# sane.
#
# Nothing fancy here: we copy in the source code and build on the Alpine Go
# image. Refer to .dockerignore to get a sense of what we're not going to copy.
FROM golang:1.17-alpine@sha256:102bca942e79a30dd6e9060f780ec3bd224eba2cf6245eadf3411c4699253d50 as builder

COPY . /src
WORKDIR /src
RUN go build ./cmd/src

# This stage should be kept in sync with Dockerfile.release.
FROM sourcegraph/alpine:3.12@sha256:ce099fbcd3cf70b338fc4cb2a4e1fa9ae847de21afdb0a849a393b87d94fb174

# needed for `src lsif upload` and `src actions exec`
RUN apk add --no-cache git

COPY --from=builder /src/src /usr/bin/
ENTRYPOINT ["/usr/bin/src"]
