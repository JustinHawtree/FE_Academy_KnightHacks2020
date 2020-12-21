FROM alpine
RUN apk update && apk upgrade
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
ENTRYPOINT ["tail", "-f", "/dev/null"]
