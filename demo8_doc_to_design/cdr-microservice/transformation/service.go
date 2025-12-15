// [AI GENERATED] LLM: GitHub Copilot, Mode: Chat, Date: 2025-12-15
package transformation

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/demo/cdr-microservice/config"
)

type Service struct {
	config config.TransformationConfig
}

func NewService(cfg config.TransformationConfig) *Service {
	return &Service{
		config: cfg,
	}
}

func (s *Service) Transform(ctx context.Context, data []byte) ([]byte, error) {
	var record map[string]interface{}
	if err := json.Unmarshal(data, &record); err != nil {
		return nil, err
	}

	for _, rule := range s.config.Rules {
		if err := s.applyRule(record, rule); err != nil {
			return nil, fmt.Errorf("failed to apply rule to field %s: %w", rule.Field, err)
		}
	}

	return json.Marshal(record)
}

func (s *Service) applyRule(record map[string]interface{}, rule config.TransformationRule) error {
	value, ok := record[rule.Field]
	if !ok {
		return nil
	}

	switch rule.Operation {
	case "convert":
		if rule.Value == "seconds" {
			if duration, ok := value.(float64); ok {
				record[rule.Field] = int64(duration)
			}
		}
	case "mask":
		if strVal, ok := value.(string); ok {
			record[rule.Field] = s.maskValue(strVal, rule.Value)
		}
	case "format_date":
		if timestamp, ok := value.(string); ok {
			t, err := time.Parse(time.RFC3339, timestamp)
			if err == nil {
				record[rule.Field] = t.Format(s.config.DateFormat)
			}
		}
	case "uppercase":
		if strVal, ok := value.(string); ok {
			record[rule.Field] = strings.ToUpper(strVal)
		}
	case "lowercase":
		if strVal, ok := value.(string); ok {
			record[rule.Field] = strings.ToLower(strVal)
		}
	}

	return nil
}

func (s *Service) maskValue(value, maskType string) string {
	switch maskType {
	case "last4":
		if len(value) > 4 {
			return strings.Repeat("*", len(value)-4) + value[len(value)-4:]
		}
		return value
	case "full":
		return strings.Repeat("*", len(value))
	default:
		return value
	}
}

func (s *Service) convertToSeconds(value interface{}) (int64, error) {
	switch v := value.(type) {
	case float64:
		return int64(v), nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return 0, fmt.Errorf("unsupported type for duration conversion")
	}
}
