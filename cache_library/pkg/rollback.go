package pkg

// Rollback : in order to maintain data integrity we'll need some rollback strategy
type Rollback interface {
	RollbackChanges()
	UpdateBackUp()
}
