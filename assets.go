// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5ee8511b3ad4076db5e8d61394d955561bf2a84ee4986499b19569b29918a5c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "70a5bc8ffc70f6e84440d665851a4c18adcca121"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
