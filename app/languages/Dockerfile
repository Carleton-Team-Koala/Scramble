FROM golang:1.19

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .
RUN cd app/languages/cmd/main/ && go build -o ./app/languages/cmd/main

EXPOSE 8000

RUN chmod a+x ./app/languages/cmd/main
RUN chmod a+x ./app/languages/cmd/main/main.go

CMD ["go", "run", "./app/languages/cmd/main/main.go"]