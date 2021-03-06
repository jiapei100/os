package templates

import (
	"encoding/json"
	"strings"
	"text/template"
)

// basicFunctions are the set of initial
// functions provided to every template.
var basicFunctions = template.FuncMap{
	"json": func(v interface{}) string {
		a, _ := json.Marshal(v)
		return string(a)
	},
	"split": strings.Split,
	"join":  strings.Join,
	"title": strings.Title,
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
}

// Parse creates a new annonymous template with the basic functions
// and parses the given format.
func Parse(format string) (*template.Template, error) {
	return NewParse("", format)
}

// NewParse creates a new tagged template with the basic functions
// and parses the given format.
func NewParse(tag, format string) (*template.Template, error) {
	return template.New(tag).Funcs(basicFunctions).Parse(format)
}
