package infra

import "testing"

func TestEnvSet(t *testing.T) {
	EnvSet()
	if AppEnv.Port == 0 {
		t.Fatalf("fail to set env port %#v", AppEnv.Port)
	}
}
