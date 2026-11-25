// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "19adabe117c006bbed1a96f4b0402d88ebe747b37c343666757d426811313aff"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "043efc46a4f6a0cfb690b81de6955e8a818c57a1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
