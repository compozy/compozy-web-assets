// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "19adabe117c006bbed1a96f4b0402d88ebe747b37c343666757d426811313aff"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ea54b6deecfad44e403928a86b11fced52220338"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
