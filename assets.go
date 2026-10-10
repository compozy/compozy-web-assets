// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5d5921a58b4cbcd84dfe8ce78a25a2cd890b5790c409aa8ecdfba223e2db197c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "88534944a7b348cbde9b47c7b5dd4f23dc7fa5b5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
