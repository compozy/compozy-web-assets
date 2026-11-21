// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "dc9f49cfa7bd6451c5874b3c32573a315894da4df1f68a1657455a56cc0b76f3"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "24bd5b54dcb96c88fda5a94833c55edcc61b83df"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
