# Backlog MVP - Micro-API de Tarefas

## Release 1 - Core

- [x] **RF-001 - Health Check**  
  **Descrição:** disponibilizar `GET /health` com status de serviço.  
  **Critérios de aceite:**  
  - retorna HTTP `200`  
  - corpo JSON contém `status: "ok"`  
  - corpo JSON contém `timestamp` em UTC (formato RFC3339)

- [x] **RF-002 - Criar tarefa**  
  **Descrição:** implementar endpoint para criação de tarefa.  
  **Critérios de aceite:**  
  - `POST /tasks` cria tarefa com `id`, `title`, `description`, `priority`, `status`, `created_at`  
  - `title` obrigatório; sem `title` retorna HTTP `400`  
  - `status` padrão definido como `todo` quando não informado

- [x] **RF-003 - Listar e consultar tarefas**  
  **Descrição:** implementar listagem de tarefas e consulta por ID.  
  **Critérios de aceite:**  
  - `GET /tasks` retorna HTTP `200`  
  - resposta é array JSON de tarefas  
  - lista todas as tarefas cadastradas no repositório atual
  - `GET /tasks/{id}` retorna HTTP `200` quando a tarefa existe
  - `GET /tasks/{id}` retorna HTTP `404` quando a tarefa não existe

  **Observação:** filtros por `status` e `priority` não fazem parte do contrato atual do MVP.

- [x] **RF-004 - Atualizar tarefa**  
  **Descrição:** permitir atualização parcial de tarefa por ID.  
  **Critérios de aceite:**  
  - `PUT /tasks/{id}` atualiza campos válidos  
  - tarefa inexistente retorna HTTP `404`  
  - payload inválido retorna HTTP `400`

- [x] **RF-005 - Remover tarefa**  
  **Descrição:** excluir tarefa por ID.  
  **Critérios de aceite:**  
  - `DELETE /tasks/{id}` retorna HTTP `204` em sucesso  
  - tarefa inexistente retorna HTTP `404`

- [x] **RT-001 - Estrutura técnica inicial**  
  **Descrição:** organizar projeto por camadas (`handler`, `service`, `repository`, `model`).  
  **Critérios de aceite:**  
  - código do domínio não depende de camada HTTP  
  - separação de responsabilidades evidenciada por pacotes

## Release 2 - Qualidade

- [x] **RT-002 - Testes automatizados unitários**  
  **Descrição:** criar suíte de testes de regra de negócio.  
  **Critérios de aceite:**  
  - comando `go test ./...` executa com sucesso  
  - fluxos críticos (criação e atualização de tarefa) cobertos por testes

- [x] **RT-003 - Testes de endpoint**  
  **Descrição:** validar contrato HTTP com testes de handler.  
  **Critérios de aceite:**  
  - testes de `GET /health`, `POST /tasks`, `GET /tasks` e `GET /tasks/{id}` implementados  
  - validação de status code e payload JSON

- [ ] **RT-004 - Padronização de erros e logs**  
  **Descrição:** implementar padrão de resposta de erro e logs estruturados.  
  **Critérios de aceite:**  
  - erros seguem formato JSON único (`code`, `message`, `details`)  
  - logs incluem método, rota, status e duração

  **Observação:** o MVP atual já retorna erros em JSON simples com campo `error`; o formato único acima é evolução de qualidade ainda pendente.

- [ ] **RT-005 - Qualidade estática**  
  **Descrição:** adicionar formatação/lint no fluxo local e CI.  
  **Critérios de aceite:**  
  - `go fmt ./...` sem alterações pendentes após execução  
  - lint configurado e sem erros bloqueantes no pipeline

- [x] **RF-006 - Priorização assistida por IA**  
  **Descrição:** aplicar sugestão de prioridade de tarefa no fluxo de criação/atualização quando a prioridade não for informada.  
  **Critérios de aceite:**  
  - mecanismo recebe contexto textual (`title`, `description`)  
  - define `priority` sugerida (`low|medium|high|critic`)  
  - indisponibilidade do provedor IA aciona fallback por heurística sem derrubar a API

  **Observação:** endpoint dedicado de sugestão e justificativa curta não fazem parte do contrato atual do MVP.

- [x] **RF-007 - Contrato de resposta JSON**  
  **Descrição:** manter respostas HTTP em JSON para sucesso e erro, conforme contrato atual do MVP.  
  **Critérios de aceite:**  
  - respostas de sucesso retornam JSON, exceto `DELETE /tasks/{id}` com HTTP `204`
  - respostas de erro retornam JSON com campo `error`
  - status HTTP principais documentados no README e cobertos por testes de rota

  **Observação:** schema único com `code`, `message` e `details` permanece como evolução em `RT-004`.

## Release 3 - Entrega Final

- [ ] **RT-006 - Persistência em banco**  
  **Descrição:** migrar armazenamento em memória para banco relacional (PostgreSQL).  
  **Critérios de aceite:**  
  - operações CRUD usam repositório persistente  
  - migração inicial de schema disponível e reproduzível

- [ ] **RT-007 - Documentação de API**  
  **Descrição:** publicar documentação técnica dos endpoints.  
  **Critérios de aceite:**  
  - especificação OpenAPI atualizada  
  - exemplos de request/response para rotas principais

- [ ] **RT-008 - Prontidão de entrega**  
  **Descrição:** garantir execução reproduzível para homologação interna.  
  **Critérios de aceite:**  
  - README com setup completo validado  
  - variáveis de ambiente documentadas  
  - build e testes passando no pipeline

- [ ] **RT-009 - Pronto para uso interno**  
  **Descrição:** disponibilizar versão estável para uso da equipe interna.  
  **Critérios de aceite:**  
  - rotas core e priorização por IA operacionais  
  - sem erros críticos abertos  
  - release taggeada e registrada no repositório
