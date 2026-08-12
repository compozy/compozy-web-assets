// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "2e188af04e69ed68d11a8aab9252f655b0a5ed6a5ae63f9bee9ebcc43156d297"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "fee93b73b740650d64fb47a749b6248889baa613"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
