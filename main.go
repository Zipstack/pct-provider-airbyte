package main

import (
	"github.com/zipstack/pct-plugin-framework/schema"
	"github.com/zipstack/pct-plugin-framework/server"

	"github.com/zipstack/pct-provider-airbyte/plugin"
)

func main() {
	server.Serve(plugin.NewProvider, []func() schema.ResourceService{
		plugin.NewSourceFakerResource,
		plugin.NewDestinationPostgresResource,
		plugin.NewConnectionResource,
	})
}
