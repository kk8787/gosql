package jsql

import (
	"testing"
)

func TestLex(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "SELECT 1 FROM users",
			want:  []string{"SELECT", "1", "FROM", "users"},
		},
	}

	for _, test := range tests {
		tokens, err := Lex(test.input)
		if err != nil {
			t.Fatalf("Lex(%q) returned error: %v", test.input, err)
		}

		for _, token := range tokens {
			token.print()
		}
	}
}
