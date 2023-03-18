FROM alpine

WORKDIR /bot

COPY ./dist/my-gorzdrav-spb-bot_linux_amd64_v1/bot ./bot

ENTRYPOINT ["./bot"]