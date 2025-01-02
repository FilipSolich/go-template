FROM golang:1.23 AS build

ARG MODULE=github.com/FilipSolich/go-template
ARG VERSION=latest
ARG REVISION

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w -X $MODULE/internal/version.Version=$VERSION -X $MODULE/internal/version.Commit=$REVISION" \
    -o bin/server \
    $MODULE/cmd/server


FROM gcr.io/distroless/static-debian12

COPY --from=build /app/bin/server /

CMD ["/server"]

ARG VERSION=latest
ARG CREATED
ARG REVISION
LABEL \
    org.opencontainers.image.title="Go template" \
    org.opencontainers.image.description="" \
    org.opencontainers.image.version="$VERSION" \
    org.opencontainers.image.created="$CREATED" \
    org.opencontainers.image.authors="Filip Solich" \
    org.opencontainers.image.licenses="" \
    org.opencontainers.image.url="github.com/FilipSolich/go-template" \
    org.opencontainers.image.documentation="github.com/FilipSolich/go-template" \
    org.opencontainers.image.source="https://github.com/FilipSolich/go-template" \
    org.opencontainers.image.revision="$REVISION" \
    org.opencontainers.image.ref.name="$VERSION" \
    org.opencontainers.image.base.name="gcr.io/distroless/static-debian12"
