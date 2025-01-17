package resources

import (
	_ "embed"
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"
)

//go:embed examples/command.yaml
var yamlCommand []byte

func TestDeserializeCorrectProcessYaml(t *testing.T) {
	// Given
	expected := Command{
		Basic: Basic{
			Kind: "Command",
		},
		Metadata: Metadata{
			Name: "my_command",
		},
		Spec: CommandSpec{
			Run: []string{`echo "Hello, World!"`},
		},
	}

	fmt.Println("File content: ", string(yamlCommand))

	// When
	command := Command{}
	err := yaml.Unmarshal(yamlCommand, &command)

	// Then
	if err != nil {
		t.Errorf("Failed to deserialize yaml: %v", err)
	}
	if command.Metadata.Name != expected.Metadata.Name {
		t.Errorf("Deserialized yaml is not as expected. Actual %v Expected %v", command, expected)
	}

	if command.Spec.Run[0] != expected.Spec.Run[0] {
		t.Errorf("Deserialized yaml is not as expected. Actual %v Expected %v", command, expected)
	}

}
