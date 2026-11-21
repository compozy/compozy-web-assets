// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "cbd77f5f57c88d5dd1dc5d580da647addbfeae96322b860b8e1192f4f8df871d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "3760651c89c8c6bd734df5a5bfacaf876f6685ad"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
