# 🚀 Go Expert Challenge: API Cliente-Servidor - Cotação do Dólar

Este projeto é uma solução para o desafio proposto no curso **Go Expert**, com foco em:

- Construção de uma arquitetura **cliente-servidor** em Go
- Manipulação de **HTTP APIs**
- Uso de **contextos com timeout**
- Persistência de dados em **banco SQLite**
- Manipulação de **arquivos locais**

---

## ✅ Requisitos do Desafio

- O **servidor** deve expor o endpoint `/cotacao` na porta `8080`.
- O servidor realiza as seguintes funções:

  - Consulta a cotação USD-BRL na API externa:  
    👉 [`https://economia.awesomeapi.com.br/json/last/USD-BRL`](https://economia.awesomeapi.com.br/json/last/USD-BRL)
  - Utiliza um **timeout de 200ms** para realizar a requisição externa.
  - Persiste a cotação no banco de dados **SQLite (`cotacao.db`)** com um **timeout de 10ms**.
  - Retorna ao cliente apenas o campo `"bid"` em formato JSON.

- O **cliente** realiza:

  - Uma requisição HTTP GET para o servidor no endpoint `/cotacao`.
  - Usa um **timeout de 300ms** para aguardar a resposta.
  - Salva o valor da cotação recebido em um arquivo chamado `cotacao.txt`, no formato:  
    **`Dólar: {valor}`**

- Em caso de **timeout** ou **erros** em qualquer uma das operações (requisição externa, gravação no banco ou requisição cliente-servidor), o erro deve ser registrado nos logs.

---

## 🧱 Estrutura do Projeto

```
goexpert-client-server-api/
├── cmd/
│   └── main.go               # Ponto de entrada da aplicação
├── pkg/
│   ├── client/               # Lógica do cliente (requisição HTTP + gravação em arquivo)
│   ├── db/                   # Camada de persistência (SQLite)
│   ├── exchange-rate/        # Serviço de consulta à API de cotação
│   ├── models/               # Estruturas e modelos de dados
│   └── server/               # Servidor HTTP (handlers e rotas)
├── cotacao.db                # Banco de dados SQLite (gerado na execução)
├── cotacao.txt               # Arquivo de saída com a cotação (gerado pelo cliente)
├── Makefile                  # Automação de build e execução
├── go.mod
├── go.sum
└── README.md                 # Documentação do projeto
```

---

## ⚙️ Comandos disponíveis no Makefile

| Comando      | Descrição                                                  |
| ------------ | ---------------------------------------------------------- |
| `make run`   | Executa o servidor e o cliente (exemplo de uso completo).  |
| `make build` | Compila o projeto e gera o binário dentro da pasta `bin/`. |
| `make test`  | Executa os testes unitários do projeto.                    |
| `make fmt`   | Formata o código-fonte utilizando `go fmt`.                |
| `make clean` | Remove artefatos de build, como a pasta `bin/`.            |

---

## 🛠️ Como Executar Manualmente (Sem Makefile)

### 1. Inicie o servidor (porta 8080):

```bash
go run cmd/main.go
```

_(O servidor ficará escutando o endpoint `/cotacao`)_

---

### 2. Execute o cliente (em outro terminal):

```bash
go run pkg/client/client.go
```

_(O cliente fará a requisição e salvará o resultado no arquivo `cotacao.txt`)_

---

## ⏱️ Comportamento de Timeout

| Contexto                              | Timeout | Comportamento em caso de falha |
| ------------------------------------- | ------- | ------------------------------ |
| Requisição à API externa (AwesomeAPI) | 200ms   | Log de erro em caso de timeout |
| Gravação no banco (SQLite)            | 10ms    | Log de erro em caso de timeout |
| Requisição HTTP do cliente            | 300ms   | Log de erro em caso de timeout |

Todos os erros (incluindo timeouts) são exibidos nos logs do console.

---

## ✅ Exemplo de Saída no Arquivo `cotacao.txt`

```
Dólar: 5.3312
```

_(O valor varia de acordo com a cotação no momento da execução)_

---
