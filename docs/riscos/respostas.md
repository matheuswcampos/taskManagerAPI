# Estrategias de Resposta a Riscos

Este documento apresenta respostas possiveis para os riscos identificados e analisados. As estrategias nao substituem decisao de projeto; elas indicam como reduzir exposicao, aceitar conscientemente limitacoes do MVP ou preparar evolucoes futuras.

## 1) Estrategias por risco

### 1. Evolucao dos requisitos gerar retrabalho e instabilidade de escopo

**Mitigar:** manter backlog priorizado, registrar decisoes, revisar impacto antes de aceitar mudancas e validar criterios de aceite em ciclos curtos.

**Aceitar:** aceitar mudancas pequenas e compativeis com o MVP quando o custo de controle formal for maior que o retrabalho esperado.

### 2. Mudancas futuras no contrato da API afetarem consumidores e testes

**Mitigar:** versionar decisoes de contrato, atualizar README/escopo/backlog juntos, adicionar testes de contrato antes de expor mudancas e documentar claramente o que e atual versus planejado.

**Evitar:** evitar integrar consumidores externos antes de estabilizar contratos como erro padronizado, filtros e OpenAPI.

### 3. Dependencia operacional da integracao opcional com LLM

**Mitigar:** manter timeout, fallback local, validacao de resposta, logs de falha e configuracao explicita por ambiente.

**Aceitar:** manter LLM como recurso complementar quando heuristica local for suficiente para o uso minimo.

### 4. Fallback por heuristica local produzir priorizacoes inadequadas

**Mitigar:** validar exemplos de priorizacao com usuarios, ajustar palavras-chave conforme casos reais e criar testes para cenarios representativos.

**Aceitar:** aceitar divergencias pontuais se a prioridade sugerida for apenas apoio e puder ser sobrescrita pelo usuario ou cliente da API.

### 5. Perda de dados por persistencia em memoria

**Evitar:** nao usar a API para dados que precisem de durabilidade enquanto a persistencia em banco nao estiver implementada.

**Mitigar:** documentar a limitacao, restringir uso a demonstracao/homologacao controlada e priorizar PostgreSQL quando houver necessidade de continuidade dos dados.

### 6. Evolucao para banco de dados exigir mudancas tecnicas relevantes

**Mitigar:** manter interface de repositorio isolada, planejar migracoes, adicionar testes de integracao e preparar configuracao reproduzivel de banco.

**Transferir/Compartilhar:** envolver responsaveis por infraestrutura e stakeholders na definicao de ambiente, backup, credenciais e operacao.

### 7. Exposicao da API sem autenticacao ou autorizacao

**Evitar:** nao expor a API fora de ambiente controlado sem autenticacao/autorizacao.

**Mitigar:** limitar rede/acesso no ambiente de execucao e explicitar que autenticacao/autorizacao estao fora do MVP atual.

### 8. Observabilidade insuficiente para diagnostico operacional

**Mitigar:** adicionar logs estruturados por requisicao, metricas basicas de latencia/erro e sinais especificos para chamadas LLM.

**Aceitar:** aceitar observabilidade minima apenas em execucao local ou demonstracao, quando investigacao operacional avancada nao for necessaria.

### 9. Validacao de entrada nao cobrir todos os valores de dominio

**Mitigar:** adicionar validacao explicita para titulo, status e prioridade; criar testes para payloads invalidos; alinhar codigos HTTP esperados.

**Evitar:** evitar expandir consumidores antes de confirmar que dados invalidos nao entram no dominio.

### 10. Coordenacao em equipe distribuida afetar decisoes e sincronizacao

**Mitigar:** definir fonte de verdade para escopo, registrar decisoes tecnicas, manter cadencia de alinhamento e sinalizar mudancas de contrato para todos os envolvidos.

**Transferir/Compartilhar:** atribuir responsabilidade por priorizacao e aceite a papeis claros de produto/tecnologia, se existirem.

### 11. Testes nao acompanharem mudancas de contrato e arquitetura

**Mitigar:** exigir atualizacao de testes junto com mudancas de contrato, adicionar testes de regressao para casos de LLM/fallback e criar testes de integracao quando banco for introduzido.

**Evitar:** evitar concluir itens de contrato ou arquitetura sem evidencia minima em `go test ./...`.

## 2) Pontos de validacao com stakeholders

- Se o uso do MVP sera apenas local/controlado ou em ambiente compartilhado.
- Se dados cadastrados precisam sobreviver a reinicios.
- Se a integracao com LLM sera habilitada em homologacao ou uso real.
- Quais criterios tornam uma prioridade aceitavel para a equipe.
- Quando schema de erro padronizado, filtros, OpenAPI e endpoint dedicado de sugestao devem entrar no escopo.
- Qual cadencia de alinhamento sera usada pela equipe distribuida.
