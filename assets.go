// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c7a3f4ad146b16943053cd3547d15473ddc349b8d2e9edc484901907532ab0e4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5f36f0a54ce03c6d3962b6c6d0e068a4f9e35da6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
