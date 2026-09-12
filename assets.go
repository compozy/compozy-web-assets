// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8986f6d2582b8b4165eb55cb6ab429d2f975b4eb48a30c2e9c74fec0258bf219"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "f8d0946bc6e221a2d8c13a184d6b29c64b0a4e28"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
