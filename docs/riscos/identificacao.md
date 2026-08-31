# Identificacao de Riscos

Lista de riscos identificados:

## 1) Evolucao dos requisitos gerar retrabalho e instabilidade de escopo

**Risco:** Evolucao dos requisitos gerar retrabalho e instabilidade de escopo.

**Descricao:** O contexto informa que os requisitos estao em evolucao. Mesmo com o contrato atual do MVP documentado, mudancas futuras podem alterar criterios de aceite, comportamento da API, priorizacao assistida por IA, persistencia ou requisitos nao funcionais.

**Contexto de ocorrencia:** Pode ocorrer quando novos entendimentos sobre o MVP ou sua evolucao forem incorporados durante o desenvolvimento. A frequencia, o responsavel por aprovacao e o processo de avaliacao de impacto precisam ser validados.

## 2) Mudancas futuras no contrato da API afetarem consumidores e testes

**Risco:** Mudancas futuras no contrato da API afetarem consumidores e testes.

**Descricao:** O contrato atual documenta `PUT /tasks/{id}`, listagem sem filtros, erro JSON simples com campo `error` e priorizacao assistida sem endpoint dedicado. Alguns itens permanecem como evolucao planejada, como schema unico de erro, filtros, OpenAPI e possiveis ajustes de contrato. Essas mudancas podem exigir adaptacao de clientes, testes e documentacao.

**Contexto de ocorrencia:** Pode se manifestar quando itens de qualidade ou evolucao forem priorizados, especialmente padronizacao de erros, filtros de listagem, documentacao OpenAPI ou endpoint dedicado de sugestao de prioridade.

## 3) Dependencia operacional da integracao opcional com LLM

**Risco:** Dependencia operacional da integracao opcional com LLM.

**Descricao:** A priorizacao pode chamar uma API externa de LLM quando `OPENAI_API_KEY` estiver configurada. Essa dependencia pode introduzir latencia, custo, indisponibilidade temporaria e respostas fora do contrato esperado.

**Contexto de ocorrencia:** Pode ocorrer em ambientes com LLM habilitado, principalmente diante de timeout, HTTP nao-2xx, JSON invalido, `choices` vazio, prioridade invalida ou configuracao incorreta de `OPENAI_BASE_URL`, `OPENAI_MODEL` e `PRIORITY_ADVISOR_TIMEOUT`.

## 4) Fallback por heuristica local produzir priorizacoes inadequadas

**Risco:** Fallback por heuristica local produzir priorizacoes inadequadas.

**Descricao:** Quando a chave da OpenAI nao existe ou a chamada externa falha, o projeto usa heuristica local baseada em palavras-chave. Esse comportamento preserva o CRUD, mas pode gerar sugestoes diferentes das esperadas pela equipe, especialmente em tarefas com contexto ambiguo.

**Contexto de ocorrencia:** Pode se manifestar em criacao ou atualizacao de tarefas sem prioridade explicita, quando o LLM estiver desabilitado, indisponivel ou retornar resposta invalida. A aceitabilidade da heuristica precisa ser validada com criterios reais de priorizacao.

## 5) Perda de dados por persistencia em memoria

**Risco:** Perda de dados por persistencia em memoria.

**Descricao:** O MVP usa persistencia em memoria com `map` e `mutex`. Os dados cadastrados nao sobrevivem a restart, falha do processo ou nova execucao da aplicacao.

**Contexto de ocorrencia:** Pode ocorrer quando a API for reiniciada, atualizada, interrompida ou usada em ambiente de homologacao/uso interno com expectativa de continuidade dos dados. A criticidade depende de o uso ser demonstrativo ou exigir durabilidade.

## 6) Evolucao para banco de dados exigir mudancas tecnicas relevantes

**Risco:** Evolucao para banco de dados exigir mudancas tecnicas relevantes.

**Descricao:** A migracao para PostgreSQL esta planejada para release futura. Essa evolucao pode exigir migracoes, configuracao de ambiente, tratamento de erros de banco, testes de integracao e revisao de desempenho.

**Contexto de ocorrencia:** Pode ocorrer quando persistencia duravel for priorizada ou quando stakeholders exigirem uso continuado da API. Ainda ha incerteza sobre momento, ambiente e requisitos operacionais dessa evolucao.

## 7) Exposicao da API sem autenticacao ou autorizacao

**Risco:** Exposicao da API sem autenticacao ou autorizacao.

**Descricao:** Autenticacao e autorizacao estao fora do escopo do MVP. Isso e aceitavel para ambiente controlado, mas se torna risco caso a API seja exposta a usuarios, redes ou ambientes nao controlados.

**Contexto de ocorrencia:** Pode ocorrer se a API for disponibilizada em homologacao aberta, rede compartilhada ou ambiente com usuarios diferentes acessando dados de tarefas. A exposicao real da API precisa ser validada.

## 8) Observabilidade insuficiente para diagnostico operacional

**Risco:** Observabilidade insuficiente para diagnostico operacional.

**Descricao:** O projeto ainda nao possui observabilidade completa, como metricas e tracing. Ha logs relacionados ao `PriorityAdvisor`, mas a ausencia de sinais operacionais mais completos pode dificultar diagnostico de latencia, erros e falhas de integracao.

**Contexto de ocorrencia:** Pode se manifestar durante investigacao de falhas, avaliacao de desempenho, acompanhamento do uso da IA ou homologacao em ambiente compartilhado. A necessidade exata depende dos requisitos operacionais.

## 9) Validacao de entrada nao cobrir todos os valores de dominio

**Risco:** Validacao de entrada nao cobrir todos os valores de dominio.

**Descricao:** O escopo exige validacao de payload para reduzir dados invalidos. Os modelos declaram valores esperados para `status` e `priority`, mas a leitura do fluxo HTTP indica que a validacao efetiva deve ser confirmada para casos como titulo vazio, status invalido e prioridade invalida.

**Contexto de ocorrencia:** Pode ocorrer quando clientes enviarem payloads incompletos ou valores fora do dominio aceito. A extensao do risco precisa ser validada com testes automatizados e revisao do fluxo de validacao.

## 10) Coordenacao em equipe distribuida afetar decisoes e sincronizacao

**Risco:** Coordenacao em equipe distribuida afetar decisoes e sincronizacao.

**Descricao:** A equipe e distribuida e os requisitos estao em evolucao. Isso pode dificultar alinhamento sobre prioridades, contratos de API, decisoes de arquitetura e momento de incorporar evolucoes planejadas.

**Contexto de ocorrencia:** Pode se manifestar quando membros diferentes atuarem em backlog, codigo, documentacao, testes e arquitetura sem rituais ou canais claros de alinhamento. Nao foram fornecidas informacoes sobre esse processo, portanto o ponto requer validacao.

## 11) Testes nao acompanharem mudancas de contrato e arquitetura

**Risco:** Testes nao acompanharem mudancas de contrato e arquitetura.

**Descricao:** A suite atual cobre rotas, servico, repositorio e `PriorityAdvisor`, mas evolucoes como schema padronizado de erro, filtros, OpenAPI, banco de dados e validacoes adicionais podem exigir novos testes unitarios, de contrato e de integracao.

**Contexto de ocorrencia:** Pode ocorrer quando novos requisitos forem incorporados ou quando a persistencia e a integracao com LLM forem alteradas. A cobertura deve ser reavaliada a cada mudanca relevante de contrato.
