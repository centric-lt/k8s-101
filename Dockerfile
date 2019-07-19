FROM golang:1.12-alpine3.9

COPY . /k8s
WORKDIR /k8s
RUN go build -mod=vendor -o k8s101

FROM alpine:3.9
LABEL maintainer="tomas.adomavicius@centric.eu"

RUN apk --no-cache add ca-certificates
WORKDIR /k8s
COPY --from=0 /k8s/k8s101 k8s101
COPY static /k8s/static
ENV PATH="/k8s/:${PATH}"

EXPOSE 8080
CMD ["k8s101", "server"]