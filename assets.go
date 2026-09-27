// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "068796d9834ce4a82400bf1a134b0233e04ee8d0603e8e7f794ceb597a3513e4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5ecb6da89774e346929e5be7a885440785e8550b"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
