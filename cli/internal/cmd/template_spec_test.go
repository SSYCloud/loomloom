package cmd

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadTemplateSpecFile_ValidSpec(t *testing.T) {
	path := filepath.Join(t.TempDir(), "spec.json")
	content := `{
  "Meta": {"Name": "Spec Test", "Description": "desc"},
  "Steps": [{"StepID": "stp_text", "DisplayName": "Text", "ExecutionUnit": "text-generate"}],
  "InputSchema": {"Fields": [{"Key": "prompt", "Label": "Prompt", "ValueType": "string"}]},
  "FieldBindings": [{"FieldKey": "prompt", "StepID": "stp_text", "ParamKey": "prompt", "BindMode": "shared"}]
}`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write spec: %v", err)
	}

	spec, raw, err := loadTemplateSpecFile(path)
	if err != nil {
		t.Fatalf("loadTemplateSpecFile() error = %v", err)
	}
	if spec.Meta.Name != "Spec Test" {
		t.Fatalf("Meta.Name = %q, want Spec Test", spec.Meta.Name)
	}
	if len(raw) == 0 || raw[0] != '{' {
		t.Fatalf("expected compact JSON bytes, got %q", string(raw))
	}
}

func TestLoadTemplateSpecFile_MissingName(t *testing.T) {
	path := filepath.Join(t.TempDir(), "spec.json")
	content := `{
  "Meta": {},
  "Steps": [{"StepID": "stp_text"}],
  "InputSchema": {"Fields": []},
  "FieldBindings": []
}`
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write spec: %v", err)
	}

	if _, _, err := loadTemplateSpecFile(path); err == nil {
		t.Fatal("loadTemplateSpecFile() error = nil, want missing name error")
	}
}
