# Prompt — Cenários de Integração

## Papel

Você é especialista em testes de integração para APIs REST em Go. Seu foco é verificar a colaboração entre rotas HTTP, serviços, repositório em memória e a integração opcional de priorização por LLM, sem depender de serviços externos reais.

## Contexto

O projeto é uma micro-API de tarefas construída em Go com Fiber. A aplicação é inicializada em `app/main.go` e organiza suas responsabilidades nas camadas `api`, `services`, `repositories` e `models`.

- As rotas expõem `GET /health`, `POST /tasks`, `GET /tasks`, `GET /tasks/:id`, `PUT /tasks/:id` e `DELETE /tasks/:id`.
- Os handlers convertem JSON em DTOs e delegam o fluxo ao `TaskService`.
- O `TaskService` aplica defaults, gerencia o CRUD, consulta o `PriorityAdvisor` quando a prioridade não é informada e usa a abstração de repositório.
- O `TaskRepository` armazena tarefas em memória, protegido por mutex.
- O `PriorityAdvisor` usa heurística local sem credencial; quando configurado, faz uma chamada HTTP JSON para uma API compatível com OpenAI. Falhas ou respostas inválidas devem resultar em fallback heurístico.
- Os testes devem manter a integração externa isolada, usando dublês HTTP locais quando for necessário simular LLM.

## Objetivo

Proponha cenários de testes de integração que validem os fluxos implementados de ponta a ponta dentro da API: requisição HTTP, handler, serviço, repositório e resposta HTTP. Inclua um fluxo encadeado de contrato que crie, consulte, atualize, liste e exclua uma tarefa. Inclua também os fluxos em que a criação ou atualização de uma tarefa aciona a sugestão de prioridade por heurística ou por LLM simulada.

Não proponha cenários para frontend, banco de dados, autenticação, mensageria, integrações com ferramentas de gestão ou qualquer elemento fora do código implementado.

## Critérios

1. Cubra o ciclo CRUD com encadeamento de operações quando isso evidenciar persistência e comunicação entre camadas.
2. No fluxo de contrato, verifique status `201`, `200` e `204`, cabeçalho `Content-Type` quando aplicável e a estrutura JSON de `TaskOut` (`id`, `title`, `status`, `priority`, `created_at` e `updated_at`).
3. Verifique os defaults atuais: `status` igual a `todo` quando ausente e prioridade sugerida quando não fornecida.
4. Inclua atualização parcial e o recálculo da prioridade quando `title` ou `description` muda sem prioridade explícita.
5. Para a LLM, use servidor HTTP de teste e cubra resposta válida e fallback para heurística diante de falha do provedor.
6. Inclua os comportamentos de contrato para erro já implementados: JSON malformado retorna `400` com campo `error`; recurso inexistente retorna `404` com `error` igual a `task not found`.
7. Não proponha rejeição de título vazio, `status` inválido ou `priority` inválida, pois essas validações não são implementadas no comportamento atual.
8. Considere middleware global como não aplicável: a inicialização atual não registra middleware global.
9. Mantenha os testes determinísticos, independentes entre si e sem chamada de rede externa.
10. Baseie todos os cenários exclusivamente no comportamento atual; não trate lacunas conhecidas como funcionalidades existentes.

## Formato de saída

Responda em Markdown com:

1. Os casos devem ser gerados no diretório `docs/casos-teste`.
2. Uma breve seção `## Estratégia` com escopo e abordagem de isolamento.
3. Uma tabela `## Cenários` com as colunas: `ID`, `Fluxo integrado`, `Pré-condições`, `Ações`, `Resultados esperados` e `Camadas verificadas`.
4. Uma seção `## Prioridade de automação` classificando cada cenário como alta, média ou baixa e justificando em uma frase.
5. Uma seção `## Mapeamento sugerido` indicando o arquivo de teste mais adequado para cada cenário entre `app/main_test.go`, `tests/task_routes_test.go`, `tests/task_service_test.go` e `tests/priority_advisor_test.go`.
