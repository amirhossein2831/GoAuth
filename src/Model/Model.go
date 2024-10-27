package Model

type Model interface {
	TableName() string
}

// Models is please you should register you model to be Migrate
func Models() []interface{} {
	return []interface{}{
		User{},
	}
}
