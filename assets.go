// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "20081ee16d53e5c09fc59f979741b2f41de0b2f276f0b1cfa42b1c4f3fb66c42"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "beaba45a0754112c46fce34e60d53a9f74b65aaa"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
