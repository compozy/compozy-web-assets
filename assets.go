// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "aa7964bec3a09720d89195e3e15ab64b3f479ab8182f806d489567e4b54442c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "25dd7b8c98859df21fa3e51539f0b0c8eff6ec07"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
