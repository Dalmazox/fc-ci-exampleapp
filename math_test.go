package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		Description    string
		X              int
		Y              int
		ExpectedResult int
	}{
		{
			Description:    "should sum two integer numbers",
			X:              10,
			Y:              20,
			ExpectedResult: 30,
		},
	}

	for _, scenario := range tests {
		t.Run(scenario.Description, func(t *testing.T) {
			t.Parallel()

			result := sum(scenario.X, scenario.Y)

			if result != scenario.ExpectedResult {
				t.Errorf("expected to be %v, but received %v", scenario.ExpectedResult, result)
			}
		})
	}
}
