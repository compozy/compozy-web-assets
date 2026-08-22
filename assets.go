// Package webassets embeds the production CompozyOS web UI bundle.
package webassets

import "embed"

// DistDir is the root directory embedded in DistFS.
const DistDir = "dist"

const (
	BuildDigest = "30cb32b750f4af6cf5b57e449094b9e88b5ce0c7435421025e49604d0d95edca"
	SourceRepository = "github.com/compozy/compozy"
	SourceCommit = "96a771efa8944d4b13c4427919ae3aeb9bd70857"
)

// DistFS embeds the generated production CompozyOS web UI bundle.
//
//go:embed all:dist
var DistFS embed.FS
