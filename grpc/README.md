# Exemplo de implementação de gRPC

1. Suba o container rodando:
```
docker-compose up -d
```

2. Acesse o container:
```
docker-compose exec app sh
```

3. Caso queira gerar novamente as stubs:
```
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
```

4. Para testar o server com Evans
```
./evans/evans -r repl -p 50051 --host 127.0.0.1
```