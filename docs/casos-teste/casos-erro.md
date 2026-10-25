# Casos de Erro e Borda da API de Tarefas

## Classificação de falhas

- **Entrada HTTP:** corpo JSON malformado em criação ou atualização.
- **Recurso inexistente:** ID ausente do repositório em leitura, atualização ou remoção.
- **Casos de borda:** coleção vazia, payload vazio e prioridade explícita.
- **Integração LLM:** indisponibilidade, conexão encerrada, timeout, resposta não-`200` ou conteúdo inválido do provedor.

## Casos de teste

| ID | Categoria | Preparação | Requisição ou estímulo | Resposta/comportamento esperado | Asserções essenciais |
| --- | --- | --- | --- | --- | --- |
| ERR-01 | Entrada HTTP | Aplicação nova | `POST /tasks` com `{"title":` | `400` com JSON contendo `error` | Não cria tarefa no repositório |
| ERR-02 | Entrada HTTP | Tarefa existente | `PUT /tasks/{id}` com `{"status":` | `400` com JSON contendo `error` | A tarefa existente não é alterada |
| ERR-03 | Recurso inexistente | Aplicação nova | `GET /tasks/id-inexistente` | `404` com `{"error":"task not found"}` | Não retorna corpo de `TaskOut` |
| ERR-04 | Recurso inexistente | Aplicação nova | `PUT` e `DELETE /tasks/id-inexistente` | `404` com `{"error":"task not found"}` | Não há criação ou alteração de registros |
| ERR-05 | Casos de borda | Aplicação nova | `GET /tasks` | `200` e lista JSON vazia | Corpo decodifica como `[]` |
| ERR-06 | Casos de borda | Aplicação nova, advisor heurístico | `POST /tasks` com `{}` | `201`; tarefa com `title` vazio, `status: "todo"` e prioridade heurística `medium` | Registra o comportamento atual; não espera `400` |
| ERR-07 | Casos de borda | Tarefa existente | `PUT /tasks/{id}` com `{}` | `200`; valores de domínio são preservados | Não exige alteração de regra de negócio; `updated_at` é renovado pelo repositório |
| ERR-08 | Casos de borda | Tarefa existente e advisor observável | Atualizar texto com `priority: "critic"` | `200`; prioridade permanece `critic` | Advisor não é chamado |
| ERR-09 | Integração LLM | LLM habilitada apontando para endereço local sem listener ou servidor que encerra a conexão | Criar tarefa sem prioridade | `201`; prioridade é a heurística correspondente | Erro de conexão não é propagado ao handler nem interrompe CRUD |
| ERR-10 | Integração LLM | Servidor `httptest` atrasa além de `PRIORITY_ADVISOR_TIMEOUT` | Criar ou atualizar tarefa sem prioridade explícita | `201` ou `200`; prioridade é a heurística correspondente | A operação conclui após o timeout controlado, sem rede externa |
| ERR-11 | Integração LLM | Servidor `httptest` retorna `500` | Criar tarefa sem prioridade | `201`; prioridade vem da heurística | Falha HTTP não-`200` é absorvida pelo advisor |
| ERR-12 | Integração LLM | Servidor `httptest` retorna JSON inválido, `choices` vazia ou prioridade fora do contrato | Criar tarefa sem prioridade | `201`; prioridade vem da heurística | Cada variação usa um subteste isolado |
| ERR-13 | Integração LLM | Servidor `httptest` retorna `{"priority":"critical"}` no conteúdo da primeira escolha | Criar tarefa sem prioridade | `201`; prioridade é `critic` | Confirma normalização do valor aceito pelo advisor |

## Casos que não devem ser propostos

- Rejeição de título vazio e de valores inválidos de `status` ou `priority`, pois não há validação efetiva desses valores no comportamento atual.
- Falhas de banco de dados, migrações ou persistência durável, pois o repositório é exclusivamente em memória.
- Autenticação, autorização, CORS, rate limiting ou outro middleware global, pois nenhum middleware global está registrado.
- Fluxos de frontend, mensageria, OpenAPI ou integrações com Jira, Trello e Asana.

## Mapeamento sugerido

| Casos | Arquivo de teste mais adequado |
| --- | --- |
| ERR-01 a ERR-08 | `tests/task_routes_test.go` e `tests/task_service_test.go` |
| ERR-09 a ERR-13 | `tests/priority_advisor_test.go`, com cobertura de rota adicional se o fluxo HTTP completo for necessário |
