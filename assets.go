// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "772a6b13772157328d15a5ba4fe1754e45b0e2c1d45f490aa28d43aab7374ef1"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b19490a327ed3c830c07fca1741bfffa5dff63eb"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
