// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "748de31e6ab104c1d4c3a5e2060d4ae2e9e45a86082fd1de54868a0798e817dc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "927b9b70398570ce5f7e1265bf4297611abffdaf"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
