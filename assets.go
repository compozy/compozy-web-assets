// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "19adabe117c006bbed1a96f4b0402d88ebe747b37c343666757d426811313aff"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6b2f89bdf837b8c30315550c02b4f907752a8ceb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
