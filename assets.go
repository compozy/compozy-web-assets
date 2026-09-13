// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "8986f6d2582b8b4165eb55cb6ab429d2f975b4eb48a30c2e9c74fec0258bf219"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "64b36b4cf674d5a71af69dbe2473ad7ee91217b2"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
