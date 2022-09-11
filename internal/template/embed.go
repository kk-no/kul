package template

import "embed"

//go:embed go.mod.txt main.go.txt
var Templates embed.FS
