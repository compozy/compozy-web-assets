// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "94520f40d29097d92bf8ac84bf70eb2f0ba297dd0541245a3fe417fdc3dcfdb8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f4e2fdfa111b02bf6c1a88c70dfa69661fd24ca9"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
