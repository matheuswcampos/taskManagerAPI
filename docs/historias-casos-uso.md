# Historias de Usuario e Casos de Uso

Este documento complementa o escopo do MVP com uma visao orientada ao uso. As historias descrevem necessidades dos usuarios e os casos de uso detalham os fluxos esperados da API.

## Atores

| Ator | Descricao | Interesse principal |
|---|---|---|
| Integrador da API | Pessoa ou sistema que consome os endpoints HTTP | Registrar, consultar, atualizar e remover tarefas por JSON. |
| Membro da equipe interna | Usuario representado pelo cliente que consome a API | Acompanhar tarefas e priorizar execucao. |
| Responsavel tecnico | Pessoa que mantem a API | Garantir funcionamento, testes, configuracao e evolucao segura. |
| Stakeholder de negocio | Pessoa que acompanha escopo e andamento do MVP | Validar valor, riscos e prontidao para uso interno. |

## Historias de Usuario

| ID | Historia | Requisitos relacionados | Prioridade | Status |
|---|---|---|---|---|
| HU-001 | Como integrador da API, quero verificar se o servico esta disponivel para decidir se posso enviar requisicoes. | RF-001, RF-007, RNF-002 | Alta | Implementada |
| HU-002 | Como integrador da API, quero criar uma tarefa com titulo e descricao para registrar demandas da equipe. | RF-002, RF-006, RF-007 | Alta | Parcial |
| HU-003 | Como membro da equipe interna, quero que uma nova tarefa receba status inicial automaticamente para manter o fluxo padronizado. | RF-002 | Alta | Implementada |
| HU-004 | Como membro da equipe interna, quero receber uma prioridade sugerida quando eu nao informar prioridade para apoiar a organizacao do trabalho. | RF-006 | Alta | Implementada |
| HU-005 | Como integrador da API, quero listar tarefas cadastradas para apresentar o backlog atual da equipe. | RF-003, RF-007 | Alta | Implementada |
| HU-006 | Como integrador da API, quero consultar uma tarefa por identificador para exibir ou processar seus detalhes. | RF-003, RF-007 | Media | Implementada |
| HU-007 | Como integrador da API, quero atualizar parcialmente uma tarefa para alterar apenas os campos necessarios. | RF-004, RF-006, RF-007 | Alta | Parcial |
| HU-008 | Como integrador da API, quero remover uma tarefa por identificador para manter o cadastro limpo. | RF-005, RF-007 | Media | Implementada |
| HU-009 | Como responsavel tecnico, quero que falhas do provedor de IA nao indisponibilizem o CRUD para manter a API operacional. | RF-006, RNF-003 | Alta | Implementada |
| HU-010 | Como responsavel tecnico, quero executar testes automatizados por comando unico para reduzir regressao em mudancas futuras. | RNF-007, RT-002, RT-003 | Alta | Implementada |
| HU-011 | Como stakeholder de negocio, quero conhecer limitacoes e riscos do MVP para decidir seu uso em ambiente interno. | RNF-004, RNF-005, docs de riscos | Media | Implementada em documentacao |

## Casos de Uso

### UC-001 - Verificar disponibilidade da API

**Ator principal:** Integrador da API.

**Objetivo:** Confirmar que a API esta ativa.

**Pre-condicoes:** A aplicacao esta em execucao.

**Fluxo principal:**

1. O ator envia `GET /health`.
2. A API retorna HTTP `200`.
3. A API retorna JSON com `status: "ok"` e `timestamp`.

**Fluxos alternativos:** Nao aplicavel no contrato atual.

**Pos-condicoes:** O ator pode usar a resposta para diagnostico simples de disponibilidade.

### UC-002 - Criar tarefa

**Ator principal:** Integrador da API.

**Objetivo:** Registrar uma nova tarefa.

**Pre-condicoes:** A aplicacao esta em execucao.

**Fluxo principal:**

1. O ator envia `POST /tasks` com JSON contendo `title` e, opcionalmente, `description`, `status` e `priority`.
2. A API interpreta o payload.
3. A API aplica `status: "todo"` quando `status` nao for informado.
4. A API sugere `priority` quando `priority` nao for informada.
5. A API persiste a tarefa no repositorio atual.
6. A API retorna HTTP `201` com a tarefa criada.

**Fluxos alternativos:**

- Payload JSON invalido: retorna HTTP `400` com campo `error`.
- Titulo ausente ou vazio: retorna HTTP `400`.
- `status` ou `priority` invalido: retorna HTTP `400`.
- Provedor de IA indisponivel: usa heuristica local e mantem o fluxo de criacao.

**Pos-condicoes:** A tarefa fica disponivel para consulta e listagem enquanto o processo da API estiver ativo.

### UC-003 - Listar tarefas

**Ator principal:** Integrador da API.

**Objetivo:** Obter todas as tarefas cadastradas.

**Pre-condicoes:** A aplicacao esta em execucao.

**Fluxo principal:**

1. O ator envia `GET /tasks`.
2. A API consulta o repositorio atual.
3. A API retorna HTTP `200` com array JSON de tarefas.

**Fluxos alternativos:** Nao ha filtros no contrato atual do MVP.

**Pos-condicoes:** Nenhum dado e alterado.

### UC-004 - Consultar tarefa por ID

**Ator principal:** Integrador da API.

**Objetivo:** Obter os dados de uma tarefa especifica.

**Pre-condicoes:** A aplicacao esta em execucao.

**Fluxo principal:**

1. O ator envia `GET /tasks/{id}`.
2. A API busca a tarefa pelo identificador.
3. A API retorna HTTP `200` com a tarefa encontrada.

**Fluxos alternativos:**

- ID inexistente: retorna HTTP `404` com campo `error`.

**Pos-condicoes:** Nenhum dado e alterado.

### UC-005 - Atualizar tarefa parcialmente

**Ator principal:** Integrador da API.

**Objetivo:** Alterar campos especificos de uma tarefa.

**Pre-condicoes:** A tarefa existe no repositorio atual.

**Fluxo principal:**

1. O ator envia `PUT /tasks/{id}` com um ou mais campos entre `title`, `description`, `status` e `priority`.
2. A API interpreta o payload.
3. A API altera somente os campos informados.
4. Quando texto muda e `priority` nao e informada, a API pode recalcular a prioridade sugerida.
5. A API retorna HTTP `200` com a tarefa atualizada.

**Fluxos alternativos:**

- Payload JSON invalido: retorna HTTP `400` com campo `error`.
- ID inexistente: retorna HTTP `404` com campo `error`.
- Titulo vazio: retorna HTTP `400`.
- `status` ou `priority` invalido: retorna HTTP `400`.
- Provedor de IA indisponivel no recalcule: usa fallback sem derrubar a API.

**Pos-condicoes:** A tarefa permanece cadastrada com os campos atualizados.

### UC-006 - Remover tarefa

**Ator principal:** Integrador da API.

**Objetivo:** Excluir uma tarefa cadastrada.

**Pre-condicoes:** A aplicacao esta em execucao.

**Fluxo principal:**

1. O ator envia `DELETE /tasks/{id}`.
2. A API remove a tarefa quando o identificador existe.
3. A API retorna HTTP `204` sem corpo.

**Fluxos alternativos:**

- ID inexistente: retorna HTTP `404` com campo `error`.

**Pos-condicoes:** A tarefa removida nao aparece em consultas posteriores.

### UC-007 - Sugerir prioridade internamente

**Ator principal:** Sistema.

**Objetivo:** Definir uma prioridade para tarefas sem prioridade explicita.

**Pre-condicoes:** Criacao ou atualizacao de tarefa sem `priority` informada.

**Fluxo principal:**

1. O servico recebe `title` e `description`.
2. O `PriorityAdvisor` verifica se ha configuracao para LLM.
3. Se houver LLM habilitado, o componente solicita uma resposta JSON com `priority`.
4. A API aceita somente `low`, `medium`, `high` ou `critic`.
5. A prioridade sugerida e aplicada a tarefa.

**Fluxos alternativos:**

- Sem chave de API: usa heuristica local.
- Timeout, HTTP nao-2xx, JSON invalido, resposta vazia ou prioridade invalida: usa heuristica local.

**Pos-condicoes:** O CRUD continua funcional mesmo quando a dependencia externa falha.

## Fora do Fluxo Atual

- Endpoint dedicado para sugestao de prioridade.
- Filtros em `GET /tasks`.
- Status `canceled`.
- Autenticacao e autorizacao.
- Persistencia duravel em banco.
- Interface web para operacao.
