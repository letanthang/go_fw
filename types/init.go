package types

type Student struct {
	ID        int    `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Age       int    `bson:"age"`
	ClassName string `bson:"class"`
	Email     string `bson:"email"`
}
