// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "96f6c9b09a0e4309f5225c34a0fc27cab986f7158b917689da7fbd9073f383bd"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7e95d949809717888b17746c0feefffe11f0ad71"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
