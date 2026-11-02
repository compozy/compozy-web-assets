// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "399a70fee3b81283ab069b9a8a3dacb9ac64124979fa7f7fe0248da65538166f"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "c7cb9d0c5312ede56f5dbe3a6a915e62db80863d"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
