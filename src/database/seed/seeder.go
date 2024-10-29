package seed

// Seed Put Seeder in here
func Seed() error {
	err := AdminSeeder()
	if err != nil {
		return err
	}

	return nil
}
