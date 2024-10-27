package Model

type Model interface {
	TableName() string
}

func Models() []interface{} {
	return []interface{}{
		User{},
	}
}
