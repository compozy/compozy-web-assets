// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "f9fc4ffa708eefda78e07f55d08001c78326e4132b8e4c9637bc7451b1a36632"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f514c79a3de16282939cfe4a08b16147e64bffc6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
