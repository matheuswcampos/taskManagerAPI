# Cenários de Integração da API de Tarefas

## Estratégia

Os cenários exercitam o fluxo HTTP do Fiber até `TaskService` e `TaskRepository`, com um repositório novo para cada teste. A LLM é isolada com `httptest.Server`; nenhuma chamada é feita a serviços externos reais. Middleware global não é aplicável, pois a inicialização atual não registra nenhum.

## Cenários

| ID | Fluxo integrado | Pré-condições | Ações | Resultados esperados | Camadas verificadas |
| --- | --- | --- | --- | --- | --- |
| INT-01 | Healthcheck | Aplicação criada por `newApp` | Enviar `GET /health` | `200`, JSON com `status: "ok"` e `timestamp` RFC3339 | Bootstrap, Fiber |
| INT-02 | Contrato CRUD encadeado | Aplicação nova, sem tarefas | Criar uma tarefa sem `status` e `priority`; consultar por ID; atualizar `title`; listar; excluir; consultar o mesmo ID | Criação `201` com `application/json` e campos de `TaskOut`; `status` inicial `todo`; consulta, atualização e listagem retornam `200`; atualização recalcula a prioridade; exclusão retorna `204`; consulta final retorna `404` com `error: "task not found"` | API, serviço, advisor, repositório, modelos |
| INT-03 | Criação com prioridade heurística | `OPENAI_API_KEY` ausente | Enviar `POST /tasks` com título que contenha `urgente` e sem `priority` | `201`; `priority` é `high`; não há chamada HTTP externa | API, serviço, `PriorityAdvisor`, repositório |
| INT-04 | Criação com resposta LLM válida | Servidor `httptest` retorna uma escolha com `{"priority":"critic"}`; LLM habilitada para o teste | Criar tarefa sem prioridade | `201`; prioridade retornada é `critic`; a requisição ao dublê usa `POST` e JSON | Serviço, `PriorityAdvisor`, cliente HTTP, repositório |
| INT-05 | Atualização parcial com prioridade explícita | Tarefa existente com prioridade `low`; advisor dublê observável | Enviar `PUT /tasks/{id}` com novo `title` e `priority: "critic"` | `200`; título e prioridade são atualizados; não há recálculo pelo advisor | API, serviço, repositório |
| INT-06 | Atualização parcial que recalcula prioridade | Tarefa existente sem prioridade explícita no payload de atualização | Enviar `PUT /tasks/{id}` alterando `description` com indício de urgência | `200`; campos não enviados são preservados e a prioridade é recalculada | API, serviço, advisor, repositório |
| INT-07 | Fallback heurístico diante de falha da LLM | Servidor `httptest` configurado como LLM retorna erro HTTP, JSON inválido ou prioridade inválida; LLM habilitada para o teste | Criar tarefa ou atualizar texto sem enviar `priority` | A operação retorna `201` ou `200` normalmente; a prioridade é calculada pela heurística local (por exemplo, `high` para texto urgente), sem expor erro do provedor | API, serviço, `PriorityAdvisor`, cliente HTTP, repositório |
| INT-08 | Erro de contrato por JSON malformado | Aplicação nova | Enviar `POST /tasks` e `PUT /tasks/{id}` com JSON incompleto | Ambos retornam `400`, JSON com campo `error` | API/handler |
| INT-09 | Recurso inexistente | Aplicação nova | Enviar `GET`, `PUT` e `DELETE` para ID não existente | Todas as operações retornam `404` com `error: "task not found"` | API, serviço, repositório |

## Prioridade de automação

| Cenário | Prioridade | Justificativa |
| --- | --- | --- |
| INT-01 | Média | Confirma que o bootstrap e a rota de disponibilidade permanecem operacionais. |
| INT-02 | Alta | Valida o principal fluxo de negócio e o contrato REST entre todas as camadas locais. |
| INT-03 | Alta | Garante o comportamento padrão sem dependência externa. |
| INT-04 | Alta | Protege a integração HTTP opcional com LLM sem usar credenciais reais. |
| INT-05 | Alta | Impede que uma prioridade fornecida pelo cliente seja sobrescrita. |
| INT-06 | Alta | Confirma a regra de recálculo acionada por mudança textual. |
| INT-07 | Alta | Garante resiliência do CRUD quando o provedor de prioridade falha ou responde fora do contrato. |
| INT-08 | Alta | Mantém o tratamento implementado de erro de entrada HTTP. |
| INT-09 | Alta | Mantém a conversão consistente de ausência no repositório para `404`. |

## Mapeamento sugerido

| Cenário | Arquivo de teste mais adequado |
| --- | --- |
| INT-01, INT-02 | `app/main_test.go` |
| INT-03, INT-07, INT-08, INT-09 | `tests/task_routes_test.go` |
| INT-05, INT-06 | `tests/task_service_test.go` |
| INT-04, INT-07 | `tests/priority_advisor_test.go` |
