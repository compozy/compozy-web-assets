// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e17b8e412d462ff9521976befe255af638265dbbddf4e92464d650fc9bcee51c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "13f4f3dbdfe86b1226ee621f63f98066199aae5e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
