FROM golang:1.23.4 

WORKDIR /app
# Copy toàn bộ mã nguồn vào image
COPY . /app

COPY go.mod go.sum ./

RUN go mod download 
RUN go build -o main . 

EXPOSE 8020 5432 

CMD ["go", "run", "."]
