package db

func Init() error {
	err := initMySQL()
	if err != nil {
		return err
	}
	return nil
}