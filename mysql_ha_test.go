package goup

import "testing"

func TestInstallMysqlHa(t *testing.T) {
	settings := MustLoadConfig("goup/testdata/config.toml")
	if err := InstallMysqlHa(settings); err != nil {
		t.Fatalf("%s", err)
	}
}
