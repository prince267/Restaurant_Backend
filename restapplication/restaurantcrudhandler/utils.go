package restaurantcrudhandler

import "MAD_Assignment/domain"

func transformobjListToResponse(resp []*domain.Restaurant) RestaurantGetListRespDTO {
	responseObj := RestaurantGetListRespDTO{}
	for _, obj := range resp {
		restaurantObj := RestaurantGetRespDTO{
			ID:           obj.DBID.String(),
			Name:         obj.Name,
			Address:      obj.Address,
			AddressLine2: obj.AddressLine2,
			URL:          obj.URL,
			Outcode:      obj.Outcode,
			Postcode:     obj.Postcode,
			Rating:       obj.Rating,
			Type_Of_Food: obj.Type_Of_Food,
		}
		responseObj.Restaurants = append(responseObj.Restaurants, restaurantObj)
	}
	responseObj.Count = len(responseObj.Restaurants)

	return responseObj
}
