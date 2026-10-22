// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5bd7ef49ca1f4639be1341edbcbf2a5589e23f2b9d007c725ed2d2fe4ef6157d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f137c1068fc93fa0f6603c801439c3f3df8db6b1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
