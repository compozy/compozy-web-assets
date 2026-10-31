// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b5878d91c88c0d1725339091879de79dc92787d5caa41476da89f2e3ecb3892d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "46d3fda4f2d24fdf06110aed2300f73b5e8389cd"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
