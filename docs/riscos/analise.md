# Analise Qualitativa de Riscos

Esta analise considera os riscos identificados em `docs/riscos/identificacao.md` e as implicacoes diretas do contexto informado: projeto de software com requisitos em evolucao, equipe distribuida, arquitetura em camadas, persistencia em memoria, priorizacao assistida por IA/LLM opcional e limitacoes explicitas do MVP. As classificacoes sao qualitativas e devem ser revisitadas conforme ambiente de uso, criticidade dos dados e expectativas dos stakeholders.

## 1) Analise estruturada dos riscos

### 1. Evolucao dos requisitos gerar retrabalho e instabilidade de escopo

**Probabilidade:** Alta

**Impacto:** Alto

**Justificativa:** A evolucao dos requisitos foi informada como condicao do projeto. Mudancas podem afetar escopo, contrato da API, testes, documentacao e decisoes tecnicas ja tomadas.

### 2. Mudancas futuras no contrato da API afetarem consumidores e testes

**Probabilidade:** Media

**Impacto:** Alto

**Justificativa:** Existem evolucoes planejadas de contrato, como erro padronizado, filtros, OpenAPI e possivel endpoint dedicado de sugestao. O impacto cresce se houver consumidores integrados antes da estabilizacao desses contratos.

### 3. Dependencia operacional da integracao opcional com LLM

**Probabilidade:** Media

**Impacto:** Medio

**Justificativa:** A integracao e opcional e possui fallback, o que reduz impacto sobre o CRUD. Ainda assim, quando habilitada, pode afetar latencia, custo e previsibilidade da priorizacao.

### 4. Fallback por heuristica local produzir priorizacoes inadequadas

**Probabilidade:** Media

**Impacto:** Medio

**Justificativa:** A heuristica e deterministica e preserva funcionamento minimo, mas pode nao refletir criterios reais de prioridade da equipe em casos ambiguos.

### 5. Perda de dados por persistencia em memoria

**Probabilidade:** Alta

**Impacto:** Alto

**Justificativa:** A persistencia em memoria e caracteristica atual do MVP. O impacto e alto se houver expectativa de uso continuado ou homologacao com dados que devam sobreviver a reinicios.

### 6. Evolucao para banco de dados exigir mudancas tecnicas relevantes

**Probabilidade:** Media

**Impacto:** Alto

**Justificativa:** A migracao para PostgreSQL esta planejada para etapa futura e pode alterar infraestrutura, testes, tratamento de erro, desempenho e operacao.

### 7. Exposicao da API sem autenticacao ou autorizacao

**Probabilidade:** Baixa

**Impacto:** Alto

**Justificativa:** Autenticacao/autorizacao estao fora do escopo atual. A probabilidade e baixa enquanto o uso for controlado, mas o impacto cresce se a API for exposta fora desse limite.

### 8. Observabilidade insuficiente para diagnostico operacional

**Probabilidade:** Media

**Impacto:** Medio

**Justificativa:** A ausencia de metricas e tracing pode dificultar diagnostico em ambiente compartilhado ou homologacao, principalmente em falhas de LLM e latencia.

### 9. Validacao de entrada nao cobrir todos os valores de dominio

**Probabilidade:** Media

**Impacto:** Medio

**Justificativa:** O risco depende da validacao efetiva no fluxo HTTP. Caso valores invalidos entrem no dominio, podem surgir inconsistencias e comportamento inesperado.

### 10. Coordenacao em equipe distribuida afetar decisoes e sincronizacao

**Probabilidade:** Media

**Impacto:** Medio

**Justificativa:** Equipe distribuida e requisitos em evolucao aumentam necessidade de alinhamento. A classificacao deve ser validada com informacoes sobre rituais, canais e papeis de decisao.

### 11. Testes nao acompanharem mudancas de contrato e arquitetura

**Probabilidade:** Media

**Impacto:** Alto

**Justificativa:** Evolucoes em contrato, validacao, LLM e banco podem introduzir regressao se a suite nao for atualizada junto com as mudancas.

## 2) Matriz Qualitativa de Riscos

| Probabilidade\Impacto | Baixo | Medio | Alto |
|-----------------------|-------|-------|------|
| **Alta** |  |  | R1 - Evolucao dos requisitos<br>R5 - Perda de dados por persistencia em memoria |
| **Media** |  | R3 - Dependencia de LLM<br>R4 - Heuristica local inadequada<br>R8 - Observabilidade insuficiente<br>R9 - Validacao de entrada<br>R10 - Coordenacao distribuida | R2 - Mudancas de contrato<br>R6 - Evolucao para banco<br>R11 - Testes nao acompanharem mudancas |
| **Baixa** |  |  | R7 - Exposicao sem autenticacao/autorizacao |

## 3) Observacoes de incerteza e validacao

- A criticidade de R5, R7 e R8 depende do ambiente real de uso: local, homologacao controlada, rede compartilhada ou producao.
- R2 e R11 devem ser revisitados quando houver decisao sobre schema de erro, filtros, OpenAPI e endpoint dedicado de sugestao.
- R3 e R4 dependem de a integracao com LLM ser habilitada e dos criterios reais de priorizacao aceitos pela equipe.
- R10 requer validacao sobre canais, rituais, responsaveis por decisao e cadencia de alinhamento da equipe distribuida.
