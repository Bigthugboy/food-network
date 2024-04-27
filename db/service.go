package db

type mainstore interface {
	Welcome()
	RegisterHandler()
	GetFoodNearbyHandler()
	SelectRestaurantHandler()
	GetMenuList()
	RateResurtant()
	SaveFavouriteResturant()
}
