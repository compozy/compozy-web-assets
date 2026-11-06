// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9e71a6277849a2b986bd0c62f100c48d815e1405a71983b3c35832b7549c473d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "8e7fa1af2bb997c0ec23abd9ad3aea97ba0ec171"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
