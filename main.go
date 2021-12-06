package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

func main() {
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("scaffolding day %v\n", day)

	dir := fmt.Sprintf("day%v", day)
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}
	os.Chdir(dir)

	writeTemplate("day_test.go", "test", day)
	writeTemplate("part1.go", "part", struct{ Day, Part int }{day, 1})
	writeTemplate("part2.go", "part", struct{ Day, Part int }{day, 2})
	writeTemplate("test.txt", "input", "test")
	writeTemplate("input.txt", "input", "real")
}

func writeTemplate(filename, template string, data interface{}) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		if os.IsExist(err) {
			fmt.Printf("Skipping existing file %v\n", filename)
			return
		}
		panic(err)
	}
	defer f.Close()
	tmpl.ExecuteTemplate(f, template, data)
}

var tmpl = template.Must(template.New("advent").Parse(`
{{define "testfn" -}}
func TestPart{{.}}(t *testing.T) {
	testCases := map[string]int{"test.txt": -1, "input.txt": -1}
	for path, expected := range testCases {
		t.Run(path, func(t *testing.T) {
			if expected == -1 {
				t.Skip("TODO: provide an expected value")
			}
			n, err := Part{{.}}(path)
			require.Nil(t, err)
			require.Equal(t, expected, n)
		})
	}
}
{{end}}
{{- define "test" -}}
package day{{.}}_test

import (
	. "ryepup/advent2021/day{{.}}"
	"testing"

	"github.com/stretchr/testify/require"
)

{{ template "testfn" 1 }}
{{ template "testfn" 2 }}
{{- end}}
{{define "part" -}}
package day{{.Day}}

import "fmt"

/*
TODO: copy in problem definition
*/
func Part{{.Part}}(path string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
{{end}}
{{define "input" -}}
// TODO: copy in {{.}} input
{{end}}
	`))
