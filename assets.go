// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "4f0852f8f0069b088371fdedc34baee8f66dcdf261c61402bec3e2ef391c344e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4bc5806891976202a3277732a1501e0171a9e63c"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
