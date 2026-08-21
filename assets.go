// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "243e298cda58f4238d4d7ef211a9232dad3ae6ef04b3a08fa045dd312d9faace"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "609f2d2249f9343836204ea8b86104ccf4eac3d3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
