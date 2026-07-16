# Estratégias de Resposta a Riscos

Este documento apresenta alternativas qualitativas de resposta para os riscos previamente identificados e analisados. As estratégias abaixo têm caráter exploratório e não representam uma escolha definitiva. A seleção final depende de informações adicionais sobre escopo, governança, capacidade da equipe, processos de validação e apetite ao risco dos stakeholders.

## 1) Estratégias Possíveis de Risco

### 1. Mudanças frequentes nos requisitos

**Risco:** Mudanças frequentes nos requisitos

**Descrição:** Os requisitos podem sofrer alterações ao longo do desenvolvimento, afetando funcionalidades planejadas, em andamento ou já implementadas.

**Estratégias de resposta possíveis:**

**Evitar:** Poderia ser aplicada por meio da estabilização prévia de um conjunto mínimo de requisitos antes do desenvolvimento de determinadas funcionalidades. Essa abordagem pode reduzir a exposição ao risco, mas pode ser limitada em contextos nos quais a evolução dos requisitos faz parte da natureza do projeto.

**Mitigar:** Poderia envolver ciclos curtos de validação, priorização frequente, registro formal de mudanças e avaliação de impacto antes de incorporar alterações. Essa estratégia reduz a probabilidade de retrabalho descontrolado e o impacto sobre prazo e esforço.

**Transferir:** Parte do impacto poderia ser compartilhada com stakeholders responsáveis pela definição de requisitos, por meio de acordos sobre prazos de resposta, critérios de aceite e responsabilidades pela aprovação de mudanças. Em projetos com fornecedores externos, cláusulas contratuais de mudança de escopo também poderiam ser consideradas.

**Aceitar:** Pode ser considerada quando a incerteza dos requisitos for esperada e o projeto tiver capacidade para absorver ajustes. Nesse caso, a aceitação deve ser consciente, com reserva de tempo ou capacidade para mudanças.

### 2. Dificuldade de alinhamento sobre escopo e prioridades

**Risco:** Dificuldade de alinhamento sobre escopo e prioridades

**Descrição:** A evolução dos requisitos pode gerar entendimentos diferentes sobre o que deve ser priorizado, o que está dentro do escopo e quais mudanças devem ser incorporadas.

**Estratégias de resposta possíveis:**

**Evitar:** Poderia ser aplicada com a definição antecipada de critérios objetivos de priorização, papéis de decisão e limites claros de escopo para cada ciclo de trabalho. Isso pode reduzir a chance de ambiguidades, mas exige disponibilidade dos stakeholders para decidir e validar.

**Mitigar:** Poderia envolver reuniões regulares de alinhamento, backlog priorizado, registro de decisões e comunicação explícita de alterações de prioridade. Essa abordagem reduz o impacto de interpretações divergentes.

**Transferir:** A responsabilidade por decisões de prioridade pode ser compartilhada com um patrocinador, product owner ou comitê de decisão, caso esses papéis existam. Essa transferência não elimina o risco, mas desloca parte da responsabilidade de decisão para quem tem autoridade sobre escopo e valor de negócio.

**Aceitar:** Pode ser considerada quando o desalinhamento potencial tiver baixo impacto imediato ou quando as decisões puderem ser ajustadas rapidamente. A aceitação exige monitoramento para evitar que pequenas divergências se acumulem.

### 3. Problemas de comunicação em equipe distribuída

**Risco:** Problemas de comunicação em equipe distribuída

**Descrição:** A distribuição da equipe pode dificultar a troca de informações, o compartilhamento de contexto e o esclarecimento de dúvidas durante o desenvolvimento.

**Estratégias de resposta possíveis:**

**Evitar:** Poderia ser aplicada reduzindo dependências críticas entre pessoas ou concentrando decisões sensíveis em momentos sincronizados. Em alguns casos, pode-se evitar atividades distribuídas para tarefas que exigem alta colaboração imediata, mas isso depende da estrutura real da equipe.

**Mitigar:** Poderia envolver canais de comunicação definidos, rituais de sincronização, documentação objetiva de decisões, critérios para escalonamento de bloqueios e registro centralizado de informações relevantes. Essa abordagem reduz atrasos e perda de contexto.

**Transferir:** Parte do risco pode ser compartilhada com lideranças de equipe, facilitadores ou responsáveis por coordenação, atribuindo a eles a responsabilidade por manter alinhamento e remover bloqueios de comunicação. Ferramentas de gestão também podem apoiar, mas não transferem totalmente o risco.

**Aceitar:** Pode ser considerada quando a equipe distribuída já tiver boa autonomia, baixa dependência entre tarefas ou quando pequenas latências de comunicação forem toleráveis. Ainda assim, deve haver acompanhamento para identificar bloqueios recorrentes.

### 4. Coordenação insuficiente entre atividades da equipe

**Risco:** Coordenação insuficiente entre atividades da equipe

**Descrição:** A equipe distribuída pode ter dificuldade para sincronizar atividades, especialmente quando há dependências entre tarefas ou componentes relacionados.

**Estratégias de resposta possíveis:**

**Evitar:** Poderia ser aplicada por meio da redução de dependências entre tarefas, divisão clara de responsabilidades e planejamento de entregas que minimizem trabalho paralelo em componentes altamente acoplados.

**Mitigar:** Poderia envolver planejamento integrado, visibilidade sobre dependências, revisões periódicas de progresso, definição de interfaces entre componentes e integração frequente. Essa estratégia reduz a chance de atrasos por dependência e problemas de integração tardia.

**Transferir:** A coordenação pode ser parcialmente atribuída a líderes técnicos, gerentes de projeto ou responsáveis por integração. Em contextos com fornecedores ou equipes externas, acordos de entrega e pontos de integração podem distribuir responsabilidades.

**Aceitar:** Pode ser considerada quando as atividades forem pouco dependentes entre si ou quando o impacto de uma coordenação imperfeita for limitado. A aceitação deve considerar se atrasos localizados podem ser absorvidos pelo planejamento.

### 5. Impacto das mudanças de requisitos na qualidade do software

**Risco:** Impacto das mudanças de requisitos na qualidade do software

**Descrição:** Alterações nos requisitos podem exigir ajustes rápidos em funcionalidades existentes ou em decisões técnicas já adotadas, afetando a consistência e a estabilidade do software.

**Estratégias de resposta possíveis:**

**Evitar:** Poderia ser aplicada evitando incorporar mudanças de requisitos em funcionalidades consideradas estáveis sem avaliação formal de impacto. Também pode envolver congelamento temporário de escopo em etapas críticas, se isso for aceitável para o projeto.

**Mitigar:** Poderia envolver revisão técnica das mudanças, testes de regressão, critérios mínimos de qualidade, refatoração controlada e validação antes da entrega. Essa abordagem reduz a chance de defeitos e inconsistências decorrentes de alterações.

**Transferir:** Parte do risco pode ser compartilhada com stakeholders que solicitam mudanças, tornando explícitos os impactos sobre prazo, custo, qualidade ou escopo. Em contratos externos, pode-se tratar mudanças como itens sujeitos a replanejamento ou renegociação.

**Aceitar:** Pode ser considerada quando a mudança for pequena, o impacto técnico for limitado ou quando o projeto aceitar um nível maior de instabilidade temporária. Essa aceitação deve ser acompanhada de monitoramento de defeitos e débito técnico.

### 6. Necessidade de validação contínua do entendimento dos requisitos

**Risco:** Necessidade de validação contínua do entendimento dos requisitos

**Descrição:** Como os requisitos evoluem, a equipe pode desenvolver com base em entendimentos desatualizados, incompletos ou não compartilhados de forma uniforme.

**Estratégias de resposta possíveis:**

**Evitar:** Poderia ser aplicada impedindo o início de atividades com requisitos considerados insuficientemente compreendidos ou sem critérios mínimos de aceite. Essa abordagem pode reduzir erros de entendimento, mas pode impactar o ritmo do desenvolvimento.

**Mitigar:** Poderia envolver validações frequentes com stakeholders, refinamento contínuo de requisitos, exemplos de uso, critérios de aceite claros e revisões de entendimento antes da implementação. Essa estratégia reduz a probabilidade de funcionalidades desalinhadas.

**Transferir:** Parte da responsabilidade pode ser compartilhada com stakeholders responsáveis pela validação dos requisitos, estabelecendo prazos e critérios para retorno. Isso não elimina a responsabilidade da equipe técnica, mas torna explícita a dependência de validação externa.

**Aceitar:** Pode ser considerada quando o custo de validação completa for maior que o risco de pequenos ajustes posteriores, especialmente em itens de menor criticidade. A aceitação deve ser limitada a casos em que o retrabalho seja administrável.

## 2) Considerações sobre aplicação das estratégias

### a) Situações em que cada tipo de estratégia tende a ser mais adequada

**Evitar:** Tende a ser mais adequada quando o risco tem impacto alto e existe uma forma viável de remover ou reduzir significativamente a exposição, como estabilizar escopo, adiar itens incertos ou impedir início de atividades sem requisitos mínimos.

**Mitigar:** Tende a ser adequada quando o risco não pode ser eliminado, mas sua probabilidade ou impacto pode ser reduzido por práticas de gestão, comunicação, validação, testes, revisão ou controle de mudanças.

**Transferir:** Tende a ser adequada quando parte do risco depende de decisões, validações ou responsabilidades de terceiros, como stakeholders, fornecedores, lideranças ou papéis formais de negócio. A transferência deve deixar claras responsabilidades, prazos e consequências.

**Aceitar:** Tende a ser adequada quando o risco e inerente ao contexto, quando seu impacto e considerado administrável ou quando o custo de evitar ou mitigar e maior que o beneficio esperado. A aceitação deve ser acompanhada de monitoramento.

### b) Limitações ou trade-offs associados a cada abordagem

**Evitar:** Pode reduzir flexibilidade, atrasar o início de atividades ou exigir maior definição antecipada, o que pode ser difícil em um contexto de requisitos em evolução.

**Mitigar:** Exige disciplina operacional, tempo da equipe e mecanismos consistentes de acompanhamento. Pode reduzir riscos, mas dificilmente os elimina por completo.

**Transferir:** Não remove totalmente o risco do projeto. Mesmo quando responsabilidades são compartilhadas, impactos em prazo, qualidade ou coordenação ainda podem afetar a equipe.

**Aceitar:** Pode ser uma escolha racional em contextos de incerteza, mas aumenta a necessidade de monitoramento. Sem critérios claros, pode se transformar em omissão de tratamento do risco.

## 3) Observações gerais

### a) Dependência de contexto adicional para tomada de decisão

A escolha das respostas depende de informações que não foram fornecidas, como criticidade do software, tamanho e maturidade da equipe, cadencia de entregas, nível de formalidade do processo, disponibilidade dos stakeholders, critérios de qualidade e tolerância a mudanças de escopo.

As sugestões apresentadas devem ser tratadas como alternativas iniciais. A decisão final deve considerar a exposição real de cada risco, a capacidade da equipe de executar as ações e o custo de cada estratégia.

### b) Pontos que exigem validação com stakeholders

- Quem tem autoridade para aprovar mudanças de requisitos e prioridades.
- Qual nível de mudança de escopo e aceitável durante o desenvolvimento.
- Quais critérios mínimos de aceite e qualidade devem ser usados.
- Qual frequência de validação dos requisitos e adequada ao projeto.
- Quais canais e rituais de comunicação devem ser adotados pela equipe distribuída.
- Como impactos de mudanças em prazo, esforço e qualidade serão comunicados e aprovados.
