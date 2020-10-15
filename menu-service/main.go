package main

func main(){
	router := mux.NewRouter()

	router.Handler("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("menu service listen on port : 8000")
	log.Panic(http.ListenAndServe(":8000", router))
}