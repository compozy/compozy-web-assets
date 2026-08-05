// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "11e6d606b74b42ce0aa4f3265a6716804b9c87cef40f0d41bc51faaa693a95ef"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7e630c6025b07b40e70035e16a564e0570711d83"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
