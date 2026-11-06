// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9e71a6277849a2b986bd0c62f100c48d815e1405a71983b3c35832b7549c473d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f8ec00ccbf8568560a8bc1acf34ce879368669db"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
