package httpapp

func StartApplication() {
	router := routes()
	router.Run(":8080")
}