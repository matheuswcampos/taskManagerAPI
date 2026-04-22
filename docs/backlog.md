# Backlog MVP - Micro-API de Tarefas

## Release 1 - Core

- [ ] **RF-001 - Health Check**  
  **Descrição:** disponibilizar `GET /health` com status de serviço.  
  **Critérios de aceite:**  
  - retorna HTTP `200`  
  - corpo JSON contém `status: "ok"`  
  - corpo JSON contém `timestamp` em UTC (formato RFC3339)

- [ ] **RF-002 - Criar tarefa**  
  **Descrição:** implementar endpoint para criação de tarefa.  
  **Critérios de aceite:**  
  - `POST /tasks` cria tarefa com `id`, `title`, `description`, `priority`, `status`, `created_at`  
  - `title` obrigatório; sem `title` retorna HTTP `400`  
  - `status` padrão definido como `todo` quando não informado

- [ ] **RF-003 - Listar tarefas**  
  **Descrição:** implementar listagem de tarefas.  
  **Critérios de aceite:**  
  - `GET /tasks` retorna HTTP `200`  
  - resposta é array JSON de tarefas  
  - filtros por `status` e `priority` funcionam quando informados

- [ ] **RF-004 - Atualizar tarefa**  
  **Descrição:** permitir atualização parcial de tarefa por ID.  
  **Critérios de aceite:**  
  - `PATCH /tasks/{id}` atualiza campos válidos  
  - tarefa inexistente retorna HTTP `404`  
  - payload inválido retorna HTTP `400`

- [ ] **RF-005 - Remover tarefa**  
  **Descrição:** excluir tarefa por ID.  
  **Critérios de aceite:**  
  - `DELETE /tasks/{id}` retorna HTTP `204` em sucesso  
  - tarefa inexistente retorna HTTP `404`

- [ ] **RT-001 - Estrutura técnica inicial**  
  **Descrição:** organizar projeto por camadas (`handler`, `service`, `repository`, `model`).  
  **Critérios de aceite:**  
  - código do domínio não depende de camada HTTP  
  - separação de responsabilidades evidenciada por pacotes

## Release 2 - Qualidade

- [ ] **RT-002 - Testes automatizados unitários**  
  **Descrição:** criar suíte de testes de regra de negócio.  
  **Critérios de aceite:**  
  - comando `go test ./...` executa com sucesso  
  - fluxos críticos (criação e atualização de tarefa) cobertos por testes

- [ ] **RT-003 - Testes de endpoint**  
  **Descrição:** validar contrato HTTP com testes de handler.  
  **Critérios de aceite:**  
  - testes de `GET /health`, `POST /tasks` e `GET /tasks` implementados  
  - validação de status code e payload JSON

- [ ] **RT-004 - Padronização de erros e logs**  
  **Descrição:** implementar padrão de resposta de erro e logs estruturados.  
  **Critérios de aceite:**  
  - erros seguem formato JSON único (`code`, `message`, `details`)  
  - logs incluem método, rota, status e duração

- [ ] **RT-005 - Qualidade estática**  
  **Descrição:** adicionar formatação/lint no fluxo local e CI.  
  **Critérios de aceite:**  
  - `go fmt ./...` sem alterações pendentes após execução  
  - lint configurado e sem erros bloqueantes no pipeline

- [ ] **RF-006 - Priorização assistida por IA**  
  **Descrição:** disponibilizar endpoint para sugestão de prioridade de tarefa.  
  **Critérios de aceite:**  
  - endpoint recebe contexto textual (`title`, `description`)  
  - retorna `priority` sugerida (`low|medium|high`) e justificativa curta  
  - indisponibilidade do provedor IA retorna erro controlado sem derrubar a API

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

- [ ] **RF-007 - Pronto para uso interno**  
  **Descrição:** disponibilizar versão estável para uso da equipe interna.  
  **Critérios de aceite:**  
  - rotas core e priorização por IA operacionais  
  - sem erros críticos abertos  
  - release taggeada e registrada no repositório
