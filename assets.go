// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e0ddb4522042b4365ccedbf6a41707159bac66793c05ff8aeeb7ad685346665a"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "9f36865b26fb6f19e0a120a73ef235669f62b437"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
