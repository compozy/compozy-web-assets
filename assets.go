// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d6b44d8f40f484ee834edfd4dfef63270189691b9e75afc600056a73925fd442"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6fc6d4a1d25af9751b79bfa039030bfd7c88a8e9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
