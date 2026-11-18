// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c3ffcd4ce468c0252655bae9de29fb1e0da1bf2c30f4a7167a0ba663b0ed1163"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "59c3a7c1e5ab80246d415314f790ba5968a5e1af"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
