package env

import "os"

type EnvironmentVariable string

func (e EnvironmentVariable) String() string {
	return string(e)
}

func (e EnvironmentVariable) GetOr(fallback string) string {
	v := os.Getenv(e.String())

	if v != "" {
		return v
	}

	return fallback
}
