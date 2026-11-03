// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e155aa8249c2e401f6d327fb57f870f02b7159ed0b743c0d2ad9eeb1c4412d9d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b57c83f051e7a78e5a9135cbf304acde876bfdd1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
