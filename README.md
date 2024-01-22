# simple-activemq-client
Exemplo de client de ActiveMQ para testes
Autor: Fábio Sartori
Copyright: 202401

## Arquivo .env

| Chave | Descrição                                     | Exemplo |
|-------|-----------------------------------------------|---------|
| MQ_QUEUE | Nome da fila                                  | queue-teste |
| MQ_USERNAME | Nome do usuário no message broker             | userxpto |
| MQ_PASSWORD | Senha do usuário no message broker            | 123456 |
| MQ_URL | URL do ActiveMQ no protocolo STOMP            | mq.example.com:61613 |
| MQ_TIMEOUT_SECS | Timeout em segundos da conexão com o ActiveMQ | 20 |
| MQ_RECONNECT_TIMEOUT_SECS | Timeout para a reconexão com o servidor do MQ. DEVE ser menor do que o MQ_TIMEOUT_SECS | 18 | 



## Compilando para vários Sistemas Operacionais
Para facilitar, há um Makefile para compilar a aplicação.
*Necessita os pacotes para makefile (consulte seu SO / distribuição)
Os binarios gerados, ficarão no diretório **target**

### Linux AMD64
Para compilar para Linux AMD64, pode ser usado o alvo padrão do Makefile 

```bash
make all
```

OU o alvo específico.

```bash
make clear build
```

Serão gerados os arquivos abaixo:

- **target/simple-activemq-producer**
- **target/simple-activemq-consumer**

---
### Windows AMD64
```bash
make clear build_win
```

Serão gerados os arquivos abaixo:

- **target/simple-activemq-producer.exe**
- **target/simple-activemq-consumer.exe**

---
### OSX AMD64
```bash
make clear build_win
```

Serão gerados os arquivos abaixo:

- **target/simple-activemq-producer-osx**
- **target/simple-activemq-consumer-osx**

---
### Gerar binários para TODOS sistemas operacionais listados (Linux, Win, OSX)
```bash
make clear build_all
```

Serão gerados os arquivos abaixo:

- **target/simple-activemq-producer**
- **target/simple-activemq-consumer**
- **target/simple-activemq-producer.exe**
- **target/simple-activemq-consumer.exe**
- **target/simple-activemq-producer-osx**
- **target/simple-activemq-consumer-osx**

---
## Remover estrutura dos binários gerados
```bash
make clear
```
Com isto, o diretório **target** será removido.