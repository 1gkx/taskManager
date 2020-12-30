package store

func migrate() {

	// Миграция схем
	x.AutoMigrate(&User{})

}
