package main

import (
	"github.com/zipstack/pct-plugin-framework/schema"
	"github.com/zipstack/pct-plugin-framework/server"

	"github.com/zipstack/pct-provider-airbyte-local/plugin"
)

// Set while building the compiled binary.
var version string

func main() {
	server.Serve(version, plugin.NewProvider, []func() schema.ResourceService{
		plugin.NewConnectionResource,

		// plugin.NewSourceFakerResource,
		plugin.NewSourcePipedriveResource,
		plugin.NewSourceStripeResource,
		plugin.NewSourceAmplitudeResource,
		plugin.NewSourceShopifyResource,
		plugin.NewSourceFreshdeskResource,
		plugin.NewSourceZendeskSupportResource,
		plugin.NewSourceHubspotResource,

		// plugin.NewDestinationPostgresResource,
		plugin.NewDestinationLocalCSVResource,
	})
}
