// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "77adbf29c0875e266a6fc2dc34141dd1d708e64308e3c832383994406de36f8d"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "2cf933f08774a7edec2c341dfc8b283e7656a736"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
