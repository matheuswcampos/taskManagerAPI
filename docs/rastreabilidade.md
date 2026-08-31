# Matriz de Rastreabilidade de Requisitos

Esta matriz relaciona os requisitos do escopo e do backlog com os artefatos que os especificam, os componentes que os implementam e as evidencias de verificacao existentes.

## Legenda

- **Implementado:** ha codigo e teste automatizado associados.
- **Parcial:** ha implementacao ou documentacao, mas ainda existe lacuna relevante.
- **Pendente:** item planejado sem implementacao principal.
- **Evidencia:** arquivo usado para confirmar especificacao, implementacao ou teste.

## Requisitos Funcionais

| ID | Requisito | Fonte | Implementacao | Evidencia de teste | Status | Observacoes |
|---|---|---|---|---|---|---|
| RF-001 | Health check da API | `docs/escopo-mvp.md`, `docs/backlog.md`, `README.md` | `app/main.go` | `app/main_test.go` | Implementado | `GET /health` retorna `200`, `status: "ok"` e `timestamp`. |
| RF-002 | Cadastro de tarefa | `docs/escopo-mvp.md`, `docs/backlog.md`, `README.md` | `app/api/task_routes.go`, `app/services/task_service.go`, `app/models/task.go`, `app/repositories/task_repository.go` | `tests/task_routes_test.go`, `tests/task_service_test.go` | Parcial | Criacao funciona; faltam validacoes verificadas para `title` ausente/vazio e valores invalidos de dominio. |
| RF-003 | Listagem e consulta de tarefas | `docs/escopo-mvp.md`, `docs/backlog.md`, `README.md` | `app/api/task_routes.go`, `app/services/task_service.go`, `app/repositories/task_repository.go` | `tests/task_routes_test.go`, `tests/task_service_test.go`, `tests/task_repository_test.go` | Implementado | Inclui `GET /tasks` e `GET /tasks/{id}`. Filtros nao fazem parte do MVP atual. |
| RF-004 | Atualizacao parcial de tarefa | `docs/escopo-mvp.md`, `docs/backlog.md`, `README.md` | `app/api/task_routes.go`, `app/services/task_service.go`, `app/models/task.go`, `app/repositories/task_repository.go` | `tests/task_routes_test.go`, `tests/task_service_test.go`, `tests/task_repository_test.go` | Parcial | Atualiza campos parciais e trata ID inexistente; validacoes de dominio devem ser implementadas/testadas. |
| RF-005 | Remocao de tarefa | `docs/escopo-mvp.md`, `docs/backlog.md`, `README.md` | `app/api/task_routes.go`, `app/services/task_service.go`, `app/repositories/task_repository.go` | `tests/task_routes_test.go`, `tests/task_service_test.go` | Implementado | `DELETE /tasks/{id}` retorna `204` em sucesso e `404` para ID inexistente. |
| RF-006 | Priorizacao assistida por IA | `docs/escopo-mvp.md`, `docs/backlog.md`, `README.md` | `app/services/priority_advisor.go`, `app/services/task_service.go`, `app/models/task.go` | `tests/priority_advisor_test.go`, `tests/task_service_test.go` | Implementado | Usa LLM quando configurado e heuristica local como fallback; nao ha endpoint dedicado no contrato atual. |
| RF-007 | Contrato de resposta JSON | `docs/escopo-mvp.md`, `docs/backlog.md`, `README.md` | `app/api/task_routes.go`, `app/models/task.go` | `tests/task_routes_test.go`, `app/main_test.go` | Parcial | Contrato minimo com JSON e campo `error` existe; schema unico `code/message/details` permanece evolucao planejada. |

## Requisitos Tecnicos e Nao Funcionais

| ID | Requisito | Fonte | Implementacao / Artefato | Evidencia | Status | Observacoes |
|---|---|---|---|---|---|---|
| RT-001 / RNF-006 | Estrutura tecnica por camadas | `docs/backlog.md`, `docs/escopo-mvp.md`, `docs/arquitetura-componentes.mmd` | `app/api`, `app/services`, `app/repositories`, `app/models` | Testes por camada em `tests` e `app/main_test.go` | Implementado | A separacao esta refletida em pacotes e interfaces. |
| RT-002 / RNF-007 | Testes unitarios automatizados | `docs/backlog.md`, `docs/escopo-mvp.md` | `tests/task_service_test.go`, `tests/task_repository_test.go`, `tests/priority_advisor_test.go` | `go test ./...` | Implementado | Cobre servico, repositorio e prioridade. |
| RT-003 / RNF-007 | Testes de endpoint | `docs/backlog.md`, `docs/escopo-mvp.md` | `tests/task_routes_test.go`, `app/main_test.go` | `go test ./...` | Implementado | Cobre status HTTP e payloads principais. |
| RT-004 / RNF-005 | Padronizacao de erros e logs estruturados | `docs/backlog.md`, `docs/escopo-mvp.md` | `app/api/task_routes.go`, `app/services/priority_advisor.go` | Testes parciais de erro HTTP | Pendente | Erro unico `code/message/details` e logs por requisicao ainda nao estao implementados. |
| RT-005 / RNF-006 | Qualidade estatica | `docs/backlog.md`, `docs/escopo-mvp.md` | Nao identificado | Nao identificado | Pendente | Falta configuracao de lint e pipeline; `go fmt` pode ser executado localmente. |
| RT-006 | Persistencia em banco | `docs/backlog.md`, `README.md`, `docs/riscos/*` | `app/repositories/task_repository.go` usa memoria | `tests/task_repository_test.go` | Pendente | PostgreSQL esta planejado; persistencia atual nao sobrevive a restart. |
| RT-007 | Documentacao tecnica da API | `docs/backlog.md`, `README.md` | `README.md` | Nao ha OpenAPI | Pendente | Falta especificacao OpenAPI versionada. |
| RT-008 | Prontidao de entrega reproduzivel | `docs/backlog.md`, `README.md` | `README.md`, `Makefile`, `.env.example` | `make test`, `go test ./...` | Parcial | Falta pipeline e validacao formal de build/testes para entrega. |
| RT-009 | Pronto para uso interno | `docs/backlog.md`, `README.md` | Multiplos componentes | Suite `go test ./...` | Pendente | Depende de estabilidade, ausencia de erros criticos e registro de release. |
| RNF-001 | Stack e compatibilidade | `docs/escopo-mvp.md`, `README.md`, `go.mod` | Projeto Go/Fiber | Build/testes locais | Parcial | README e `go.mod` devem permanecer alinhados. |
| RNF-002 | Desempenho inicial | `docs/escopo-mvp.md` | Implementacao em memoria | Nao ha teste de carga | Pendente | Meta p95 <= 300 ms em carga baixa nao possui medicao automatizada. |
| RNF-003 | Confiabilidade | `docs/escopo-mvp.md`, `docs/riscos/*` | Tratamento de erros em rotas, fallback do advisor | Testes de erro e fallback | Parcial | Casos de payload invalido e dominio devem ser ampliados. |
| RNF-004 | Seguranca basica | `docs/escopo-mvp.md`, `README.md`, `.env.example`, `.gitignore` | Configuracao via ambiente, chave LLM opcional | Revisao documental | Parcial | Sem autenticacao no MVP; deve permanecer restrito a ambiente controlado. |
| RNF-005 | Observabilidade minima | `docs/escopo-mvp.md`, `docs/backlog.md` | Logs do `PriorityAdvisor` | Nao ha teste especifico | Parcial | Logs por requisicao, metricas e tracing nao fazem parte da implementacao atual. |
| RNF-006 | Qualidade e manutenibilidade | `docs/escopo-mvp.md`, `docs/backlog.md` | Camadas, testes, Makefile | `go test ./...` | Parcial | Lint e pipeline ainda pendentes. |
| RNF-007 | Testabilidade | `docs/escopo-mvp.md`, `docs/backlog.md` | Testes automatizados | `go test ./...` | Implementado | Suite cobre fluxos principais e fallback de prioridade. |

## Pontos de Atencao de Rastreabilidade

- IDs funcionais foram normalizados para `RF-001` a `RF-007`.
- `RT-009 - Pronto para uso interno` substitui o antigo uso de `RF-007` para evitar confundir criterio de entrega com requisito funcional.
- O valor `canceled` fica fora do contrato atual do MVP; se for necessario, deve virar novo requisito ou mudanca aprovada.
- A obrigatoriedade de `title` e a rejeicao de `status`/`priority` invalidos estao especificadas, mas ainda precisam de teste e confirmacao de implementacao.
- A documentacao OpenAPI ainda nao existe, apesar de prevista em `RT-007`.
