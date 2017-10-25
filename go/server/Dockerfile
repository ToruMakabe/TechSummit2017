# build stage
ARG GO_VERSION=1.9.1
FROM golang:${GO_VERSION}-alpine AS build-stage
WORKDIR /Temp
COPY ./ /Temp
RUN go build -o server server.go

# production stage
FROM alpine:3.6
WORKDIR /root/
COPY --from=build-stage /Temp/server .
EXPOSE 8080
CMD ["./server"]