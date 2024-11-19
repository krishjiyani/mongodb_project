package db

type Student struct {
	ID       string   `bson:"_id"`
	Name     string   `bson:"name"`
	RollNo   int      `bson:"roll_no"`
	Division string   `bson:"division"`
	Class    string   `bson:"class"`
	Subjects []string `bson:"subjects"`
}
