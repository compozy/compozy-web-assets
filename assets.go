// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e749b3654d6c498f9af08be7f83e7d245d7622d5b8b6cd89fe7813760de9913d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "187c82a3e1ad93bea33b14e1aa5581ae8431c468"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
