// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e2b1acfe6dc9961f31aca460fd9832d9dd0ebd78231956e92813e1923f837a21"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a626b69948f6d2a9204efd46755860b39b780990"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
