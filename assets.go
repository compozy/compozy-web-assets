// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1607d272771346637e46f5c92f12da59d07b09288189708669bcd72ea47e464e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "5e03e889407f08aa7c5d735e7e3aa9f602260f2e"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
