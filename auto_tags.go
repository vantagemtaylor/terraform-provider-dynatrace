/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package main

import (
	"context"
	"log"
	"reflect"

	api "github.com/dtcookie/dynatrace/api/config"
	"github.com/dtcookie/dynatrace/api/config/autotags"
	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/terraform"
	"github.com/dtcookie/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type autoTags struct {
}

// Resource produces terraform resource definition for Management Zones
func (aps *autoTags) Resource() *schema.Resource {
	resource := terraform.ResourceFor(new(autotags.AutoTag))
	resource.CreateContext = logging.Enable(aps.Create)
	resource.UpdateContext = logging.Enable(aps.Update)
	resource.ReadContext = logging.Enable(aps.Read)
	resource.DeleteContext = logging.Enable(aps.Delete)

	return resource
}

// Create expects the configuration of a Management Zone within the given ResourceData
// and send them to the Dynatrace Server in order to create that resource
func (aps *autoTags) Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if config.HTTPVerbose {
		log.Println("AutoTags.Create(..)")
	}
	var err error
	var resolver terraform.Resolver
	if resolver, err = terraform.NewResolver(d); err != nil {
		return diag.FromErr(err)
	}
	var untypedConfig interface{}
	if untypedConfig, err = resolver.Resolve(reflect.TypeOf(autotags.AutoTag{})); err != nil {
		return diag.FromErr(err)
	}

	typedConfig := untypedConfig.(autotags.AutoTag)
	typedConfig.ID = nil
	conf := m.(*config.ProviderConfiguration)
	apiService := autotags.NewService(conf.DTenvURL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	var objStub *api.EntityShortRepresentation
	if objStub, err = apiService.Create(&typedConfig); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(objStub.ID)

	return aps.Read(ctx, d, m)
}

// Update expects the configuration of a Management Zone within the given ResourceData
// and send them to the Dynatrace Server in order to update that resource
func (aps *autoTags) Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if config.HTTPVerbose {
		log.Println("AutoTags.Update(..)")
	}
	var err error
	var resolver terraform.Resolver
	if resolver, err = terraform.NewResolver(d); err != nil {
		return diag.FromErr(err)
	}
	var untypedConfig interface{}
	if untypedConfig, err = resolver.Resolve(reflect.TypeOf(autotags.AutoTag{})); err != nil {
		return diag.FromErr(err)
	}

	typedConfig := untypedConfig.(autotags.AutoTag)
	typedConfig.ID = opt.NewString(d.Id())
	conf := m.(*config.ProviderConfiguration)
	apiService := autotags.NewService(conf.DTenvURL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	if err = apiService.Update(&typedConfig); err != nil {
		return diag.FromErr(err)
	}
	return aps.Read(ctx, d, m)
}

// Read queries the Dynatrace Server for the configuration of a Management Zone
// identified by the ID within the given ResourceData
func (aps *autoTags) Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if config.HTTPVerbose {
		log.Println("AutoTags.Read(..)")
	}
	var err error
	conf := m.(*config.ProviderConfiguration)
	apiService := autotags.NewService(conf.DTenvURL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	typedConfig := new(autotags.AutoTag)
	if typedConfig, err = apiService.Get(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	if err = terraform.ToTerraform(typedConfig, d); err != nil {
		return diag.FromErr(err)
	}

	return diag.Diagnostics{}
}

// Delete a Management Zone on the Dynatrace Server
func (aps *autoTags) Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if config.HTTPVerbose {
		log.Println("AutoTags.Delete(..)")
	}
	var err error
	conf := m.(*config.ProviderConfiguration)
	apiService := autotags.NewService(conf.DTenvURL, conf.APIToken)
	rest.Verbose = config.HTTPVerbose
	if err = apiService.Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}
