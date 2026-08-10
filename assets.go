// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "3cd9d5380fb2bed75993ac513da02501f1eb8063b5601420a498a76a2cda8d98"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fb21a4994fa1dc33ed20c0ac14945b143079e27c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
