// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "266cfb356eeff9f1440fdbbc8111313a91ab0b76a720d894f8f3327dd84340e3"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "6da3b271d0872555eed9298d2084b83057118949"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
