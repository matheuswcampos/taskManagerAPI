# Análise Qualitativa de Riscos

Esta análise considera apenas os riscos previamente identificados em `docs/riscos/identificacao.md` e as implicações diretas do contexto informado: projeto de desenvolvimento de software, requisitos em evolução e equipe distribuída. As classificações são qualitativas e devem ser validadas conforme mais informações sobre processo, equipe, governança de requisitos e cadência de comunicação estejam disponíveis.

## 1) Análise estruturada dos riscos

### 1. Mudanças frequentes nos requisitos

**Risco:** Mudanças frequentes nos requisitos

**Descrição:** Os requisitos podem sofrer alterações ao longo do desenvolvimento, afetando funcionalidades planejadas, em andamento ou já implementadas.

**Possíveis impactos no projeto:** Pode gerar retrabalho, aumento de esforço, alterações no prazo, dificuldade de estabilizar o escopo e necessidade de revisar decisões já tomadas.

**Fatores que influenciam a ocorrência:** A ocorrência tende a aumentar quando os requisitos ainda não estão suficientemente validados ou quando novas necessidades surgem durante o desenvolvimento. A probabilidade pode ser reduzida se houver mecanismos claros de priorização, validação e controle de mudanças, mas essas práticas não foram informadas e precisam ser validadas.

**Probabilidade (qualitativa):** Alta

**Impacto (qualitativo):** Alto

**Justificativa da classificação:** A probabilidade foi classificada como alta porque o contexto informa explicitamente que os requisitos estão em evolução. O impacto foi classificado como alto porque mudanças de requisitos podem afetar prazo, esforço, qualidade e escopo, especialmente quando atingem itens já desenvolvidos ou em desenvolvimento.

### 2. Dificuldade de alinhamento sobre escopo e prioridades

**Risco:** Dificuldade de alinhamento sobre escopo e prioridades

**Descrição:** A evolução dos requisitos pode gerar entendimentos diferentes sobre o que deve ser priorizado, o que está dentro do escopo e quais mudanças devem ser incorporadas.

**Possíveis impactos no projeto:** Pode causar decisões inconsistentes, retrabalho, atraso em definições, conflitos de prioridade e execução de atividades com base em premissas diferentes.

**Fatores que influenciam a ocorrência:** O risco aumenta quando a equipe precisa tomar decisões antes de haver clareza sobre requisitos e prioridades. Também pode aumentar em equipes distribuídas, caso as decisões não sejam comunicadas de forma clara. Há incerteza sobre a existência de um processo formal de priorização, o que requer validação.

**Probabilidade (qualitativa):** média

**Impacto (qualitativo):** Alto

**Justificativa da classificação:** A probabilidade foi classificada como média porque o risco decorre diretamente da evolução dos requisitos, mas depende de como as decisões de escopo são conduzidas. O impacto foi considerado alto porque desalinhamentos de escopo e prioridade podem afetar prazo, esforço e coordenação da equipe.

### 3. Problemas de comunicação em equipe distribuída

**Risco:** Problemas de comunicação em equipe distribuída

**Descrição:** A distribuição da equipe pode dificultar a troca de informações, o compartilhamento de contexto e o esclarecimento de dúvidas durante o desenvolvimento.

**Possíveis impactos no projeto:** Pode gerar atrasos na resolução de bloqueios, entendimento incompleto de requisitos, decisões técnicas desalinhadas e redução da coordenação entre os envolvidos.

**Fatores que influenciam a ocorrência:** O risco aumenta quando há dependência frequente entre membros da equipe para esclarecer requisitos, tomar decisões ou integrar atividades. Pode ser reduzido por rotinas de comunicação, registros de decisão e canais bem definidos, mas essas informações não foram fornecidas e precisam ser validadas.

**Probabilidade (qualitativa):** média

**Impacto (qualitativo):** médio

**Justificativa da classificação:** A probabilidade foi classificada como média porque a equipe distribuída foi informada como uma condição do projeto, mas não há detalhes sobre fusos, ferramentas ou cadência de comunicação. O impacto foi considerado médio porque falhas de comunicação podem afetar coordenação e prazo, embora sua gravidade dependa da intensidade das dependências entre atividades.

### 4. coordenação insuficiente entre atividades da equipe

**Risco:** coordenação insuficiente entre atividades da equipe

**Descrição:** A equipe distribuída pode ter dificuldade para sincronizar atividades, especialmente quando há dependências entre tarefas ou componentes relacionados.

**Possíveis impactos no projeto:** Pode gerar duplicidade de esforço, atrasos por dependências não resolvidas, integração tardia de componentes e inconsistências entre partes do software.

**Fatores que influenciam a ocorrência:** O risco aumenta quando diferentes membros ou subgrupos trabalham em itens relacionados sem alinhamento frequente. A probabilidade e o impacto dependem do nível de interdependência técnica entre as atividades, informação que não foi fornecida e precisa ser validada.

**Probabilidade (qualitativa):** média

**Impacto (qualitativo):** médio

**Justificativa da classificação:** A probabilidade foi classificada como média porque a distribuição da equipe cria uma condição favorável ao risco, mas não confirma sua ocorrência. O impacto foi classificado como médio porque os efeitos tendem a afetar coordenação, prazo e integração, mas sua severidade depende de dependências técnicas ainda não detalhadas.

### 5. Impacto das mudanças de requisitos na qualidade do software

**Risco:** Impacto das mudanças de requisitos na qualidade do software

**Descrição:** Alterações nos requisitos podem exigir ajustes rápidos em funcionalidades existentes ou em decisões técnicas já adotadas, afetando a consistência e a estabilidade do software.

**Possíveis impactos no projeto:** Pode reduzir a qualidade da implementação, enfraquecer a cobertura de validação, introduzir defeitos, aumentar retrabalho e dificultar a manutenção do produto.

**Fatores que influenciam a ocorrência:** O risco aumenta quando mudanças são incorporadas sem tempo adequado para análise, implementação e validação. Pode ser reduzido por práticas de revisão, testes e controle de mudanças, mas a existência dessas práticas não foi informada e requer validação.

**Probabilidade (qualitativa):** média

**Impacto (qualitativo):** Alto

**Justificativa da classificação:** A probabilidade foi classificada como média porque depende de como as mudanças de requisitos são tratadas durante o desenvolvimento. O impacto foi classificado como alto porque problemas de qualidade podem afetar diretamente os resultados do projeto, a estabilidade do software e o esforço de correção.

### 6. Necessidade de validação contínua do entendimento dos requisitos

**Risco:** Necessidade de validação contínua do entendimento dos requisitos

**Descrição:** Como os requisitos evoluem, a equipe pode desenvolver com base em entendimentos desatualizados, incompletos ou não compartilhados de forma uniforme.

**Possíveis impactos no projeto:** Pode gerar funcionalidades desalinhadas, retrabalho, atrasos por necessidade de correção e inconsistências entre o que foi desenvolvido e o que era esperado.

**Fatores que influenciam a ocorrência:** O risco aumenta quando não há mecanismos frequentes de validação do entendimento dos requisitos. Como a existência e a frequência desses mecanismos não foram informadas, há necessidade de validação adicional.

**Probabilidade (qualitativa):** Alta

**Impacto (qualitativo):** Médio

**Justificativa da classificação:** A probabilidade foi classificada como alta porque a evolução dos requisitos torna necessária a validação contínua do entendimento. O impacto foi classificado como médio porque o risco pode gerar retrabalho e desalinhamento, mas sua gravidade depende da rapidez com que erros de entendimento são identificados e corrigidos.

## 2) Matriz Qualitativa de Riscos

| Probabilidade\Impacto | Baixo | Médio | Alto |
|-----------------------|-------|-------|------|
| **Alta** |  | R6 - Necessidade de validação contínua do entendimento dos requisitos | R1 - Mudanças frequentes nos requisitos |
| **Média** |  | R3 - Problemas de comunicação em equipe distribuída<br>R4 - Coordenação insuficiente entre atividades da equipe | R2 - Dificuldade de alinhamento sobre escopo e prioridades<br>R5 - Impacto das mudanças de requisitos na qualidade do software |
| **Baixa** |  |  |  |

## Observações de incerteza e validação

- As classificações devem ser validadas com informações sobre processo de gestão de requisitos, critérios de priorização e controle de mudanças.
- A análise da equipe distribuída depende de detalhes não fornecidos, como fusos horários, canais de comunicação, frequência de sincronização e nível de dependência entre atividades.
- O impacto sobre qualidade depende da existência de práticas de revisão, testes e validação, que não foram informadas.
