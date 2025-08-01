package views

import "embed"

//go:embed assets
var AssetFiles embed.FS

//go:embed templates
var TemplateFiles embed.FS
