# Especificacao de Requisitos Verificaveis

Este artefato consolida os requisitos do MVP em formato verificavel. Ele nao substitui o escopo, o backlog ou a matriz de rastreabilidade; seu papel e deixar claro o comportamento esperado, o criterio objetivo de verificacao, a evidencia existente e as lacunas que ainda precisam ser implementadas ou testadas.

## Criterios de Verificabilidade

Um requisito e considerado verificavel quando possui:

- ID unico.
- Comportamento esperado sem ambiguidade.
- Entrada, saida ou condicao observavel.
- Criterio de aceite passavel/reprovavel.
- Evidencia de teste automatizado, teste manual ou inspecao documental.
- Status atualizado.

## Requisitos Funcionais Verificaveis

| ID | Requisito | Condicao de entrada | Resultado esperado | Evidencia atual | Status |
|---|---|---|---|---|---|
| RF-001 | Health check | `GET /health` | HTTP `200`, JSON com `status: "ok"` e `timestamp` UTC | `app/main_test.go` | Implementado |
| RF-002.1 | Criar tarefa com payload valido | `POST /tasks` com `title` valido | HTTP `201`, tarefa com `id`, `title`, `status`, `priority`, `created_at`, `updated_at` | `tests/task_routes_test.go`, `tests/task_service_test.go` | Implementado |
| RF-002.2 | Aplicar status padrao | Criacao sem `status` | Tarefa criada com `status: "todo"` | `tests/task_service_test.go` | Implementado |
| RF-002.3 | Aplicar prioridade sugerida | Criacao sem `priority` | Tarefa criada com `priority` em `low`, `medium`, `high` ou `critic` | `tests/task_service_test.go`, `tests/priority_advisor_test.go` | Implementado |
| RF-002.4 | Rejeitar JSON invalido na criacao | `POST /tasks` com JSON malformado | HTTP `400`, JSON com campo `error` | `tests/task_routes_test.go` | Implementado |
| RF-002.5 | Rejeitar titulo ausente ou vazio na criacao | `POST /tasks` sem `title` ou com `title` vazio | HTTP `400`, JSON com campo `error` | Nao identificado | Pendente |
| RF-002.6 | Rejeitar dominio invalido na criacao | `status` fora de `todo|doing|done` ou `priority` fora de `low|medium|high|critic` | HTTP `400`, JSON com campo `error` | Nao identificado | Pendente |
| RF-003.1 | Listar tarefas | `GET /tasks` | HTTP `200`, array JSON com tarefas do repositorio atual | `tests/task_routes_test.go`, `tests/task_service_test.go` | Implementado |
| RF-003.2 | Listar sem tarefas | `GET /tasks` sem registros | HTTP `200`, array JSON vazio | Nao identificado de forma explicita | Parcial |
| RF-003.3 | Consultar tarefa existente | `GET /tasks/{id}` com ID existente | HTTP `200`, JSON da tarefa | `tests/task_routes_test.go` | Implementado |
| RF-003.4 | Consultar tarefa inexistente | `GET /tasks/{id}` com ID inexistente | HTTP `404`, JSON com campo `error` | `tests/task_routes_test.go`, `tests/task_service_test.go` | Implementado |
| RF-004.1 | Atualizar tarefa parcialmente | `PUT /tasks/{id}` com payload valido | HTTP `200`, somente campos informados alterados | `tests/task_routes_test.go`, `tests/task_service_test.go` | Implementado |
| RF-004.2 | Recalcular prioridade ao alterar texto sem prioridade explicita | `PUT /tasks/{id}` alterando `title` ou `description` sem `priority` | Prioridade recalculada quando advisor retorna sugestao valida | `tests/task_service_test.go` | Implementado |
| RF-004.3 | Preservar prioridade explicita na atualizacao | `PUT /tasks/{id}` com `priority` informada | Prioridade final igual ao valor informado | `tests/task_service_test.go` | Implementado |
| RF-004.4 | Rejeitar JSON invalido na atualizacao | `PUT /tasks/{id}` com JSON malformado | HTTP `400`, JSON com campo `error` | `tests/task_routes_test.go` | Implementado |
| RF-004.5 | Retornar erro para atualizacao de ID inexistente | `PUT /tasks/{id}` com ID inexistente | HTTP `404`, JSON com campo `error` | `tests/task_service_test.go` | Parcial |
| RF-004.6 | Rejeitar titulo vazio ou dominio invalido na atualizacao | `title` vazio, `status` invalido ou `priority` invalida | HTTP `400`, JSON com campo `error` | Nao identificado | Pendente |
| RF-005.1 | Remover tarefa existente | `DELETE /tasks/{id}` com ID existente | HTTP `204`, sem corpo obrigatorio | `tests/task_routes_test.go`, `tests/task_service_test.go` | Implementado |
| RF-005.2 | Remover tarefa inexistente | `DELETE /tasks/{id}` com ID inexistente | HTTP `404`, JSON com campo `error` | `tests/task_routes_test.go`, `tests/task_service_test.go` | Implementado |
| RF-006.1 | Usar heuristica local sem chave de LLM | Criacao/atualizacao sem `priority` e sem `OPENAI_API_KEY` | Prioridade sugerida pela heuristica | `tests/priority_advisor_test.go` | Implementado |
| RF-006.2 | Usar LLM quando resposta for valida | LLM habilitado e resposta JSON valida | Prioridade sugerida pelo LLM | `tests/priority_advisor_test.go` | Implementado |
| RF-006.3 | Aplicar fallback quando LLM falhar | Timeout, HTTP nao-2xx, JSON invalido, resposta vazia ou prioridade invalida | CRUD continua e prioridade vem da heuristica | `tests/priority_advisor_test.go` | Implementado |
| RF-007.1 | Retornar JSON em sucesso | Operacoes de leitura/criacao/atualizacao com sucesso | Payload JSON compativel com modelo de tarefa | `tests/task_routes_test.go`, `app/main_test.go` | Implementado |
| RF-007.2 | Retornar JSON em erro | Erros de payload invalido ou recurso inexistente | JSON com campo `error` | `tests/task_routes_test.go` | Parcial |
| RF-007.3 | Manter schema unico de erro | Erros da API | JSON com `code`, `message`, `details` | Nao identificado | Pendente como evolucao |

## Requisitos Nao Funcionais Verificaveis

| ID | Requisito | Criterio verificavel | Evidencia atual | Status |
|---|---|---|---|---|
| RNF-001 | Stack e compatibilidade | Projeto compila/testa com Go e Fiber documentados | `go.mod`, `README.md`, suite de testes | Parcial |
| RNF-002 | Desempenho inicial | CRUD com p95 <= 300 ms em carga baixa; `/health` sem dependencia externa | Nao ha teste de carga | Pendente |
| RNF-003 | Confiabilidade | Erros de requisicao e falhas de LLM nao encerram o processo | Testes de rota e fallback | Parcial |
| RNF-004 | Seguranca basica | Segredos fora do repositorio; configuracao por ambiente | `.env.example`, `.gitignore`, `README.md` | Parcial |
| RNF-005 | Observabilidade minima | Falhas de IA logadas; logs por requisicao com metodo, rota, status e duracao | Logs do `PriorityAdvisor` | Parcial |
| RNF-006 | Qualidade e manutenibilidade | Camadas separadas; `go fmt` e lint sem pendencias antes de release | Estrutura de pacotes e testes | Parcial |
| RNF-007 | Testabilidade | `go test ./...` executa suite automatizada dos fluxos principais | Testes em `tests` e `app/main_test.go` | Implementado |

## Decisoes de Contrato

- IDs funcionais oficiais: `RF-001` a `RF-007`.
- `RF-003` cobre listagem e consulta por ID.
- `RF-007` cobre o contrato JSON atual; erro padronizado com `code/message/details` e evolucao tecnica, nao obrigacao do MVP atual.
- `RT-009` representa prontidao de uso interno e substitui o antigo uso de `RF-007` no backlog.
- Valores oficiais de `status`: `todo`, `doing`, `done`.
- `status: "canceled"` fica fora do contrato atual.
- Valores oficiais de `priority`: `low`, `medium`, `high`, `critic`.
- Valores invalidos de `status` ou `priority` devem retornar HTTP `400`, mas a implementacao/teste ainda precisa ser confirmada.

## Lacunas Prioritarias

| Lacuna | Requisito afetado | Acao necessaria |
|---|---|---|
| Titulo ausente ou vazio sem evidencia de rejeicao | RF-002, RF-004 | Implementar/confirmar validacao e adicionar testes de rota/servico. |
| Valores invalidos de `status` e `priority` sem evidencia de rejeicao | RF-002, RF-004 | Implementar/confirmar validacao e adicionar testes para HTTP `400`. |
| Array vazio em listagem sem teste explicito | RF-003 | Adicionar teste de `GET /tasks` sem registros. |
| Atualizacao de ID inexistente sem teste de rota identificado | RF-004 | Adicionar teste HTTP para `PUT /tasks/{id}` inexistente. |
| Schema OpenAPI ausente | RF-007, RT-007 | Criar `docs/api/openapi.yaml`. |
| Desempenho e observabilidade sem evidencia objetiva | RNF-002, RNF-005 | Definir medicao e criterios minimos verificaveis. |
