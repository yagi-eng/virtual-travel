package infrastructure

import (
	"github.com/yagi-eng/place-search/interfaces/controllers"

	"github.com/labstack/echo"
)

// Router ルーティング
type Router struct {
	e  *echo.Echo
	lc *controllers.LinebotController
	ac *controllers.APIController
}

// NewRouter コンストラクタ
func NewRouter(e *echo.Echo, lc *controllers.LinebotController, ac *controllers.APIController) *Router {
	return &Router{e: e, lc: lc, ac: ac}
}

// Init ルーティング設定
func (r *Router) Init() {
	r.e.POST("/linebot/callback", r.lc.CatchEvents())

	api := r.e.Group("/api")
	{
		api.GET("/search", r.ac.Search())

		favorite := api.Group("/favorite")
		{
			favorite.POST("/get", r.ac.GetFavorites())
			favorite.POST("/add", r.ac.AddFavorites())
			favorite.POST("/remove", r.ac.RemoveFavorites())
		}
	}
}
