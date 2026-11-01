// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f5a01242e0be7db936a6b56cf4960bee2b2930e2d2b72f5c0fee37f04a316009"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4026c7091c9c593f78376c110f95b2fe49b31413"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
