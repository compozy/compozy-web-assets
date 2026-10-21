// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "5bd7ef49ca1f4639be1341edbcbf2a5589e23f2b9d007c725ed2d2fe4ef6157d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "491cb19a36acb593b7af70e1ac7ee39cf451b925"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
