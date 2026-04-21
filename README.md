# Micro-API de Gestão de Tarefas com Priorização Assistida por IA

## Objetivo
Construir um MVP de micro-API para cadastro e gestão de tarefas, com suporte a priorização assistida por IA para ajudar na organização do backlog.

A API terá foco em:
- criação, listagem, atualização e remoção de tarefas
- classificação de prioridade (`baixa`, `média`, `alta`, `crítica`)
- sugestão de prioridade com base no contexto da tarefa via modelo de IA
- evoluir gradualmente para um fluxo de priorização assistida por IA em produção

## Stack
- Go 1.25+
- API HTTP REST em Go (estrutura pronta para uso com `net/http` ou framework leve como Gin)
- JSON como formato de entrada/saída
- Variáveis de ambiente para configuração (`.env`)
- Integração com provedor de IA via API key
- Testes com `go test`

## Como rodar localmente

### 1. Pré-requisitos
- Go instalado (`go version`)
- Git
- PowerShell (Windows)

### 2. Ativar ambiente local do projeto
Na raiz do repositório:

```powershell
.\activate-go-env.ps1
```

### 3. Inicializar módulo Go (se ainda não existir)

```powershell
go mod init task-prioritization-api
```

### 4. Criar arquivo de variáveis de ambiente
Crie um arquivo `.env` com valores iniciais:

```env
APP_ENV=development
APP_PORT=8080
AI_API_KEY=your_api_key_here
AI_MODEL=gpt-4.1-mini
```

### 5. Executar a API
Quando o `main.go` estiver criado, execute:

```powershell
go run ./...
```

### 6. Executar testes

```powershell
go test ./...
```

## Roadmap de Releases

### v0.1.0 - Base da API
- estrutura inicial do projeto em Go
- endpoint de health check (`GET /health`)
- CRUD básico de tarefas em memória
- validação mínima de payload

### v0.2.0 - Priorização Assistida por IA
- endpoint para sugestão de prioridade
- prompt inicial para classificar urgência/impacto
- fallback manual quando IA estiver indisponível
- logs básicos de requisição e erro

### v0.3.0 - Persistência e Qualidade
- persistência em banco (PostgreSQL)
- migrações e repositório de dados
- testes unitários e de integração para fluxos principais
- documentação OpenAPI/Swagger

### v1.0.0 - MVP de Produção
- autenticação simples por token
- observabilidade mínima (logs estruturados + métricas)
- hardening de erros e timeouts
- versionamento estável de API (`/v1`)

## Status
Projeto em fase inicial de estruturação do MVP.
