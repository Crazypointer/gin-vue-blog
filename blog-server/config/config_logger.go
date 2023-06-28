package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	ShowLine     bool   `yaml:"show_line"`
	Director     string `yaml:"director"`
	LogInConsole bool   `yaml:"log_in_console"`
}
