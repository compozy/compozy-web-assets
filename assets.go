// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7e1c8bc760213e890896889ef8e5f38bbe1289b56369253fa90a0ce35001d92c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4960784df84c2f101eb157aee82e555e1e7e8722"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
