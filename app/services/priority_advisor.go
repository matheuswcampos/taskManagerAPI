package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"task-prioritization-api/app/models"
)

const (
	defaultOpenAIBaseURL = "https://api.openai.com/v1"
	defaultOpenAIModel   = "gpt-4.1-mini"
	defaultTimeout       = 4 * time.Second
)

// DefaultPriorityAdvisor suggests task priority using local heuristics
// and optionally an LLM call when OPENAI_API_KEY is configured.
type DefaultPriorityAdvisor struct {
	apiKey  string
	baseURL string
	model   string
	timeout time.Duration
	client  *http.Client
}

// NewPriorityAdvisor creates a fail-safe priority advisor.
func NewPriorityAdvisor() *DefaultPriorityAdvisor {
	apiKey := strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))
	baseURL := strings.TrimSpace(os.Getenv("OPENAI_BASE_URL"))
	if baseURL == "" {
		baseURL = defaultOpenAIBaseURL
	}

	model := strings.TrimSpace(os.Getenv("OPENAI_MODEL"))
	if model == "" {
		model = defaultOpenAIModel
	}

	timeout := defaultTimeout
	if raw := strings.TrimSpace(os.Getenv("PRIORITY_ADVISOR_TIMEOUT")); raw != "" {
		if parsed, err := time.ParseDuration(raw); err == nil && parsed > 0 {
			timeout = parsed
		}
	}

	return &DefaultPriorityAdvisor{
		apiKey:  apiKey,
		baseURL: strings.TrimRight(baseURL, "/"),
		model:   model,
		timeout: timeout,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// SuggestPriority returns a safe priority suggestion.
// It always falls back to local heuristics when LLM is unavailable.
func (a *DefaultPriorityAdvisor) SuggestPriority(title, description string) (models.TaskPriority, error) {
	fallback := heuristicPriority(title, description)

	if a == nil || strings.TrimSpace(a.apiKey) == "" {
		return fallback, nil
	}

	priority, err := a.suggestWithLLM(title, description)
	if err != nil {
		return fallback, nil
	}

	if !isValidPriority(priority) {
		return fallback, nil
	}

	return priority, nil
}

func (a *DefaultPriorityAdvisor) suggestWithLLM(title, description string) (models.TaskPriority, error) {
	endpoint := a.baseURL + "/chat/completions"

	payload := map[string]any{
		"model":       a.model,
		"temperature": 0,
		"response_format": map[string]string{
			"type": "json_object",
		},
		"messages": []map[string]string{
			{
				"role": "system",
				"content": "Classifique prioridade de tarefa para backlog interno. " +
					"Retorne apenas JSON no formato {\"priority\":\"low|medium|high|critic\"}.",
			},
			{
				"role": "user",
				"content": fmt.Sprintf("title: %s\ndescription: %s", strings.TrimSpace(title), strings.TrimSpace(description)),
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("openai status: %d", resp.StatusCode)
	}

	var llmResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&llmResp); err != nil {
		return "", err
	}
	if len(llmResp.Choices) == 0 {
		return "", errors.New("empty choices")
	}

	content := strings.TrimSpace(llmResp.Choices[0].Message.Content)
	if content == "" {
		return "", errors.New("empty content")
	}

	var out struct {
		Priority string `json:"priority"`
	}
	if err := json.Unmarshal([]byte(content), &out); err != nil {
		return "", err
	}

	priority := models.TaskPriority(strings.ToLower(strings.TrimSpace(out.Priority)))
	if !isValidPriority(priority) {
		return "", errors.New("invalid priority")
	}

	return priority, nil
}

func heuristicPriority(title, description string) models.TaskPriority {
	text := strings.ToLower(strings.TrimSpace(title + " " + description))
	if text == "" {
		return models.PriorityMedium
	}

	score := 0

	criticHints := []string{"critico", "critical", "incidente", "outage", "security", "seguranca", "vazamento", "produção", "producao", "bloqueia"}
	highHints := []string{"urgente", "asap", "prazo", "deadline", "cliente", "bug", "falha", "erro", "alto impacto"}
	mediumHints := []string{"melhoria", "improvement", "refactor", "ajuste", "ajustar", "manutencao"}
	lowHints := []string{"backlog", "opcional", "nice to have", "estudo", "spike", "pesquisa"}

	for _, hint := range criticHints {
		if strings.Contains(text, hint) {
			score += 4
		}
	}
	for _, hint := range highHints {
		if strings.Contains(text, hint) {
			score += 2
		}
	}
	for _, hint := range mediumHints {
		if strings.Contains(text, hint) {
			score += 1
		}
	}
	for _, hint := range lowHints {
		if strings.Contains(text, hint) {
			score -= 1
		}
	}

	switch {
	case score >= 6:
		return models.PriorityCritic
	case score >= 3:
		return models.PriorityHigh
	case score >= 1:
		return models.PriorityMedium
	default:
		return models.PriorityLow
	}
}

func isValidPriority(p models.TaskPriority) bool {
	switch p {
	case models.PriorityLow, models.PriorityMedium, models.PriorityHigh, models.PriorityCritic:
		return true
	default:
		return false
	}
}
