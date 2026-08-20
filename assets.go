// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ef280a292ed888495c6758ae93b2ff552e2f6fc5d3b36b1eefc5961ccd92dd18"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "636bc823ad8711d7964e7c1d59193930e9906454"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
