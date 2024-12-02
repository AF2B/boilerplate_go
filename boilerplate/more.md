boilerplate/
├── cmd/
│   └── main.go                  # Entrada principal
├── config/
│   ├── config.go                # Configuração de envs
├── pkg/
│   ├── logger/
│   │   └── zap.go               # Configuração do Zap
│   └── database/
│       └── postgres.go          # Configuração do Postgres
├── internal/
│   ├── routes/
│   │   └── router.go            # Configuração de rotas do Gin
│   ├── handlers/
│   │   └── health.go            # Handler de exemplo
│   ├── middleware/
│   │   └── logger.go            # Middleware de Logging
│   └── models/
│       └── base.go              # Model base e estrutura
├── tests/
│   ├── mocks/
│   │   └── mock_user_repo.go    # Mocks para testes
│   ├── stubs/
│   │   └── stub_user_repo.go    # Stubs para testes
│   ├── fakes/
│   │   └── fake_user_repo.go    # Fakes para testes
│   └── helpers/
│       └── test_helpers.go      # Funções auxiliares para testes
├── .env                         # Variáveis de ambiente
├── go.mod                       # Dependências Go
├── go.sum                       # Checksum das dependências
├── Makefile                     # Automação de comandos
