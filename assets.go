// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5a17a64864dc2a1ba46d066028e4b0bf4bdb0c0ffb210b38105f69faaab445a4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "aca71f7bb3aa0a43d4b2b5b846b4f330b7c8d360"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
