package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	
	// plug in plugins here, for example:
	// _ "import/path/here"

	// http.supervisor - https://github.com/lucaslorentz/caddy-supervisor/blob/master/README.md
	//_ "github.com/lucaslorentz/caddy-supervisor/httpplugin"
  	//_ "github.com/lucaslorentz/caddy-supervisor/servertype"

	// dns.gandiv5
	_ "github.com/caddyserver/dnsproviders/gandiv5"
)

func main() {
	// optional: disable telemetry
	// caddymain.EnableTelemetry = false
	caddymain.Run()
}
