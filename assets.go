// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "a878b319a3a0f509ee961b5452d879aea906173fa43ccde3a07d4c77ae10181b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4eb4968fead6a2b72e626980924dd4d6291a37d9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
