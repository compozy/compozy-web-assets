// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cf4c61cd0b202a3c965e6d1de641baba14098c0dd7c39f45239c199f2ad33d2f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0b740cff417767c7296c547b1bb8d93810b4ed14"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
