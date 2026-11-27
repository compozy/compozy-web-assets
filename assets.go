// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7547da3a95d82250e6bbec875556e6415f38e80eedea0f7efb332401e26b6db5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "afa96ff2194d173fd51c2ef1b8e01c486c500dc5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
