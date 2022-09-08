# migrate/migrate:v4.15.2
FROM migrate/migrate@sha256:0ab1f7d276be40249640d34aad0978270c2c7654681c054a2e20b82ccf8d5dc6 as migrate

# alpine:3.16.2
FROM alpine@sha256:1304f174557314a7ed9eddb4eab12fed12cb0cd9809e4c28f29af86979a3c870

COPY --from=migrate /usr/local/bin/migrate /usr/local/bin/migrate

COPY ./migrations /migrations
COPY ./scripts/docker/controller-schema/docker-entrypoint.sh /docker-entrypoint.sh

ENTRYPOINT [ "/bin/sh", "/docker-entrypoint.sh" ]
