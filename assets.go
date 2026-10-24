// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "d45729d2675ed5fe2a3c8391d682aabef4256245626b6b7ba7b52c5a56f61d1d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ad24d66d2e37c2fd80f2429a368e1a5967020934"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
