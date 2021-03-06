package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Environment struct {
	Port       string
	HeaderName string
	KubeConfig string
	Timeout    time.Duration
}

func Get() (Environment, error) {
	var missing []string
	var env Environment
	var timeout string

	for k, v := range map[string]*string{
		"KUBECONFIG": &env.KubeConfig,
		"TIMEOUT_MS": &timeout,
	} {
		*v = os.Getenv(k)
	}

	for k, v := range map[string]*string{
		"PORT":        &env.Port,
		"HEADER_NAME": &env.HeaderName,
	} {
		*v = os.Getenv(k)

		if *v == "" {
			missing = append(missing, k)
		}
	}

	if len(missing) > 0 {
		return env, fmt.Errorf("missing env(s): %s", strings.Join(missing, ", "))
	}

	if timeout != "" {
		t, err := strconv.Atoi(timeout)
		if err != nil {
			return env, fmt.Errorf("timeout is invalid: %w", err)
		}
		env.Timeout = time.Duration(t)
	} else {
		env.Timeout = 30000 * time.Millisecond
	}

	return env, nil
}
