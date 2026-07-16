# Capitulo 1 - Prompts Codex/Copilot
## Prompt 1 - `.gitignore`
```
Contexto: Estou iniciando uma API Python com FastAPI em um repositório de produto.
Objetivo: Gere um arquivo .gitignore para Python, ambiente virtual, cache de testes e configuracoes locais do editor.
Estilo: Organize pro secoes com comentarios.
Resposta: Forneca apenas o conteudo do arquivo .gitignore
```

## Prompt 2 - README inicial
```
Contexto: MVP de micro-API para gestao de tarefas com priorizacao assistida por IA.
Objetivo: Escrever um README inicial com objetivo, stack, como rodar localmente e roadmap de releases.
Estilo: Markdown simples, direto e profissional.
Resposta: Forneca o arquivo README completo.
```

## Prompt 3 - Endpoint de healthcheck
```
Contexto: Projeto em Go 1.25+ com Fiber.
Objetivo: Criar app/main.go com uma instancia Fiber e endpoint GET /health retornando status ok e timestamp.
Estilo: Tipagem e codigo limpo.
Resposta: Forneca apenas o codigo de app/main.go
```

## Prompt 4 - Revisao critica
```
Analise o codigo gerado para app/main.go e responda:
1) Quais riscos tecnicos existem?
2) O que pode quebrar em producao?
3) Quais testes minimos devo criar agora?
Resposta curta em checklist.
```

## Prompt 5 - Mensagem de commit
```
Contexto: Adicionei estrutura inicial, README, .gitignore e endpoint /health. Veja tambem outras alteracoes que nao estao aqui listadas.
Objetivo: Gerar uma mensagem de commit no padrao Conventional Commits.
Resposta: Apenas uma linha de commit.
```

# Capitulo 2 - Prompts para Codex/Copilot
## Prompt 1 - Escopo MVP
```
Contexto: MVP de micro-API de tarefas para uso de equipe interna
Objetivo: Gerar documento de escopo com objetivo, requisitos funcionais, nao funcionais e fora de escopo
Estilo: Linguagem tecnica, direta, em Markdown
Resposta: Forneca o conteudo completo de docs/escopo-mvp.md
```

## Prompt 2 - Backlog por releases
```
Contexto: O produto sera entregue em 3 releases: core, qualidade e entrega final
Objetivo: Criar backlog minimo com IDs RF/RT e criterios de aceite
Estilo: Checklist Markdown
Resposta: Conteudo de docs/backlog.md
```

## Prompt 3 - Arquitetura Mermaid
```
Contexto: Fiber com camadas API, Service, Repository e componente PriorityAdvisor
Objetivo: Gerar diagramas Mermaid de componentes e fluxo de dados
Estilo: Simples, legivel e versionavel
Resposta: Apenas bloco Mermaid com o arquivo no repositorio
```

## Prompt 4 - Conventional Commits
```
Contexto: Adicionei docs/escopo-mvp.md, docs/arquitetura.md e docs/backlog.md
Objetivo: Sugerir 3 mensagens de commit no padrao Conventional Commits
Resposta: Apenas as 3 linhas de commit
```

## Prompt 5 - Revisao de planejamento
```
> Analise este escopo e backlog e responda:</br>
> 1) O que esta grande demais para a release inicial?</br>
> 2) O que esta faltando para testabilidade?</br>
> 3) Quais riscos tecnicos devo mitigar antes da implementacao?</br>
> Resposta em bullets curtos.
```

# Capitulo 3 - Prompts para Codex/Copilot
## Prompt 1 - Modelo Struct
```
Contexto: Micro-API de gerenciamento de tarefas em Go para uso interno de equipe.
> Stack: Go 1.22+, encoding/json da stdlib, validação com github.com/go-playground/validator/v10.
Objetivo: Gerar os modelos Task, TaskCreate, TaskUpdate e TaskOut, equivalentes aos modelos Pydantic do projeto, usando structs, tags json e validate.
Estilo: Go idiomático, limpo, sem código desnecessário, pronto para compilação.
Resposta: Apenas código do arquivo app/models/task.go.
```
## Prompt 2 - Repositorio Inicial
```
Contexto: Preciso de persistencia inicial enxuta para viabilizar a primeira release.
Objetivo: Criar TaskRepository em memoria com create, list, get_by_id, update e delete.
Estilo: Go tipado, sem dependencias externas.
Resposta: Codigo completo de app/repositories/task_repository.go
```
## Prompt 3 - Service com regra de prioridade
```
Contexto: A prioridade da tarefa pode ser sugerida automaticamente.
Objetivo: Criar TaskService que use TaskRepository e PriorityAdvisor.
Estilo: Separar regra de negocio da camada de API
Resposta: Codigo de app/services/task_service.go
```

## Prompt 4 - PriorityAdvisor com fallback
```
Contexto: Quero rodar sem custo de API quando nao houver chave.
Objetivo: Implementar PriorityAdvisor com heuristica local e chamada opcional a LLM quando OPENAI_API_KEY existir.
Estilo: Falha segura, timeout e fallback obrigatorio.
Resposta: Codigo de app/services/priority_advisor.go
```

## Prompt 5 - Rotas CRUD
```
Contexto: Fiber com TaskService pronto.
Objetivo: Criar rotas POST/GET/PUT/DELETE para tarefas com status HTTP corretos e tratamento de 404.
Estilo: Router separado em app/api/task_routes.go.
Resposta: Apenas o codigo de arquivo.
```

## Prompt 6 - Revisao Tecnica
```
Revise os arquivos do core da API e responda:
1) Quais pontos de acoplamento estao altos?
2) Onde faltam validacoes?
3) Quais 5 testes devo priorizar na proxima release?
Resposta em checklist
```

# Capitulo 4 - Prompts para Codex/Copilot
## Prompt 1 - Testes de service
```
Contexto: Tenho TaskService com CRUD de tarefas.
Objetivo: Gerar suite de testes com uma biblioteca semelhante ao Pytest mas para Go, cobrindo criacao, listagem, atualizacao, exclusao e caso de erro por ID inexistente.
Estilo: Testes claros, nomes descritivos e fixtures simples.
Resposta: Codigo completo de testes no arquivo tests/task_service_test.go.
```
## Prompt 2 - Testes do PriorityAdvisor
```
Contexto: PriorityAdvisor possui heuristica local e fallback quando a chamada externa falha.
Objetivo: Gerar testes para os tres niveis de prioridade e para fallback.
Estilo: Usar monkeypatch quando necessario. Se monkeypatch so existir para python, usar a ferramenta equivalente para Go.
Resposta: Codigo completo de testes no arquivo tests/priority_advisor_test.go.
```
## Prompt 3 - Testes de API
```
Contexto: API Fiber com endpoints CRUD de /tasks.
Objetivo: Criar testes de rota com TestClient para status 201, 200, 204 e 404.
Estilo: Isolar dependencia de repositorio para evitar estado global entre testes.
Resposta: Codigo de tests/task_routes_test.go
```

## Prompt 4 - Refatoração DRY/SRP
```
Contexto: Analise os arquivos app/services/task_service.go e app/repositories/task_repository.go.
Objetivo: Sugerir refatoracao com foco em DRY e SRP sem mudar comportamento externo.
Resposta: 1) lista de mudancas propostas e 2) patch sugerido por arquivo.
```

## Prompt 5 - README final tecnico
```
Contexto: MVP de micro-API de tarefas com prioridade assistida por IA.
Objetivo: Atualizar o arquivo README atual, com o conteudo completo com instalacao, execucao, testes, arquitetura, uso da IA, limitacoes e proximos passos.
Estilo: Markdown profissional e objetivo.
Resposta: README inteiro.
```

## Prompt 6 - Revisao final de qualidade
```
Com base no codigo e nos testes atuais, gere um checklist com:
1) Riscos tecnicos restantes
2) Gaps de cobertura de teste
3) Melhorias prioritarias para a proxima release
Resposta em bullets curtos.
```

# Capitulo 5 - Prompts para Codex/Copilot
## Prompt 1 - Makefile
```
Contexto: Projeto Fiber com comandos de instalar dependencias, executar API e rodar testes.
Objetivo: Gerar arquivo Makefile com targets install, run e test.
Estilo: Simples e portavel.
Resposta: Conteudo completo de Makefile.
```
## Prompt 2 - `.env.example`
```
Contexto: O projeto usa integracao opcional de LLM.
Objetivo: Criar arquivo de exemplo do .env que ja existe no repositorio, com variaveis necessarias e valores placeholder seguros.
Estilo: Minimalista.
Resposta: Apenas conteudo do arquivo.
```
## Prompt 3 - Revisao de README
```
Analise meu README e responda:
1) O que falta para ser reproduzivel em maquina limpa?
2) Quais secoes estao fracas para onboard tecnico?
3) Como melhorar a secao de uso da IA?
Resposta em checklist objetivo.
Atualize para que tenha essas correcoes
```

## Prompt 4 - Checklist de release
```
Contexto: Vou publicar a versão v1.0.0 no GitHub.
Objetivo: Gerar checklist final com itens tecnicos e de documentacao para release.
Estilo: Checkbox Markdown.
Resposta: Lista completa
```

## Prompt 5 - Script de demo tecnica
```
Contexto: Preciso apresentar o MVP em 5 minutos para stakeholders tecnicos.
Objetivo: Criar roteiro com tempos por bloco e comandos para mostrar no terminal.
Estilo: Direto, sem marketing.
Resposta: Roteiro em Markdown num arquivo em docs, mas esse arquivo não pode subir para o GitHub.
```

## Prompt 6 - Revisao final de qualidade
```
Com base no codigo e nos testes atuais, gere um checklist com:
1) Riscos tecnicos restantes
2) Gaps de cobertura de teste
3) Melhorias prioritarias para a proxima release
Resposta em bullets curtos.
```

# Capitulo 6 - Prompts de Gerenciamento de Riscos e Comunicacao
## Prompt 1 - Identificacao de Riscos
```
Persona: Voce atua como um especialista em gerenciamento de riscos em projetos de software, com experiencia na identificacao de riscos tecnicos, organizacionais e de processo, especialmente em contextos de alta incerteza.
Tarefa: Identificar possiveis riscos associados ao projeto, com base nas informacoes fornecidas, considerando fatores que podem impactar o andamento, a qualidade ou os resultados do projeto.
Contexto: Considere que:
1) Trata-se de um projeto de desenvolvimento de software;
2) Os requisitos estão em evolucao, podendo sofrer mudancas ao longo do tempo;
3) A equipe e distribuida, podendo envolver desafios de comunicacao e coordenacao;
A analise deve se limitar as informacoes fornecidas e suas implicacoes diretas.
Saida: Apresente o resultado em formato estruturado, conforme o modelo abaixo:
Lista de riscos identificados. Para cada risco:
1) Risco: descricao clara e objetiva;
2) Descricao: explicacao breve do risco;
3) Contexto de ocorrencia: situacao ou condicao em que o risco pode se manifestar.
Restricoes e diretrizes: Nao assumir informacoes que nao foram fornecidas explicitamente; Indicar claramente quando houver incerteza ou necessidade de validacao; Evitar generalizacoes excessivas ou riscos genericos.
```
## Prompt 1 - Analise de Riscos
```
Persona: Voce atua como um analista de riscos em projetos de software, com experiencia na analise qualitativa de riscos e aplicacao de matrizes de probabilidade e impacto.
Tarefa: Realizar uma analise dos riscos previamente identificados, descrevendo seus impactos e fatores associados, e posteriormente organiza-los em uma matriz qualitativa de probabilidade vs. impacto, para apoiar a compreensao dos riscos.
Contexto: Considere que:
1) Os riscos ja foram previamente identificados;
2) O projeto esta em contexto de desenvolvimento de software;
3) As informacoes disponiveis podem variar em nivel de detalhe;
A analise deve se limitar as informacoes fornecidas e suas implicacoes diretas. Pode considerar o output feito no arquivo identificacao.md.
Saida: Apresente o resultado em formato estruturado, conforme o modelo abaixo:
1) Analise estruturada dos riscos. Para cada risco:
a) Risco: nome ou descricao breve;
b) Descricao: detalhamento do risco;
c) Possiveis impactos no projeto: impactos potenciais em prazo, esforco, qualidade ou coordenacao.
d) Fatores que influenciam a ocorrencia: condicoes, dependencias ou caracteristicas do projeto que aumentam ou reduzem a probabilidade de ocorrencia.
e) Probabilidade (qualitativa): Baixa / Media / Alta.
f) Impacto (qualitativo): Baixo / Medio / Alto.
g) Justificativa da classificacao
2) Matriz Qualitativa de Riscos
|Probabilidade\Impacto|Baixo|Medio|Alto|
|---------------------|-----|-----|----|
|**Alta**       |   |   |   |
|**Media**      |   |   |   |
|**Baixa**      |   |   |   |
Restricoes e diretrizes: Nao atribuir valores numericos de probabilidade ou impacto; Indicar claramente quando houver incerteza ou necessidade de validacao.
Saida: Escreva o output num arquivo Markdown nesse diretorio: docs/riscos/analise.md
```
## Prompt 3 - Definicao das Estrategias de Resposta
```
Persona: Voce atua como um gerente de projetos com experiencia em resposta a riscos, com conhecimento em estrategias classicas de tratamento de riscos em projetos de software, como evitar, mitigar, transferir e aceitar riscos, considerando diferentes contextos e niveis e incerteza.
Tarefa: Sugerir possiveis estrategias de resposta para os riscos previamente identificados e analisados, explorando alternativas viaveis e suas implicacoes, sem realizar um escolha definitiva.
Contexto: Considere que:
1) Os requisitos ja foram identificados e analisados previamente;
2) O projeto esta em contexto de desenvolvimento de software;
3) As informacoes disponiveis podem variar em nivel de detalhe;
4) A analise servira como apoio a tomada de decisao posterior.
Saida: Apresente o resultado em formato estruturado, conforme o modelo abaixo:
1) Estrategias Possiveis de Risco. Para cada risco:
a) Risco: nome ou descricao breve;
b) Descricao: detalhamento do risco;
c) Estrategias de resposta possiveis:
i) Evitar: Descricao de como esssa abordagem poderia ser aplicada ao risco;
ii) Mitigar: Descricao de como reduzir a probabilidade ou impacto;
iii) Transferir: Descricao de como o risco poderia ser compartilhado ou transferido;
iv) Aceitar: Descricao de quando e por que essa aborsagem pode ser considera.
2) Consideracoes sobre aplicacao das estrategias
a) Situacoes em que cada tipode estrategia tende a ser mais adequada;
b) Limitacoes ou trade-offs associados a cada abordagem;
3) Observacoes gerais:
a) Dependencia de contexto adicional para tomada de decisao;
b) Pontos que exigem validacao com stakeholders.
Restricoes e diretrizes: Indicar claramente as limitacoes das sugestoes apresentadas; Manter a analise em nivel exploratorio e qualitativo.
Saida: Escreva o output num arquivo Markdown nesse diretorio: docs/riscos/respostas.md
```
