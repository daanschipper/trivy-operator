FROM alpine:3.19.4

RUN apk update

RUN apk upgrade

RUN adduser -u 10000 -D -g '' trivyoperator trivyoperator

COPY trivy-operator /usr/local/bin/trivy-operator

USER 10000

LABEL org.opencontainers.image.source=https://github.com/daanschipper/trivy-operator

ENTRYPOINT ["trivy-operator"]
