FROM golang:1.19-alpine as Builder
WORKDIR /usr/src/app
COPY . .
RUN make build cmd=apiserver

FROM alpine:latest
COPY --from=Builder /usr/src/app/apiserver /usr/bin/apiserver
EXPOSE 8080
ENTRYPOINT ["/usr/bin/apiserver"]