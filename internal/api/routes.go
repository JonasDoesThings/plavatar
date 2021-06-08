package api

func (server *Server) routes() {
	server.echoRouter.GET("/pixel/:size/:name", server.HandleGetPixelAvatar())
	server.echoRouter.GET("/pixel/:size", server.HandleGetPixelAvatar())
	server.echoRouter.GET("/solid/:size/:name", server.HandleGetSolidAvatar())
	server.echoRouter.GET("/solid/:size", server.HandleGetSolidAvatar())
	server.echoRouter.GET("/gradient/:size/:name", server.HandleGetGradientAvatar())
	server.echoRouter.GET("/gradient/:size", server.HandleGetGradientAvatar())
	server.echoRouter.GET("/marble/:size/:name", server.HandleGetMarbleAvatar())
	server.echoRouter.GET("/marble/:size", server.HandleGetMarbleAvatar())
	server.echoRouter.GET("/laughing/:size/:name", server.HandleGetLaughingAvatar())
	server.echoRouter.GET("/laughing/:size", server.HandleGetLaughingAvatar())
	server.echoRouter.GET("/smiley/:size/:name", server.HandleGetSmileyAvatar())
	server.echoRouter.GET("/smiley/:size", server.HandleGetSmileyAvatar())
}
