package resources

type Command struct {
	Basic    `json:",inline" yaml:",inline"`
	Metadata Metadata    `json:"metadata" yaml:"metadata"`
	Spec     CommandSpec `json:"spec" yaml:"spec"`
}

type CommandSpec struct {
	Run []string `json:"run" yaml:"run"`
}
