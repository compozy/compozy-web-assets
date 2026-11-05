// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "fcd1089256a4f0a7c529679ca0a57efcd8fa6d5092b78ef7e3dfafd651a3b5dc"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "839f8a621bc09ad146234dc0ed790028439fb276"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
