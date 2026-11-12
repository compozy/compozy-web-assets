// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "75d954711e92886712d0021821b29eb68a4e72a8bed46c8f7127db6b96787868"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a145e07cbe549dad54d237645c003c5a81c99fbe"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
