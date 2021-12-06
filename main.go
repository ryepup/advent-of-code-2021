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

	writeTemplate(
		day,
		fmt.Sprintf("day%v_test.go", day),
		"test",
		day)

	for _, part := range []int{1, 2} {
		writeTemplate(
			day,
			fmt.Sprintf("part%v.go", part),
			"part",
			struct{ Day, Part int }{day, part})
	}

	err = os.WriteFile("test.txt", []byte("TODO: copy in test input\n"), 0644)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("input.txt", []byte("TODO: copy in real input\n"), 0644)
	if err != nil {
		panic(err)
	}
}

func writeTemplate(day int, filename, template string, data interface{}) {
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
			n, err := Part1(path)
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
	`))
