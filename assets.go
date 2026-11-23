// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "9bb47bcb52325cd796d19c0b6b11fd6889f1210d756b4ea3ae3fb3be14461567"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3afb5f76858b63af12bfe05b8150f8ce8c57aca6"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
