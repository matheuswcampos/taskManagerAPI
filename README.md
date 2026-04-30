# Micro-API de Tarefas com Priorizacao Assistida por IA

MVP de uma API REST em Go para gestão de tarefas com sugestão de prioridade baseada em heurística local e opcionalmente em LLM.

## Objetivo

Entregar uma base enxuta para:

- criar, listar, atualizar e excluir tarefas
- padronizar estados e niveis de prioridade
- sugerir prioridade automaticamente com fallback seguro
- evoluir para persistência e operação em ambiente de produção

## Stack

- Go `1.25+`
- Fiber `v2`
- Arquitetura em camadas: `api -> services -> repositories`
- Persistência atual em memória (map protegido por mutex)
- Testes com `go test` (unitários e rotas)

## Estrutura do Projeto

```text
app/
  api/                # Handlers e rotas HTTP
  models/             # Entidades e DTOs
  repositories/       # Persistência em memória
  services/           # Regras de negócio e PriorityAdvisor
  main.go             # Bootstrap HTTP atual (healthcheck)
docs/
  arquitetura-componentes.mmd
  backlog.md
  escopo-mvp.md
tests/
  *_test.go
```

## Instalação

### Pre-requisitos

- Go instalado (`go version`)
- Git
- PowerShell (Windows)

### Passos

```powershell
git clone <url-do-repositorio>
cd "Labotatorio Projeto"
go mod download
```

Opcional (helper local do projeto):

```powershell
.\activate-go-env.ps1
```

## Configuração

A aplicação usa variáveis de ambiente para o módulo de priorização por IA:

- `OPENAI_API_KEY`: chave de API (opcional)
- `OPENAI_BASE_URL`: endpoint base (padrao: `https://api.openai.com/v1`)
- `OPENAI_MODEL`: modelo (padrão: `gpt-4.1-mini`)
- `PRIORITY_ADVISOR_TIMEOUT`: timeout HTTP (padrão: `4s`)

Exemplo (PowerShell):

```powershell
$env:OPENAI_API_KEY="sua_chave"
$env:OPENAI_MODEL="gpt-4.1-mini"
$env:PRIORITY_ADVISOR_TIMEOUT="4s"
```

## Execução

Subir a API:

```powershell
go run ./app
```

Endpoint disponivel no bootstrap atual:

- `GET /health` -> `200 OK`

Exemplo:

```powershell
curl http://localhost:8080/health
```

Resposta esperada:

```json
{
  "status": "ok",
  "timestamp": "2026-04-30T12:00:00Z"
}
```

## Testes

Executar toda a suite:

```powershell
go test ./...
```

Cobertura atual de testes inclui:

- `TaskService` (CRUD + casos de erro)
- `PriorityAdvisor` (heurística e fallback)
- rotas `/tasks` com validação de status (`201`, `200`, `204`, `404`)

## Arquitetura

Separação principal:

- `app/api`: traduz HTTP <-> DTO
- `app/services`: aplica regras de negócio
- `app/repositories`: abstrai armazenamento
- `PriorityAdvisor`: encapsula lógica de sugestão de prioridade

Fluxo de alto nivel:

1. Cliente envia requisição HTTP.
2. Handler faz parse e chama `TaskService`.
3. `TaskService` aplica regras (defaults, sugestão de prioridade).
4. Repositório persiste/consulta tarefa.
5. Handler retorna resposta JSON.

Diagrama Mermaid em: `docs/arquitetura-componentes.mmd`.

## Uso da IA na Priorizacao

O componente `PriorityAdvisor` opera em dois modos:

1. **Heurística local (sempre disponivel)**
- analisa termos de urgência/impacto
- classifica em `low`, `medium`, `high` ou `critic`

2. **LLM opcional (quando `OPENAI_API_KEY` existe)**
- envia contexto (`title`, `description`) para o modelo configurado
- espera JSON com campo `priority`
- valida retorno e aplica somente valores permitidos

Fallback de segurança:

- se chamada externa falhar (timeout, erro HTTP, payload invalido), retorna heurística local sem quebrar o fluxo do CRUD.

## Endpoints de Tarefas

As rotas CRUD estão implementadas em `app/api/task_routes.go`:

- `POST /tasks`
- `GET /tasks`
- `GET /tasks/:id`
- `PUT /tasks/:id`
- `DELETE /tasks/:id`

Status esperados no contrato atual:

- criação: `201`
- leitura/listagem: `200`
- exclusão: `204`
- não encontrado: `404`

Observação: o `main.go` atual registra apenas `/health`. O registro das rotas de tarefas já existe e pode ser integrado no bootstrap sem alterar as camadas de dominio.

## Limitações do MVP

- armazenamento em memória (sem persistencia entre reinicios)
- sem autenticação/autorização
- sem migrações de banco
- sem validação estruturada de payload com retorno detalhado por campo
- sem observabilidade completa (métricas/tracing)
- bootstrap atual ainda não expõe `/tasks` em runtime por padrão

## Próximos Passos

1. Integrar `TaskRepository` e `TaskService` no `main.go` e publicar `/tasks` no runtime principal.
2. Adicionar validação de entrada com mensagens consistentes por campo.
3. Introduzir persistência real (PostgreSQL + migracoes).
4. Incluir testes de integração com banco e cenários de concorrência.
5. Adicionar autenticação simples e versionamento de API (`/v1`).
6. Evoluir observabilidade (logs estruturados, metricas, tracing).

## Licença

Definir conforme politica do projeto.