package operation

type Movie struct {
	Title       string `json:"title" binding:"min=2,max=10"`
	Description string `json:"desc" binding:"max=20"`
	Trailer     string `json:"trailer" binding:"required,url"`
	Price       string `json:"price" binding:"required"`
	LeadActor  Person `json:"actor" binding:"required"`
}

type Person struct {
	FirstName string `json: "first" binding:"required"`
	LastName string `json: "last" binding:"required"`
	Age int `json: "age" binding:"get=1,lte=130"`
	Email string `json: "email" validate:"required,email"`
}
