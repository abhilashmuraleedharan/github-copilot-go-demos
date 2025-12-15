// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package filtering

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/demo/cdr-microservice/config"
)

type Service struct {
	config config.FilteringConfig
}

func NewService(cfg config.FilteringConfig) *Service {
	return &Service{
		config: cfg,
	}
}

func (s *Service) ShouldProcess(ctx context.Context, data []byte) bool {
	if !s.config.Enabled {
		return true
	}

	var record map[string]interface{}
	if err := json.Unmarshal(data, &record); err != nil {
		return false
	}

	for _, condition := range s.config.Conditions {
		if s.evaluateCondition(record, condition) {
			return s.config.Action != "drop"
		}
	}

	return s.config.Action == "drop"
}

func (s *Service) evaluateCondition(record map[string]interface{}, rule config.FilterRule) bool {
	fieldValue, ok := record[rule.Field]
	if !ok {
		return false
	}

	switch rule.Operator {
	case "equals":
		return s.compareEqual(fieldValue, rule.Value)
	case "not_equals":
		return !s.compareEqual(fieldValue, rule.Value)
	case "less_than":
		return s.compareLessThan(fieldValue, rule.Value)
	case "greater_than":
		return s.compareGreaterThan(fieldValue, rule.Value)
	case "contains":
		if str, ok := fieldValue.(string); ok {
			return contains(str, rule.Value)
		}
	}

	return false
}

func (s *Service) compareEqual(fieldValue interface{}, ruleValue string) bool {
	switch v := fieldValue.(type) {
	case string:
		return v == ruleValue
	case float64:
		if num, err := strconv.ParseFloat(ruleValue, 64); err == nil {
			return v == num
		}
	case int64:
		if num, err := strconv.ParseInt(ruleValue, 10, 64); err == nil {
			return v == num
		}
	}
	return false
}

func (s *Service) compareLessThan(fieldValue interface{}, ruleValue string) bool {
	switch v := fieldValue.(type) {
	case float64:
		if num, err := strconv.ParseFloat(ruleValue, 64); err == nil {
			return v < num
		}
	case int64:
		if num, err := strconv.ParseInt(ruleValue, 10, 64); err == nil {
			return v < num
		}
	}
	return false
}

func (s *Service) compareGreaterThan(fieldValue interface{}, ruleValue string) bool {
	switch v := fieldValue.(type) {
	case float64:
		if num, err := strconv.ParseFloat(ruleValue, 64); err == nil {
			return v > num
		}
	case int64:
		if num, err := strconv.ParseInt(ruleValue, 10, 64); err == nil {
			return v > num
		}
	}
	return false
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || (len(s) > 0 && indexOfSubstring(s, substr) >= 0))
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
