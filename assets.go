// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e8f60ed876f6e8737131800eef51866fa6619707e51eb1f9f9db63496654437e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "19b8ae06c03954205295478f982c32f3658ca94e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
