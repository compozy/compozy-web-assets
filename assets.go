// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d94f126eb9503dd8c04fbeddc1d4598d41577f9fdcf849b957b3f7fe89798746"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fc56d6300be0aae495e4a9d6ddf2a7276358a07e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
