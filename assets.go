// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "73dedd7755a882936410a86d1322fcd5c4ab5dddda3e0d24eda495a368e5be9e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5d71954480cd56aaf766f6525a1f8c8c88b43a6a"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
