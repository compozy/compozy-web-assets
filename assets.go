// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2b9823074700067d6b36b94c14aee6608b78721067a55c53b8940f5806358a46"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "eadbf8ff0836202d4433506c0127d3e2c027e33e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
