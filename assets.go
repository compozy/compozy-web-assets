// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2fcfb4f8bab6ebf7c86006691ff1102c0d94d158a3a6c174b20511680cb50942"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "11549d10e2f1b6cc69b7c311c37dca741d30a9ff"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
