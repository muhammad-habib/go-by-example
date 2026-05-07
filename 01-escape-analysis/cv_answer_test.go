package main

import (
	"strings"
	"testing"
)

func TestProductionExperienceAnswer(t *testing.T) {
	answer := ProductionSystemsExperienceAnswer()

	requiredPhrases := []string{
		"8+ years",
		"30+ endpoints",
		"1M+ users",
		"Kubernetes",
		"AWS Lambda",
		"does not explicitly mention on-call ownership",
	}

	for _, phrase := range requiredPhrases {
		if !strings.Contains(answer, phrase) {
			t.Fatalf("answer missing expected phrase %q: %q", phrase, answer)
		}
	}
}
