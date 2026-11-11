// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1607d272771346637e46f5c92f12da59d07b09288189708669bcd72ea47e464e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "0934534cd4cd4aa37f958c5c9aa37016fb0d41ec"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
