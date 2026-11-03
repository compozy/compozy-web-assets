// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "417a49f60347872dad101f2df9b8e3db62b224aa4506532fbe4ffef9f28ab464"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "a8e8730f042c64e327fdb1dfa61dbc1f98cd38da"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
