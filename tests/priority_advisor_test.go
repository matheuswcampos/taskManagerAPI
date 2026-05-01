package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"task-prioritization-api/app/models"
	"task-prioritization-api/app/services"
)

func TestPriorityAdvisor_HeuristicLevels(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("OPENAI_BASE_URL", "")
	t.Setenv("OPENAI_MODEL", "")

	advisor := services.NewPriorityAdvisor()

	tests := []struct {
		name        string
		title       string
		description string
		want        models.TaskPriority
	}{
		{
			name:        "low priority from backlog wording",
			title:       "Pesquisa opcional",
			description: "Item de backlog para estudo",
			want:        models.PriorityLow,
		},
		{
			name:        "medium priority from maintenance hint",
			title:       "Melhoria de texto",
			description: "Aprimorar mensagens de retorno",
			want:        models.PriorityMedium,
		},
		{
			name:        "high priority from urgency hint",
			title:       "Urgente para cliente",
			description: "Solicitacao importante para entrega",
			want:        models.PriorityHigh,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := advisor.SuggestPriority(tc.title, tc.description)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tc.want {
				t.Fatalf("expected priority %q, got %q", tc.want, got)
			}
		})
	}
}

func TestPriorityAdvisor_FallbackWhenExternalCallFails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("PRIORITY_ADVISOR_ENABLE_LLM_IN_TESTS", "1")
	t.Setenv("OPENAI_BASE_URL", server.URL)
	t.Setenv("OPENAI_MODEL", "gpt-4.1-mini")

	advisor := services.NewPriorityAdvisor()

	title := "Urgente para cliente"
	description := "Solicitacao importante para entrega"

	got, err := advisor.SuggestPriority(title, description)
	if err != nil {
		t.Fatalf("expected no error on fallback path, got %v", err)
	}

	// External call fails, so advisor must return local heuristic.
	if got != models.PriorityHigh {
		t.Fatalf("expected fallback priority %q, got %q", models.PriorityHigh, got)
	}
}

func TestPriorityAdvisor_UsesLLMWhenValidResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"priority\":\"critic\"}"}}]}`)
	}))
	defer server.Close()

	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("PRIORITY_ADVISOR_ENABLE_LLM_IN_TESTS", "1")
	t.Setenv("OPENAI_BASE_URL", server.URL)
	t.Setenv("OPENAI_MODEL", "gpt-4.1-mini")

	advisor := services.NewPriorityAdvisor()

	got, err := advisor.SuggestPriority("titulo neutro", "descricao neutra")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != models.PriorityCritic {
		t.Fatalf("expected llm priority %q, got %q", models.PriorityCritic, got)
	}
}

func TestPriorityAdvisor_FallbackWhenLLMReturnsInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"choices":[{"message":{"content":"not-a-json"}}]}`)
	}))
	defer server.Close()

	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("PRIORITY_ADVISOR_ENABLE_LLM_IN_TESTS", "1")
	t.Setenv("OPENAI_BASE_URL", server.URL)
	t.Setenv("OPENAI_MODEL", "gpt-4.1-mini")

	advisor := services.NewPriorityAdvisor()
	got, err := advisor.SuggestPriority("Urgente para cliente", "Solicitacao importante para entrega")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != models.PriorityHigh {
		t.Fatalf("expected heuristic fallback %q, got %q", models.PriorityHigh, got)
	}
}

func TestPriorityAdvisor_FallbackWhenLLMReturnsEmptyChoices(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"choices":[]}`)
	}))
	defer server.Close()

	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("PRIORITY_ADVISOR_ENABLE_LLM_IN_TESTS", "1")
	t.Setenv("OPENAI_BASE_URL", server.URL)
	t.Setenv("OPENAI_MODEL", "gpt-4.1-mini")

	advisor := services.NewPriorityAdvisor()
	got, err := advisor.SuggestPriority("Urgente para cliente", "Solicitacao importante para entrega")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != models.PriorityHigh {
		t.Fatalf("expected heuristic fallback %q, got %q", models.PriorityHigh, got)
	}
}

func TestPriorityAdvisor_FallbackWhenLLMReturnsInvalidPriority(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"choices":[{"message":{"content":"{\"priority\":\"p0\"}"}}]}`)
	}))
	defer server.Close()

	t.Setenv("OPENAI_API_KEY", "test-key")
	t.Setenv("PRIORITY_ADVISOR_ENABLE_LLM_IN_TESTS", "1")
	t.Setenv("OPENAI_BASE_URL", server.URL)
	t.Setenv("OPENAI_MODEL", "gpt-4.1-mini")

	advisor := services.NewPriorityAdvisor()
	got, err := advisor.SuggestPriority("Urgente para cliente", "Solicitacao importante para entrega")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != models.PriorityHigh {
		t.Fatalf("expected heuristic fallback %q, got %q", models.PriorityHigh, got)
	}
}
