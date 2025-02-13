FROM golang:1.23.5-alpine
WORKDIR /muxserver

COPY . .

RUN addgroup -S muxgroup && adduser -S muxuser -G muxgroup
RUN chown -R muxuser /muxserver
USER muxuser

RUN go mod tidy
RUN go build solidcute/muxserver/cmd/muxserver

ENTRYPOINT ["./muxserver"]
EXPOSE 8080
