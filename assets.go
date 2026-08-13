// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d5fa9bb9fc8d2030deb06e01ce98b9247b357c32460b9d67624f1f79e5b227cf"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c86ea430e11d5a1ca07e683a331fe25090c87de9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
