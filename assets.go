// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "00a36abc4f9c14ac4a4d9e3738ec9aaf8ec6b4a8fa0e63f323cbc3e5ca753092"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7ccc00670d1e278515049d6ff7ba8b3c00b7ee25"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
