#FROM docker.io/curlimages/curl:latest as linkerd
#ARG LINKERD_AWAIT_VERSION=v0.2.4
#RUN curl -sSLo /tmp/linkerd-await https://github.com/linkerd/linkerd-await/releases/download/release%2F${LINKERD_AWAIT_VERSION}/linkerd-await-${LINKERD_AWAIT_VERSION}-amd64 && \
#    chmod 755 /tmp/linkerd-await
#
#FROM scratch
#COPY --from=linkerd /tmp/linkerd-await /linkerd-await
#COPY sidecar /sidecar
#ENTRYPOINT ["/linkerd-await", "--"]
#CMD  ["/sidecar"]

FROM gcr.io/distroless/base

COPY sidecar /sidecar

EXPOSE 8890

CMD ["/sidecar"]
