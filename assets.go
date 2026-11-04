// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c2c2c0cfb7f1813f2163e58b0a549bc7725ea5d53d1a25b38ac96ce7ed2c626d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "dc37fd7f0fbaa0c4cd1c5329582d94be92187f88"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
