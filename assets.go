// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "1607d272771346637e46f5c92f12da59d07b09288189708669bcd72ea47e464e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "56ce976f7be6d263213b4d45c2d793f1266bf914"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
