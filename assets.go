// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5a8803b73b46294003781e4c916be0e71418e73822762b9cb477425c1713246e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a0b1a0f59ff62ea7b58cabed8d413b9bf6f2aef9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
