// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bfd597bf37f52884ba9aa9131170e7817a97d3615fc9003a76c92c53808799f8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "e4361378bdaa35212e244afb4418f06a43f54967"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
