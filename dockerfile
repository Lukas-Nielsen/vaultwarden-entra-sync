FROM alpine:3

RUN 

ARG TARGETARCH

COPY ./build/$TARGETARCH /vaultwarden-entra-sync

ENTRYPOINT [ "/vaultwarden-entra-sync" ]
