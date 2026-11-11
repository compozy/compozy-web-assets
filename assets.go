// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "48b264b655a4dccfed6739e1d1ca9e85bd3563a7778ce5a4624eda6d70398502"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3bb17c1788b4ac945a144d4419621846537a6a55"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
