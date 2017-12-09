package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/muziyoshiz/ansible2tab/parser"
	"strings"
)

type Formatter interface {
	GetHeader() string
	Format(result parser.Result) string
	GetFooter() string
}

type TsvFormatter struct {
	Formatter
}

func (self *TsvFormatter) GetHeader() string {
	return ""
}

func (self *TsvFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, " ")
	return fmt.Sprintf("%s\t%s\n", result.Host, values)
}

func (self *TsvFormatter) GetFooter() string {
	return ""
}

type JsonFormatter struct {
	Formatter
	trailingLine bool
}

func (self *JsonFormatter) GetHeader() string {
	return "{"
}

func (self *JsonFormatter) Format(result parser.Result) string {
	jsHost, _ := json.Marshal(result.Host)
	jsValues, _ := json.Marshal(strings.Join(result.Values, "\n"))
	if self.trailingLine {
		return fmt.Sprintf(",%s:%s", jsHost, jsValues)
	} else {
		self.trailingLine = true
		return fmt.Sprintf("%s:%s", jsHost, jsValues)
	}
}

func (self *JsonFormatter) GetFooter() string {
	return "}\n"
}

type MarkdownFormatter struct {
	Formatter
}

func (self *MarkdownFormatter) GetHeader() string {
	return `|Host|Value|
|---|---|
`
}

func (self *MarkdownFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, " ")
	return fmt.Sprintf("|%s|%s|\n", result.Host, values)
}

func (self *MarkdownFormatter) GetFooter() string {
	return ""
}

type MarkdownCodeFormatter struct {
	Formatter
	trailingLine bool
}

func (self *MarkdownCodeFormatter) GetHeader() string {
	return ""
}

func (self *MarkdownCodeFormatter) Format(result parser.Result) string {
	values := strings.Join(result.Values, "\n")
	if self.trailingLine {
		// We can not escape backquote inside backquotes
		return fmt.Sprintf("\n## %s\n\n```\n%s\n```\n", result.Host, values)
	} else {
		self.trailingLine = true
		return fmt.Sprintf("## %s\n\n```\n%s\n```\n", result.Host, values)
	}
}

func (self *MarkdownCodeFormatter) GetFooter() string {
	return ""
}
