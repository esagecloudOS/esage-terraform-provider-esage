package main

import (
	"fmt"
	"net/url"
	"path"

	"github.com/abiquo/ojal/abiquo"
	"github.com/abiquo/ojal/core"
	"github.com/hashicorp/terraform/helper/schema"
)

var vdcDataSchema = map[string]*schema.Schema{
	"name":      attribute(required, text),
	"tiers":     attribute(computed, text),
	"location":  attribute(computed, text),
	"network":   attribute(computed, text),
	"templates": attribute(computed, text),
	"device":    attribute(computed, text),
}

func vdcNetwork(r core.Resource) string {
	vdc := r.(*abiquo.VirtualDatacenter)
	network := vdc.Links.Find(func(l *core.Link) bool {
		return l.Rel == "defaultvlan"
	})
	return network.URL()
}

func virtualdatacenterFind(d *resourceData) (err error) {
	enterprise := abq.Enterprise()
	id := path.Base(enterprise.URL())
	vdcsLink := enterprise.Rel("cloud/virtualdatacenters")
	vdcs := vdcsLink.Collection(url.Values{
		"enterprise": {id},
		"has":        {d.string("name")},
	})
	vdc := vdcs.First()
	if vdc == nil {
		return fmt.Errorf("vdc %q was not found", d.string("name"))
	}

	d.SetId(vdc.URL())
	d.Set("device", vdc.Rel("device").Href)
	d.Set("tiers", vdc.Rel("tiers").Href)
	d.Set("network", vdcNetwork(vdc))
	d.Set("location", vdc.Rel("location").Href)
	d.Set("templates", vdc.Rel("templates").Href)
	return
}
