package resources

import "github.com/samber/lo"

func FromDBResource(dbResource *DBResource) *Resource {
	return &Resource{
		ID:         dbResource.ID,
		Name:       dbResource.Name,
		Type:       dbResource.Type,
		Region:     dbResource.Region,
		CustomerID: dbResource.CustomerID,
		CreatedAt:  dbResource.CreatedAt,
		UpdatedAt:  dbResource.UpdatedAt,
	}
}

func FromDBResourceList(dbResources []*DBResource) []*Resource {
	return lo.Map(dbResources, func(resource *DBResource, _ int) *Resource {
		return FromDBResource(resource)
	})
}
