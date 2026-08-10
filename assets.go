// Package webassets embeds the production Compozy web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "e88211a8a1b9cc6bf08cd1accddef3fc9aee6f822ae742dfb6e16e5913895ffa"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "467a0f88cd1128aa258cc411219a19ad7cd660f5"
)

// DistFS embeds the generated production Compozy web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
