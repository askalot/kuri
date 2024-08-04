package config

import (
	"net/http"
)

const (
	staticAssetsDirectory = "/assets"
	StaticAssetsPath      = "/static/"

	JavscriptLink  = StaticAssetsPath + "scripts/"
	StylesheetLink = StaticAssetsPath + "styles/"
	ImageLink      = StaticAssetsPath + "images/"
)

func SetupStaticAssets() {
	fs := http.FileServer(http.Dir(staticAssetsDirectory))

	http.Handle(StaticAssetsPath, http.StripPrefix(StaticAssetsPath, fs))
}
