# Nome do projeto
PROJECT_NAME = boilerplate_go

# Comandos
GO = go
GOTEST = go test
GOBUILD = go build
GORUN = go run
GOCLEAN = go clean
GOFMT = gofmt -w
ENTRY_DIR = ./cmd/main.go # Ponto de entrada do projeto

# Variáveis para testes
TEST_DIRS = ./... # Diretórios onde os testes estão localizados
BENCH_FILTER = .  # Padrão de filtro para benchmarks, use com '-bench=<regex>'
BENCH_MEM = -benchmem # Adicionar análise de memória aos benchmarks

# Alvos do Makefile
.PHONY: all build run test clean format benchmark test-with-doubles

# Alvo padrão
all: build

# Compila o projeto
build:
	$(GOBUILD) -o $(PROJECT_NAME)

# Executa o projeto
run:
	$(GORUN) .

# Executa os testes unitários normais
test:
	$(GOTEST) $(TEST_DIRS)

# Executa os benchmarks
benchmark:
	$(GOTEST) -bench=$(BENCH_FILTER) $(BENCH_MEM) $(TEST_DIRS)

# Gera o arquivo de cobertura dos testes
coverage:
	$(GOTEST) --coverprofile=coverage.out $(TEST_DIRS)
	@echo "Cobertura gerada em coverage.out"
	@echo "Para visualizar a cobertura execute:"
	@echo "go tool cover -html=coverage.out -o coverage.html"

# FIXME: go tool cover -html=coverage.out -o coverage.html

# Limpa os arquivos gerados pelo Go
clean:
	$(GOCLEAN)

# Formata o código usando gofmt
format:
	$(GOFMT) .

# Testes com environment de testes-doubles
test-with-doubles:
	@echo "Executando testes com Test Doubles"
	GO_ENV=test $(GOTEST) $(TEST_DIRS)
