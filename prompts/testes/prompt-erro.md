# Prompt — Casos de Erro

## Papel

Você é especialista em testes voltados para casos de erro de APIs REST em Go. Seu foco é validar que falhas previstas no contrato atual sejam convertidas em respostas HTTP consistentes e que a integração opcional com LLM não interrompa o CRUD de tarefas.

## Contexto

O projeto é uma micro-API de tarefas em Go com Fiber, organizada nas camadas `api`, `services`, `repositories` e `models`.

- A camada HTTP expõe operações CRUD em `/tasks` e healthcheck em `/health`.
- JSON malformado em criação ou atualização deve produzir `400` com corpo JSON contendo o campo `error`.
- Consultar, atualizar ou excluir uma tarefa inexistente deve produzir `404` com `{"error":"task not found"}`.
- A integração opcional com LLM é executada internamente pelo `PriorityAdvisor` quando não há prioridade explícita.
- Se a LLM não estiver configurada, falhar na conexão, expirar, retornar HTTP não-2xx, JSON inválido, lista de escolhas vazia ou prioridade inválida, o `PriorityAdvisor` deve usar a heurística local e o fluxo de criação ou atualização deve continuar.
- O repositório é em memória; erros de infraestrutura persistente não pertencem ao comportamento atual.

## Objetivo

Defina casos de teste para falhas que já possuem tratamento no projeto. Avalie tanto a tradução de erros para o contrato HTTP quanto a resiliência do fluxo interno de priorização, usando dublês HTTP locais para simular falhas da LLM.

Ignore validações, integrações e componentes que não estejam implementados. Não inclua cenários de frontend, banco de dados, autenticação, autorização, mensageria, OpenAPI ou ferramentas externas de gestão.

## Critérios

1. Diferencie erros de entrada, recurso inexistente, casos de borda do contrato atual e degradação da integração com LLM.
2. Para cada caso HTTP, informe método, rota, payload quando houver, status esperado e corpo esperado.
3. Inclua casos de borda implementados: `GET /tasks` sem registros, criação com payload `{}`, atualização com payload `{}` e prioridade explícita impedindo o recálculo pela LLM ou heurística.
4. Para cada falha de LLM, confirme que a prioridade vem da heurística e que a operação CRUD retorna `201` ou `200` conforme o fluxo.
5. Simule conexão recusada ou encerrada por servidor HTTP local e timeout por resposta deliberadamente atrasada; não use rede externa.
6. Inclua resposta LLM com `"priority":"critical"` e confirme a normalização para `critic`.
7. Verifique que nenhuma credencial real, rede externa ou estado compartilhado seja necessária.
8. Não atribua ao sistema respostas ou validações que o código atual não realiza, como rejeição de título vazio ou valores inválidos de domínio.
9. Considere middleware global como não aplicável, pois a aplicação atual não registra middleware global.
10. Priorize cenários reproduzíveis com `httptest`, `app.Test` e um repositório novo por teste.

## Formato de saída

Responda em Markdown com:

1. Os casos devem ser gerados no diretório `docs/casos-teste`.
2. Uma seção `## Classificação de falhas` separando `Entrada HTTP`, `Recurso inexistente`, `Casos de borda` e `Integração LLM`.
3. Uma tabela `## Casos de teste` com as colunas: `ID`, `Categoria`, `Preparação`, `Requisição ou estímulo`, `Resposta/comportamento esperado` e `Asserções essenciais`.
4. Uma seção `## Casos que não devem ser propostos` listando explicitamente os limites adotados, com base no escopo implementado.
5. Uma seção `## Mapeamento sugerido` relacionando cada caso aos arquivos de teste existentes mais adequados.
