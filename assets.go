// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7053a418262456066c28ffe380ebe523b48212c225d281dd80bdbe5701cd2190"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "ab03af1e55b402bda1d1637868fb24872b03c013"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
