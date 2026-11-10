// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "105521818d4f2dc7f50d8bd51077ef58e2532781a34f54b24efdf843d6c371ba"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a7cdb94bf72b76177e1a1c7924bf108e191cb4a8"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
