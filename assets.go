// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "810ec19ec5c6ea7906a01f97d499c42928392c22f67d997cfdc1dc771f4bd9f1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "50db2cd6d59fa40e93199d3a8dada402086be398"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
