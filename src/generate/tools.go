package generate

import (
	"os"
	"strings"
	"text/template"
	"unsafe"
)

type Result struct {
	output string
}

func (r *Result) Write(b []byte) (n int, err error) {
	r.output += string(b)
	return len(b), nil
}

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Render(name string, path string, data interface{}) (string, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	writer := &Result{}
	if err := t.Execute(writer, data); err != nil {
		return "", err
	}

	return writer.output, nil
}

func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func ConvertModel(model string) string {
	if strings.Contains(model, "-") { // 如果含有中杠，则认为该数据模型是由多个单词组成的
		var target string
		words := strings.Split(model, "-")

		for _, word := range words {
			target = target + strings.Title(word)
		}

		return target
	}

	return strings.Title(model)
}

func Convert2Camel(model string) string {
	if strings.Contains(model, "-") {
		words := strings.Split(model, "-")
		target := strings.ToLower(words[0])
		wordsLength := len(words)
		for i := 1; i < wordsLength; i++ {
			target = target + strings.Title(words[i])
		}

		return target
	}

	return strings.ToLower(model)
}

func GenerateSimplyName(model string) string {
	if strings.Contains(model, "-") {
		words := strings.Split(model, "-")

		target := strings.ToLower(string(words[0][0]))
		wordsLength := len(words)
		for i := 1; i < wordsLength; i++ {
			target = target + string(strings.ToLower(words[i])[0])
		}

		return target
	}

	return strings.ToLower(string(model[0]))
}
