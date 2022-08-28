package main

import (
	"bytes"
	"fmt"
	"io"
)

func DBEscapeString(txt string) string {
	var (
		esc string
		buf bytes.Buffer
	)
	last := 0
	for ii, bb := range txt {
		switch bb {
		case 0:
			esc = `\0`
		case '\n':
			esc = `\n`
		case '\r':
			esc = `\r`
		case '\\':
			esc = `\\`
		case '\'':
			esc = `\'`
		case '"':
			esc = `\"`
		case '\032':
			esc = `\Z`
		default:
			continue
		}
		io.WriteString(&buf, txt[last:ii])
		io.WriteString(&buf, esc)
		last = ii + 1
	}
	io.WriteString(&buf, txt[last:])
	return buf.String()
}

func DBSprintf(format string, a ...interface{}) string {

	// Reemplazamos \n por spacio => sólo en la query, no en los argumentos
	ReplaceAll(format, "\n", " ")

	for index, value := range a {
		// Vemos si es un string y lo escapamos
		switch value.(type) {
		case []byte:
			valueStr := string(value.([]byte))
			a[index] = DBEscapeString(valueStr)
		case []rune:
			valueStr := string(value.([]rune))
			a[index] = DBEscapeString(valueStr)
		case string:
			valueStr := value.(string)
			a[index] = DBEscapeString(valueStr)
		}
	}
	return fmt.Sprintf(format, a...)
}
func DBSprintfAlreadyEscaped(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
func DBUuid() string {
	return UuidSorted()
}