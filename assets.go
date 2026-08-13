// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d5fa9bb9fc8d2030deb06e01ce98b9247b357c32460b9d67624f1f79e5b227cf"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "06dc5b50cc46d6a1c68bd9d407b17445bd23670c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
