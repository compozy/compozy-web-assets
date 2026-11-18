// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c3ffcd4ce468c0252655bae9de29fb1e0da1bf2c30f4a7167a0ba663b0ed1163"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a72d4f8c75527bd0d7903c999faad55845d4a09f"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
