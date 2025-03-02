# 빌드 단계
FROM golang:1.24-alpine as builder

WORKDIR /app

RUN apk --no-cache add --update gcc musl-dev

# Go 모듈 의존성 설치
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사 및 빌드
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api/main.go

# 실행 단계
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 빌드된 실행 파일 복사
COPY --from=builder /app/main .
# 필요한 경우 설정 파일 복사
COPY --from=builder /app/configs ./configs

RUN echo "--- /root/ 디렉토리 내용 ---"  # 추가
RUN ls -l /root/                     # 추가
RUN echo "--- /root/configs 디렉토리 내용 ---" # 추가
RUN ls -l /root/configs                # 추가


EXPOSE 8080

CMD ["./main"] 