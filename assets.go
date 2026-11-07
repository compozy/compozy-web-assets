// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "04813c0baf3fb309ff9a37cab9ac2e41d7e1c0735c9fa8e56e3ecf21d4dce1e8"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "59fd1fc7d582cc304e21085095a8619a90ce1e41"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
