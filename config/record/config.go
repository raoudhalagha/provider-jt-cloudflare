package record

import "github.com/crossplane/terrajet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("cloudflare_record", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "github" 
        r.ShortGroup = "record"

        // Identifier for this resource is assigned by the provider. In other
        // words it is not simply the name of the resource.
        //r.ExternalName = config.IdentifierFromProvider

        // This resource need the repository in which branch would be created
        // as an input. And by defining it as a reference to Repository
        // object, we can build cross resource referencing. See 
        // repositoryRef in the example in the Testing section below.
        r.References["zone"] = config.Reference{
            Type: "github.com/crossplane-contrib/provider-jt-cloudflare/apis/zone/v1alpha1.Zone",
        }
    })
}