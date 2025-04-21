# GoAuth

Um serviço de autenticação robusto desenvolvido em Go, utilizando PostgreSQL como banco de dados.

## 🚀 Tecnologias

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## 📋 Pré-requisitos

- Go 1.20+
- Docker
- Docker Compose

## 🔧 Configuração do Ambiente

1. Clone o repositório:

```bash
git clone https://github.com/luansantosco/goauth.git
cd GoAuth
```

2. Crie um arquivo `.env` na raiz do projeto:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=goauth_user
DB_PASSWORD=think
DB_NAME=goauth
JWT_SECRET=sua_chave
```

3. Inicie o container do PostgreSQL:

```bash
docker-compose up -d
```

4. Execute a aplicação:

```bash
go run cmd/server/main.go
```

## 🐳 Docker

O projeto utiliza Docker para facilitar o desenvolvimento. O arquivo `docker-compose.yml` configura:

- Container PostgreSQL 15
- Persistência de dados via volumes
- Variáveis de ambiente para conexão segura

## 🛠️ Estrutura do Projeto

```bash
GoAuth/
├── cmd/
│   └── server/
│       └── main.go            # Ponto de entrada da aplicação
├── configs/
│   ├── config.go              # Carregar variáveis de ambiente
│   └── db.go                  # Conexão com o banco de dados
├── internal/
│   ├── controllers/           # Controllers HTTP
│   ├── services/              # Regras de negócio
│   ├── repositories/          # Acesso a banco de dados
│   ├── models/                # Definições de dados
│   ├── middlewares/           # Middlewares
│   └── utils/                 # Utilitários (JWT, Hash, etc.)
├── docker-compose.yml         # Configuração do Docker
├── go.mod                     # Módulo Go
├── go.sum                     # Checksum das dependências
├── .env                       # Variáveis de ambiente
└── README.md                  # Documentação
```

## ✅ Status do Projeto

- ✅ Estrutura modular (cmd, configs, internal)
- ✅ Docker configurado
- ✅ Banco de dados PostgreSQL funcionando
- ✅ Serviço de registro de usuários funcionando
- ✅ Banco de dados conectado com sucesso

## 📝 Como contribuir

1. Faça um Fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'feat: add AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👤 Autor

**Luan Rodrigues**

- GitHub: [@luansantosco](https://github.com/luansantosco)
- LinkedIn: [Luan Santos](https://www.linkedin.com/in/luansantosco/)
