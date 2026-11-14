// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "108181a6002daf130a04e529b2432a89b368da1852e30e1274d17a9ff4ce8f50"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "cdcefc3ee082b603aed8ca1f6414264dd572bfdb"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
