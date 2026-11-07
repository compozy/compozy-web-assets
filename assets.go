// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cac068e53d465a1cd331e3947db39e62c7e4d63175e68222665131f35d2173a4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f0734247100d037eb9c8e4b26c9db2de9246b32d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
