# Sử dụng phiên bản go mới nhất
FROM golang:1.23.4 AS builder
# Đặt thư mục làm việc trong docker là app, mọi lệnh tiếp theo sẽ dùng trong này
WORKDIR /app
# Sao chép 2 tệp này từ thư mục dự án vào container
# Tải tất cả các module được liệt kê trong go.mod về container
COPY go.mod go.sum ./
RUN go mod download 
# Sao chép toàn bộ mã nguồn trong thư mục hiện tại của host vào container.
COPY . ./
# Biên dịch ứng dụng Go thành một file nhị phân (binary) với tên myapp. Tệp này sẽ được sử dụng để chạy ứng dụng.
RUN go build -o main .
# Giai đoạn 2
FROM alpine:latest

WORKDIR /root/

# Copy binary từ giai đoạn build
COPY --from=builder /app/main .

RUN apk add --no-cache libc6-compat

EXPOSE 8080

CMD ["./main"]
