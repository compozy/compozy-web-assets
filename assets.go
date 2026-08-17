// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "41463dda32ee83299449ee2c49c569ee678e654623b4af239a432034e96bb824"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "4681cc2a94e02006e30c629c9c8d50b39c1bc684"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
