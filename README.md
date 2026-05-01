# Micro-API de Tarefas com Priorizacao Assistida por IA

API REST em Go (Fiber) para gestao de tarefas com sugestao de prioridade por heuristica local e integracao opcional com LLM.

## Objetivo

- Criar, listar, atualizar e excluir tarefas
- Padronizar status e niveis de prioridade
- Sugerir prioridade automaticamente com fallback seguro

## Stack

- Go `1.25.x`
- Fiber `v2`
- Arquitetura em camadas: `api -> services -> repositories`
- Persistencia em memoria (map com mutex)

## Estrutura do Projeto

```text
app/
  api/                # Handlers e rotas HTTP
  models/             # Entidades e DTOs
  repositories/       # Persistencia em memoria
  services/           # Regras de negocio + PriorityAdvisor
  main.go             # Bootstrap HTTP
docs/
tests/
```

## Reproducao Rapida (Maquina Limpa)

### Pre-requisitos

- Go `1.25.x` instalado (`go version`)
- Git
- PowerShell (Windows)

### Passo a passo

```powershell
git clone <url-do-repositorio>
cd "Labotatorio Projeto"
copy .env.example .env
make install
make run
```

Em outro terminal:

```powershell
curl http://localhost:8080/health
curl http://localhost:8080/tasks
```

## Configuracao

Arquivo recomendado: `.env` (baseado em `.env.example`).

Variaveis:

- `OPENAI_API_KEY`: chave da API OpenAI (opcional)
- `OPENAI_BASE_URL`: default `https://api.openai.com/v1`
- `OPENAI_MODEL`: default `gpt-4.1-mini`
- `PRIORITY_ADVISOR_TIMEOUT`: default `4s`

Observacao:

- Sem `OPENAI_API_KEY`, a API funciona normalmente usando apenas heuristica local.

## Comandos de Desenvolvimento

```powershell
make install
make run
make test
```

Equivalentes diretos:

```powershell
go mod download
go run ./app
go test ./...
```

## Endpoints

Base URL local: `http://localhost:8080`

### Healthcheck

- `GET /health`
- `200 OK`

Exemplo de resposta:

```json
{
  "status": "ok",
  "timestamp": "2026-05-01T00:00:00Z"
}
```

### Tasks

- `POST /tasks`
- `GET /tasks`
- `GET /tasks/:id`
- `PUT /tasks/:id`
- `DELETE /tasks/:id`

Status esperados:

- Criacao: `201`
- Leitura/listagem: `200`
- Exclusao: `204`
- Nao encontrado: `404`
- JSON invalido: `400`

Exemplo de criacao:

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d "{\"title\":\"Corrigir bug de login\",\"description\":\"Erro 500 em producao\"}"
```

Exemplo de listagem:

```bash
curl http://localhost:8080/tasks
```

## Uso da IA na Priorizacao

`PriorityAdvisor` opera em 2 modos:

1. Heuristica local (sempre disponivel)
2. LLM opcional (quando `OPENAI_API_KEY` esta configurada)

### Contrato esperado do LLM

Resposta de conteudo JSON com:

```json
{"priority":"low|medium|high|critic"}
```

### Fallback Seguro

Se a chamada externa falhar, a API nao quebra o CRUD e aplica heuristica local nos casos:

- timeout
- HTTP nao-2xx
- JSON invalido
- `choices` vazio
- prioridade fora do contrato

### Exemplos de Heuristica

- `"incidente em producao"` -> tende a `critic`
- `"urgente para cliente"` -> tende a `high`
- `"melhoria de texto"` -> tende a `medium`
- `"item opcional de backlog"` -> tende a `low`

### Custo e Latencia

- Heuristica local: custo externo zero, latencia minima
- LLM: maior latencia/custo por chamada; recomendado quando precisa de contexto semantico adicional

### Seguranca

- Nunca versionar `OPENAI_API_KEY`
- Use `.env` local
- Em ausencia de credencial, comportamento padrao continua funcional

## Testes

Executar suite:

```powershell
make test
```

Ou:

```powershell
go test ./...
```

Cobertura atual inclui cenarios de:

- rotas (`2xx`, `4xx`)
- servico (CRUD + casos de erro)
- repositorio (ordenacao e atualizacao)
- priority advisor (heuristica, sucesso LLM e fallback)

## Limitacoes do MVP

- Persistencia apenas em memoria (sem dados apos restart)
- Sem autenticacao/autorizacao
- Sem migracoes de banco
- Sem observabilidade completa (metricas/tracing)

## Licenca

Definir conforme politica do projeto.
