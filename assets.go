// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "aa7964bec3a09720d89195e3e15ab64b3f479ab8182f806d489567e4b54442c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fc5b189c919b64e0f039c291206afe533249a1f2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
