package xsaml_admin

import (
	admin_plugin "github.com/ecletus-pkg/admin"
	"github.com/ecletus-pkg/siteconf"
	"github.com/ecletus/admin"
	"github.com/ecletus/db"
	"github.com/ecletus/plug"
	ect_samlidp "github.com/moisespsena/go-ecletus-samlidp"
)

type Plugin struct {
	plug.EventDispatcher
	db.DBNames
	admin_plugin.AdminNames

	SamlIdpKey string
}

func (this *Plugin) OnRegister(options *plug.Options) {
	admin_plugin.Events(this).InitResources(func(e *admin_plugin.AdminEvent) {
		e.Admin.OnResourceValueAdded(&siteconf.SiteConfigMain{}, func(e *admin.ResourceEvent) {
			idp := options.GetInterface(this.SamlIdpKey).(*ect_samlidp.SamlIDP)
			idp.ConfigureResource(e.Resource)
		})
	})
}
