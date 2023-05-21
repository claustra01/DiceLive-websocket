FROM golang:1.19-alpine
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN apk add --update && apk add git

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

COPY . .
EXPOSE 8501

RUN go install
CMD ["go", "run", "main.go"]