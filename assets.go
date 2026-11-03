// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e5ac7113c9e933bfdc58dab5fb8e6e427f601ad267d0337b49d8e25b7211966c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "07484f7cc2b1d7485d41a1fc3f725db969545713"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
