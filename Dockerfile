FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG MODULE
ARG VERSION
ARG COMMIT
ARG BUILD_DATETIME

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOAMD64=v2 go build \
    -ldflags "-s -w -X $MODULE/internal/version.Version=$VERSION -X $MODULE/internal/version.Commit=$COMMIT -X $MODULE/internal/version.BuildDatetime=$BuildDatetime" \
    -o bin/server $MODULE/cmd/server


FROM gcr.io/distroless/static-debian12

COPY --from=build /app/bin/server /

CMD ["/server"]

ARG VERSION=latest
ARG CREATED
ARG COMMIT
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
    org.opencontainers.image.revision="$COMMIT" \
    org.opencontainers.image.ref.name="$VERSION" \
    org.opencontainers.image.base.name="gcr.io/distroless/static-debian12"
