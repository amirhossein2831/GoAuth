package models

type Model interface {
	TableName() string
}

// Models is please you should register you models to be Migrate
func Models() []interface{} {
	return []interface{}{
		User{},
		Token{},
	}
}

func ToModel[T Model](models []T) []Model {
	var all []Model
	for _, model := range models {
		all = append(all, model)
	}
	return all
}
