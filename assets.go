// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "db9acc5f9ec05ebd40d812bfe8788ed23957992b4ca91e53ffdfe1cae20273f5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "35d873f2ef4791e929f773c2e2d925e3ef1217c0"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
