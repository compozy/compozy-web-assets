// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "7b4d907c7877c3aa20d07352ae8d0c71e4f05d70a9171be5ecb00ae531c7c61e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "d6084fdc929e8e175335ed198472a1f501e7a642"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
