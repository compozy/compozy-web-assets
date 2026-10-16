// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "44905d3c64c8ee02b1c9b8c4a284b1e128c3a8f0a4084ebbd338a92ba6c27d2d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "bdc535a0ab505ca131673119b80f5cb5d76e8fb8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
