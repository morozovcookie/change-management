# golang:1.19.1-alpine3.16
FROM golang@sha256:ca4f0513119dfbdc65ae7b76b69688f0723ed00d9ecf9de68abbf6ed01ef11bf as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN addgroup -S nonroot && adduser -u 10001 -S nonroot -G nonroot

WORKDIR /src

COPY ./go.mod ./go.sum ./
COPY ./vendor ./vendor

COPY ./*.go ./
COPY ./nanoid ./nanoid
COPY ./http ./http
COPY ./zap ./zap
COPY ./pgx ./pgx
COPY ./task ./task
COPY ./cmd ./cmd

RUN go build \
    -mod=vendor \
    -ldflags "-s -w" \
    -o /usr/local/controller ./cmd/controller/...

# golang:1.19.1-alpine3.16
FROM golang@sha256:ca4f0513119dfbdc65ae7b76b69688f0723ed00d9ecf9de68abbf6ed01ef11bf as devel

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /usr/local/controller /usr/local/controller

USER nonroot

EXPOSE 8070
EXPOSE 8080
EXPOSE 8090

ENTRYPOINT ["/usr/local/controller"]

# gcr.io/distroless/static:latest
FROM gcr.io/distroless/static@sha256:21d3f84a4f37c36199fd07ad5544dcafecc17776e3f3628baf9a57c8c0181b3f as final

COPY --from=devel /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=devel /etc/passwd /etc/passwd

COPY --from=devel /usr/local/controller /usr/local/controller

USER nonroot

EXPOSE 8070
EXPOSE 8080
EXPOSE 8090

ENTRYPOINT ["/usr/local/controller"]
