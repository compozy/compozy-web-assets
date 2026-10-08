// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "b44d469c9dca456ca9b3c7556e93697971a009d9b3eaea5568fd508cbac5f5c6"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c15729cfb2838010ca80b6f62838fe9a33cb6881"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
