// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8400d86224fedc1641a76a05f0b8acbf0801b943d8b2c7bc78c696e0098f8385"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "43c0ab170e90a483e2a3e559bccb5233aea33971"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
