// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "eb397a51e6f5be88186b64aa97bfc891733a932ef74a74550038e893ae5ab253"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fb78851b276875fe7c9317388df4254c20eedd5d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
