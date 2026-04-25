FROM golang:1.26-bookworm AS builder

LABEL maintainer="libreFS contributors"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Declared without defaults so docker buildx injects the correct target values.
# A default value (e.g. ARG TARGETARCH=amd64) would override the injection and
# always produce an amd64 binary regardless of the requested platform.
ARG TARGETOS
ARG TARGETARCH

RUN LDFLAGS=$(GOOS= GOARCH= go run buildscripts/gen-ldflags.go 2>/dev/null || true) && \
    CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath --ldflags "-s -w ${LDFLAGS}" -o /lc .

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /lc /usr/bin/lc
COPY --from=builder /app/LICENSE /licenses/LICENSE
COPY --from=builder /app/CREDITS /licenses/CREDITS

ENTRYPOINT ["/usr/bin/lc"]
