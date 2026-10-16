// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "c73542255478c4bbe6e8fe178dba15682fa3ad762cd83dd24130caf4c7a4633b"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "746e46d54f46484d9b80b1201bf7ffc48cd132b3"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
