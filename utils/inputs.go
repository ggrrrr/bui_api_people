package utils

import (
	"log"
	"strings"
)

var A = []string{
	"o",
	"O",
	" ",
	`"`,
	"О",
	"o",
	"о",
	"O",
}

func NormalizeNumber(str string) string {
	var out1 = str
	// strings.NewReplacer()
	for v := range A {
		out1 = strings.Replace(out1, A[v], "0", -1)

	}
	// out1 = strings.Replace(out1, "o", "0", -1)
	log.Printf("out1: %s", out1)
	return out1
}
