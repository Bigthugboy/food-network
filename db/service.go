package db

type mainstore interface {
	Welcome()
	RegisterHandler()
	GetAllFoodNearbyHandler()
	SelectRestaurantHandler()
	GetMenuList()
	RateResurtant()
	SaveFavouriteResturant()
}
