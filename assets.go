// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2b4e3dd52769bed4f2de2818c940b1253015e0cbd6b0dab037bb0f231d1c72a5"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ac41319354916e139b211f2a0cccea44299dd252"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
