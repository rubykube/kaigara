package config

import "testing"

func TestSetAndGet(t *testing.T) {
	setDefaults()

	def := true
	def = Get("home") == "/opt/provision"
	def = Get("path") == "/opt/provision/operations"
	def = Get("tmpl") == "/opt/provision/resources"
	def = Get("env") == "development"

	if !def {
		t.Fail()
	}
}
