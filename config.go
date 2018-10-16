package main

type ReplaceConfig struct {

	// Field to do replacement on. The default is message.
	Field string `config:"field"`

	// Target field for the replaced value. The default is same as field.
	Target string `config:"target"`

	// Replace from.
	Find string `config:"find"`

	// Replace to.
	Replace string `config:"replace"`

	// Use regex. default false.
	Regex bool `config:"regex"`
}

var defaultReplaceConfig = ReplaceConfig{
	Field:  "message",
	Target:  "",
	Find:  "",
	Replace: "",
	Regex: false,
}
