// The following directive is necessary to make the package coherent:
// +build ignore

// This program generates code automatically and it should be run before build.
// It can be invoked by running: go generate

package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// max payload size

const payloadLength = 1e6

var messageConstTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
package message

const payload = "{{ .Payload }}"
`))

func main() {
	f, err := os.Create("message-no-headers-const-gen.go")
	if err != nil {
		log.Fatalf("Err: %+v", err)
	}
	defer f.Close()

	var b strings.Builder
	for i := 0; i < payloadLength; i++ {
		fmt.Fprintf(&b, "%d", i%10)
	}
	messageConstTemplate.Execute(f, struct {
		Payload string
	}{
		Payload: b.String(),
	})
}
