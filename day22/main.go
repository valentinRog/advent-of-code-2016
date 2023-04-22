package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)
	raw := strings.TrimSpace(string(bytes))
	raw = strings.ReplaceAll(raw, "\r", "")
	Part1(raw)
	// Part2(raw)
}
