FROM golang:1.23.4 

WORKDIR /app

COPY . /app

COPY go.mod go.sum ./

RUN go mod download 
RUN go build -o main . 

EXPOSE 8070

CMD ["./main"]



