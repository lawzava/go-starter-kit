package env_test

import (
	"goapp/pkg/env"
	"os"
	"testing"
)

func TestString(t *testing.T) {
	tt := []struct {
		in  env.EnvironmentVariable
		out string
	}{
		{"TEST", "TEST"},
	}

	for _, tc := range tt {
		if tc.in.String() != tc.out {
			t.Fatal(tc)
		}
	}
}

func TestGetOr(t *testing.T) {
	tt := []struct {
		in  env.EnvironmentVariable
		out string
	}{
		{"TEST", "VAR"},
	}

	for _, tc := range tt {
		err := os.Setenv(tc.in.String(), tc.out)
		if err != nil {
			t.Fatal(err)
		}

		if tc.in.GetOr("") != tc.out {
			t.Fatal(tc)
		}
	}
}
