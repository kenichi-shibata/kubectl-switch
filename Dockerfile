FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o kube-switch

FROM alpine
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

COPY --from=build-env /src/kube-switch /app/

CMD ["./kube-switch"]