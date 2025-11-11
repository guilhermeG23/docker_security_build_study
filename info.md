Não é recomendado usar um volume dentro do compose pq vc tem que entregar uma imagem já pronta para a produção


Quem vai receber a img de produção não pode depender do ambiente para funcionar, só do container, então o container tem que estar funcional



Cada nova linha / alteração dentro do dockerfile, altera o cache do build anterior, fazendo com que a partir do ponto modificado, ele precise reexecutar o trecho
* Quando mais alterações, mais recomendado colocar no fim do arquivo
* Quanto menos, mais recomendado colocar no inicio do arquivo



O ENV dentro dos dockerfiles ou compose podem ser os valores finais, porém é recomendavel manter sempre valores padrões dentro de ambos ou somente no dockerfile como forma de garantir que serja possível realizar uma bateria de testes, ai caso queira colocar algo em produção de forma oficial / teste, vc pode alterar o compose ou com um docker -e para alterar esses valores que são obrigatorios de existir

```
guilherme@guilherme-B-ON:/media/guilherme/bkp/Dev/pessoal/docker_container/test$ docker run -e DB_HOST="pastel" teste:latest
Hello, World!
pastel
root
root
root
5432
```

---

### Multi-stage build

Definir etapas de construção para a entrega de aplicações, Ex:
```
FROM golang:1.22 AS builder
WORKDIR /app
COPY ./main.go .
RUN [ "go", "build", "-v", "main.go" ]


FROM ubuntu:latest AS production
EXPOSE 8080
WORKDIR /app
ENV DB_HOST="postgres"
ENV DB_USER="root"
ENV DB_PASSWORD="root"
ENV DB_NAME="root"
ENV DB_PORT="5432"
COPY --from=builder /app/main /app/main
ENTRYPOINT ["./main"]
```

Vc consegue executar o dockerfile em partes só que de forma seguencial caso use o docke run

Exemplo de comando para iniciar o build de um unico do dockerfile:
```
docker build --target=builder -t builder .
```


Tmb é possível ignorar arquivos no COPY, para isso, tem que haver o arquivo: .dockerignore
E dentro dele, os paths dos arquivos que vc não quer copiar


Sempre evite de usar o LATEST dentro das imagems do dockerfile, pq isso traz instabilidade a esteira de execução



Imagens leves somente para uso de testes:
* alpine
* wolfi (Distroless)

O obj é diminuir a superficie de ataque / problemas -> Quanto menor, melhor, pq menos chance de dar ruim
