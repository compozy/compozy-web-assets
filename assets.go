// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "105521818d4f2dc7f50d8bd51077ef58e2532781a34f54b24efdf843d6c371ba"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5d88f412f4622b35ce3f9ff19b1c1b3191cb78c2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
