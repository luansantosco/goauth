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
git clone https://github.com/luansantosco/goauth
cd GoAuth
```

2. Crie um arquivo `.env` na raiz do projeto:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=goauth_user
DB_PASSWORD=think
DB_NAME=goauth
DATABASE_URL=postgres://goauth_user:think@localhost:5432/goauth?sslmode=disable
```

3. Inicie o container do PostgreSQL:
```bash
docker-compose up -d
```

4. Execute a aplicação:
```bash
go run main.go
```

## 🐳 Docker

O projeto utiliza Docker para facilitar o desenvolvimento. O arquivo `docker-compose.yml` configura:

- Container PostgreSQL 15
- Persistência de dados via volumes
- Variáveis de ambiente para conexão segura

## ✅ Status do Projeto

- ✅ Docker configurado
- ✅ Banco PostgreSQL funcionando
- ✅ Usuário goauth_user criado
- ✅ Banco goauth criado
- ✅ Projeto GoAuth conectando ao banco

## 🛠️ Estrutura do Projeto

```
GoAuth/
├── db/
│   └── db.go              # Configuração da conexão com banco
├── .env                   # Variáveis de ambiente
├── docker-compose.yml     # Configuração do Docker
├── go.mod
├── go.sum
└── main.go                # Entrada da aplicação
```

## 📝 Contribuição

1. Faça um Fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👤 Autor

Luan Rodrigues

## 🔗 Contato

- GitHub: [luanrodrigues](https://github.com/luansantosco)
- LinkedIn: [Luan Santos](https://www.linkedin.com/in/luansantosco/)
