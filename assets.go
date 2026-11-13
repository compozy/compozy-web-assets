// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7f4b56ecab6e4949c52c9df0ea36cb2df628e208954fd8e0122372b2effce943"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f7b5bf13bb4ff417f9f1e935103515e2c8d46219"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
