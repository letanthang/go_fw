package db

import (
	"context"

	"github.com/letanthang/go_fw/types"
)

func GetAllStudent() (*[]types.Student, error) {

	var students []types.Student
	filter := map[string]interface{}{}

	cursor, err := Client.Database(DBName).Collection(Col).Find(
		context.TODO(),
		filter,
	)

	if err != nil {
		return nil, err
	}

	cursor.All(context.TODO(), &students)

	// cursor //.Next .FetchAll

	return &students, nil
}
