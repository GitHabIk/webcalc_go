package calculation_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GitHabIk/webcalc_go/pkg/calculation"
)

type TestCase struct {
	Name         string
	RequestBody  string
	ExpectedCode int
	ExpectedBody map[string]string
}

func TestHandleCalculate(t *testing.T) {
	tests := []TestCase{
		{
			Name:         "Valid Expression",
			RequestBody:  `{"expression": "3 + 5 * 2"}`,
			ExpectedCode: http.StatusOK,
			ExpectedBody: map[string]string{"result": "13.000000"},
		},
		{
			Name:         "Invalid Characters in Expression",
			RequestBody:  `{"expression": "3 + a * 2"}`,
			ExpectedCode: http.StatusUnprocessableEntity,
			ExpectedBody: map[string]string{"error": "Expression is not valid"},
		},
		{
			Name:         "Empty Expression",
			RequestBody:  `{"expression": ""}`,
			ExpectedCode: http.StatusUnprocessableEntity,
			ExpectedBody: map[string]string{"error": "Expression is not valid"},
		},
		{
			Name:         "Invalid Request Format",
			RequestBody:  `{"expr": "3 + 5"}`, // некорректное поле
			ExpectedCode: http.StatusUnprocessableEntity,
			ExpectedBody: map[string]string{"error": "Invalid request format"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(tt.RequestBody)))
			if err != nil {
				t.Fatalf("Error creating request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(calculation.HandleCalculate)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.ExpectedCode {
				t.Errorf("Expected status %d, got %d", tt.ExpectedCode, rr.Code)
			}

			var response map[string]string
			if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
				t.Fatalf("Error decoding response: %v", err)
			}

			for key, expectedValue := range tt.ExpectedBody {
				if response[key] != expectedValue {
					t.Errorf("Expected %s: %s, got %s", key, expectedValue, response[key])
				}
			}
		})
	}
}
