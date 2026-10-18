// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "79805125f902e5655304675c1d30b96ea4085380e9f28f34fb18d941d6d036f4"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8e9ed1c10ed2887cd09bd5d54ea70a4d1d7b8157"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
