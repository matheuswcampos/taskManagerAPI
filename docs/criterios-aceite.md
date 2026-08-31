# Criterios de Aceite

Este documento consolida criterios de aceite dos requisitos funcionais, tecnicos e nao funcionais do MVP. Os criterios aqui descritos devem orientar testes automatizados, validacao manual e futuras revisoes de escopo.

## Requisitos Funcionais

### RF-001 - Health Check

- Dado que a API esta em execucao, quando `GET /health` for chamado, entao deve retornar HTTP `200`.
- A resposta deve ser JSON.
- O campo `status` deve ter valor `ok`.
- O campo `timestamp` deve estar presente e representar data/hora em UTC.
- O endpoint nao deve depender de servicos externos.

### RF-002 - Cadastro de Tarefa

- Dado um payload JSON valido com `title`, quando `POST /tasks` for chamado, entao deve retornar HTTP `201`.
- A resposta deve conter `id`, `title`, `status`, `priority`, `created_at` e `updated_at`.
- `description` deve ser aceito como campo opcional.
- Quando `status` nao for informado, a tarefa deve ser criada com `status: "todo"`.
- Quando `priority` nao for informada, a API deve tentar aplicar prioridade sugerida.
- Quando o JSON estiver malformado, a API deve retornar HTTP `400`.
- Quando `title` estiver ausente ou vazio, a API deve retornar HTTP `400`.
- Valores aceitos para `status` no contrato atual: `todo`, `doing`, `done`.
- Valores aceitos para `priority`: `low`, `medium`, `high`, `critic`.
- Valores invalidos de `status` ou `priority` devem retornar HTTP `400`.

### RF-003 - Listagem e Consulta de Tarefas

- Dado que a API esta em execucao, quando `GET /tasks` for chamado, entao deve retornar HTTP `200`.
- A resposta de `GET /tasks` deve ser um array JSON.
- Todas as tarefas cadastradas no repositorio atual devem ser retornadas.
- Quando nao houver tarefas, a resposta deve ser um array vazio.
- Dado um ID existente, quando `GET /tasks/{id}` for chamado, entao deve retornar HTTP `200`.
- A resposta de `GET /tasks/{id}` deve conter os dados da tarefa correspondente.
- Dado um ID inexistente, quando `GET /tasks/{id}` for chamado, entao deve retornar HTTP `404`.
- Em erro de nao encontrado, a resposta deve conter JSON com campo `error`.
- Filtros por `status` e `priority` nao fazem parte do aceite do MVP atual.

### RF-004 - Atualizacao Parcial de Tarefa

- Dado um ID existente e payload valido, quando `PUT /tasks/{id}` for chamado, entao deve retornar HTTP `200`.
- A API deve atualizar apenas os campos informados no payload.
- Campos aceitos para atualizacao: `title`, `description`, `status`, `priority`.
- Quando `title` ou `description` forem alterados e `priority` nao for informada, a API deve recalcular ou preservar uma prioridade sugerida conforme regra vigente.
- Dado um ID inexistente, deve retornar HTTP `404`.
- Dado JSON malformado, deve retornar HTTP `400`.
- Valores invalidos de `status` ou `priority` devem retornar HTTP `400`.
- Quando `title` for atualizado para valor vazio, a API deve retornar HTTP `400`.

### RF-005 - Remocao de Tarefa

- Dado um ID existente, quando `DELETE /tasks/{id}` for chamado, entao deve retornar HTTP `204`.
- A resposta de sucesso nao deve exigir corpo.
- A tarefa removida nao deve aparecer em consultas posteriores.
- Dado um ID inexistente, deve retornar HTTP `404`.

### RF-006 - Priorizacao Assistida por IA

- Quando uma tarefa for criada sem `priority`, o mecanismo deve receber `title` e `description` como contexto.
- Quando uma tarefa tiver texto alterado sem `priority` explicita, o mecanismo deve poder recalcular a prioridade.
- A prioridade final deve pertencer ao conjunto `low`, `medium`, `high`, `critic`.
- Quando `OPENAI_API_KEY` nao estiver configurada, deve usar heuristica local.
- Quando o provedor externo falhar por timeout, HTTP nao-2xx, JSON invalido, resposta vazia ou prioridade invalida, deve usar heuristica local.
- A indisponibilidade do provedor externo nao deve derrubar o CRUD.
- Endpoint dedicado de sugestao nao faz parte do aceite do MVP atual.

### RF-007 - Contrato de Resposta JSON

- Respostas de sucesso devem usar JSON, exceto `DELETE` com HTTP `204`.
- Respostas de erro devem usar JSON com mensagem tecnica curta.
- O formato atual minimo de erro e `{"error":"mensagem"}`.
- Schema unico com `code`, `message` e `details` nao e criterio obrigatorio do MVP atual; permanece como evolucao planejada.

## Requisitos Tecnicos e Nao Funcionais

### RNF-001 - Stack e Compatibilidade

- O projeto deve ser implementado em Go conforme versao definida no `go.mod` e documentada no README.
- A API HTTP deve usar Fiber.
- O projeto deve ter instrucoes de execucao local para PowerShell.
- A execucao deve ser compativel com ambientes Linux em CI/CD quando pipeline for introduzido.

### RNF-002 - Desempenho Inicial

- Endpoints CRUD devem ter meta inicial de p95 <= 300 ms em carga baixa.
- `/health` deve responder sem chamada a banco, LLM ou outro servico externo.
- Medicao objetiva de desempenho deve ser adicionada antes de uso em ambiente mais amplo.

### RNF-003 - Confiabilidade

- Requisicoes invalidas nao devem encerrar o processo da API.
- Falhas de LLM devem ser tratadas com fallback.
- IDs inexistentes devem retornar erro controlado.
- Erros internos devem retornar resposta JSON tecnica curta.

### RNF-004 - Seguranca Basica

- Segredos como `OPENAI_API_KEY` nao devem ser versionados.
- Configuracoes sensiveis devem vir de variaveis de ambiente ou `.env` local.
- `.env.example` deve conter apenas placeholders seguros.
- Sem autenticacao/autorizacao, a API deve permanecer limitada a ambiente controlado.

### RNF-005 - Observabilidade Minima

- Falhas de integracao com IA devem ser registradas em log.
- Logs por requisicao com metodo, rota, status e duracao devem ser implementados como evolucao de qualidade.
- Metricas e tracing nao sao obrigatorios no MVP atual, mas devem ser considerados antes de producao.

### RNF-006 - Qualidade e Manutenibilidade

- O codigo deve permanecer organizado por camadas: API, service, repository e models.
- Regras de negocio nao devem depender diretamente da camada HTTP.
- `go test ./...` deve executar com sucesso.
- `go fmt ./...` nao deve deixar alteracoes pendentes antes de release.
- Lint e pipeline sao criterios planejados, ainda pendentes.

### RNF-007 - Testabilidade

- A suite deve ser executavel por comando unico: `go test ./...`.
- Devem existir testes automatizados para `GET /health`.
- Devem existir testes automatizados para criacao, listagem, consulta, atualizacao e remocao de tarefas.
- Devem existir testes para heuristica, sucesso com LLM simulado e fallback da priorizacao.
- Mudancas futuras de contrato devem vir acompanhadas de testes de regressao.

## Criterios de Aceite da Entrega do MVP

- Endpoints de health e CRUD estao funcionais.
- Priorizacao assistida esta operacional com fallback controlado.
- README contem setup, execucao, testes, variaveis de ambiente e limitacoes.
- Testes automatizados principais passam com `go test ./...`.
- Limitacoes do MVP estao documentadas: persistencia em memoria, ausencia de autenticacao, ausencia de filtros e ausencia de endpoint dedicado para sugestao.

## Lacunas de Implementacao ou Verificacao

- Adicionar teste explicito para `title` ausente ou vazio.
- Implementar e testar rejeicao HTTP `400` para valores invalidos de `status` e `priority`.
- Avaliar em mudanca futura se `status: "canceled"` deve entrar no contrato; ele nao faz parte do MVP atual.
- Criar OpenAPI para formalizar schemas, exemplos e status HTTP.
- Definir evidencia objetiva para desempenho e observabilidade.
