# Escopo do MVP - Micro-API de Tarefas

## Objetivo
Entregar um MVP de micro-API para gestão de tarefas de uma equipe interna, com foco em registrar tarefas, acompanhar status e priorizar execução com apoio de IA.

O MVP deve ser suficientemente simples para validação rápida de uso, mas com base técnica adequada para evolução incremental em produção.

## Requisitos Funcionais

### RF-01 - Health Check
A API deve disponibilizar endpoint de verificação de disponibilidade:
- `GET /health`
- resposta com `status` e `timestamp` em UTC

### RF-02 - Cadastro de Tarefa
A API deve permitir criar tarefa com os campos mínimos:
- `title` (obrigatório)
- `description` (opcional)
- `priority` (opcional: `low`, `medium`, `high`)
- `status` (padrão inicial: `todo`)

### RF-03 - Listagem de Tarefas
A API deve permitir listar tarefas cadastradas, incluindo filtros básicos:
- por `status`
- por `priority`

### RF-04 - Atualização de Tarefa
A API deve permitir atualização parcial dos campos da tarefa:
- `title`
- `description`
- `priority`
- `status`

### RF-05 - Remoção de Tarefa
A API deve permitir exclusão de tarefa por identificador único.

### RF-06 - Sugestão de Prioridade Assistida por IA
A API deve oferecer mecanismo para sugerir prioridade de tarefa com base em título e descrição:
- entrada: contexto textual da tarefa
- saída: prioridade sugerida (`low`, `medium`, `high`) e justificativa curta
- em indisponibilidade do provedor de IA, retornar erro tratável sem indisponibilizar a API inteira

### RF-07 - Contrato de Resposta Padronizado
A API deve retornar respostas JSON consistentes para sucesso e erro, incluindo:
- código HTTP adequado (`200`, `201`, `204`, `400`, `404`, `422`, `500`)
- Payloads JSON tipados via schema
- mensagem técnica curta para troubleshooting

## Requisitos Não Funcionais

### RNF-01 - Stack e Compatibilidade
- implementação em Go 1.25+
- framework HTTP: Fiber
- execução local em ambiente Windows (PowerShell) e compatível com ambientes Linux em CI/CD

### RNF-02 - Desempenho Inicial
- endpoints CRUD devem responder com latência adequada para uso interno (meta inicial p95 <= 300 ms em carga baixa)
- endpoint `/health` deve ser leve e não depender de serviços externos

### RNF-03 - Confiabilidade
- tratamento explícito de erros de validação e erros internos
- API não deve encerrar processo por erro de requisição inválida

### RNF-04 - Segurança Básica para MVP
- não expor segredos em código-fonte
- configuração via variáveis de ambiente
- validação de payload de entrada para reduzir risco de dados inválidos

### RNF-05 - Observabilidade Mínima
- logs estruturados por requisição com nível de severidade
- registro de falhas de integração com IA

### RNF-06 - Qualidade e Manutenibilidade
- organização de código por camadas (handler, service, repository)
- lint/testes automatizados no pipeline mínimo
- cobertura inicial em testes unitários para fluxos críticos

### RNF-07 - Testabilidade

#### Objetivo de Testabilidade
Garantir que os componentes da API sejam verificáveis de forma isolada e integrada, com execução automatizada em ambiente local e pipeline.

#### Opções de Estratégia de Testes
- testes unitários para regras de negócio (services e validações), sem dependência externa
- testes de handler com `net/http/httptest` para validar contrato HTTP (status code, payload e erros)
- testes de integração com infraestrutura real (ex.: banco) para fluxos críticos de persistência
- testes de contrato para garantir estabilidade dos schemas JSON dos endpoints principais
- testes de carga básicos para validar comportamento inicial de latência e throughput

#### Ferramentas Recomendadas
- ferramenta principal de automação: `go test` (pacote padrão `testing`)
- asserções e utilitários: `testify`
- testes HTTP: `net/http/httptest`
- integração com dependências reais: `testcontainers-go` (quando houver banco/serviços externos)
- cobertura: `go test -cover ./...`

#### Critérios Mínimos para o MVP
- suíte automatizada executável via comando único (`go test ./...`)
- validação obrigatória do endpoint `GET /health`
- validação dos fluxos de criação e atualização de tarefa
- falha no pipeline em caso de regressão em testes automatizados

## Fora de Escopo (MVP)
- autenticação/autorização com gestão de usuários e perfis
- interface web (frontend) para operação de tarefas
- notificações por e-mail, Slack ou outros canais
- anexos de arquivos nas tarefas
- priorização multi-critério avançada com regras customizáveis por equipe
- auditoria completa e trilha de compliance
- multi-tenant
- alta disponibilidade com escalabilidade horizontal formal
- versionamento avançado de API além da baseline inicial
- integrações com ferramentas externas de gestão (Jira, Trello, Asana)

## Critério de Conclusão do MVP
O MVP será considerado concluído quando:
- os endpoints de health e CRUD estiverem funcionais e documentados
- a sugestão de prioridade por IA estiver operacional com fallback de erro controlado
- o projeto puder ser executado localmente com instruções reproduzíveis
- existir base mínima de testes automatizados para os fluxos principais
