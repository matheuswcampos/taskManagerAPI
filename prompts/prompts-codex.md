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