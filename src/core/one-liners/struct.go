package one_liners

type Metadata struct {
	Verified bool `yaml:"verified"`
}

type Variable struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default,omitempty"`
}

type OneLiner struct {
	ID        string              `yaml:"id"`
	Info      Info                `yaml:"info"`
	Variables map[string]Variable `yaml:"variables"`
	OneLiner  OneLinerCmd         `yaml:"one-liner"`
}

type Info struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Category    string   `yaml:"default-category"`
	Author      string   `yaml:"author"`
	Complexity  string   `yaml:"complexity"`
	Metadata    Metadata `yaml:"metadata"`
	Tags        []string `yaml:"tags"`
}

type OneLinerCmd struct {
	Cmd string `yaml:"cmd"`
}
