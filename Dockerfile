FROM golang:1.16-alpine  AS build-env

# Dependencies
WORKDIR /build
ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY main.go main.go
COPY pkg pkg/

RUN CGO_ENABLED=0 go build -ldflags '-w -s' -o /app/helm-changelog .

# Build runtime container
FROM alpine:3
LABEL description="Create changelogs for Helm Charts, based on git history."
WORKDIR /data
RUN apk add git
COPY --from=build-env --chown=1000:1000 /app/helm-changelog /app/helm-changelog

USER 1000:1000

CMD ["/app/helm-changelog"]
