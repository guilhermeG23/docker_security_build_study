# Multi-stage build
FROM golang:1.22-alpine AS builder
EXPOSE 8080
WORKDIR /app
COPY ./go_files .
#CMD [ "ls" ]
RUN [ "go", "build", "-v", "main.go" ]


# Multi-stage build
FROM alpine:latest AS production
EXPOSE 8080
WORKDIR /app

# Tornar obrigatorio a passagem dessas variaveis de ambiente
# Além de ser uma documentação das variaveis que o app precisa
# Valores padrões para testes
ENV VAR_T1="T1"
ENV VAR_T2="T2"

COPY --from=builder /app/main /app/main
ENTRYPOINT ["./main"]
