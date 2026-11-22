// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "bedab7f2d92f001c0ff1d0dff13fb023ffac4cd6030e1bb44261b6eeef372f8c"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2a84efbe25281642daed544c30c2673cfbfc63f1"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
