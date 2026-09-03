// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cabc14ce963b02b16878edb82e14e1dbe37707f36f3628a2dbe5a38e4c76aef0"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "b7667d8f7c1c70eb6c943b757fd7771f746683c5"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
