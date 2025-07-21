package halper

import "database/sql"

func CommitOrRollbeck(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollbeck := tx.Rollback()
		PanicIfError(errorRollbeck)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
