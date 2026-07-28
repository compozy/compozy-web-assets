// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "ee4bb6e1da65a96c9b9de12db9bcd7c5c35ce7fa366c8f154c3a5abbd4bbfc4e"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "7b66bf60f17fdb25514a516f5f5ec7c1bb08f0c7"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
