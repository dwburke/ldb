package ldb_test

import (
	"testing"

	"github.com/dwburke/ldb"
)

func TestOpen(t *testing.T) {

	ldb, err := ldb.NewLdb(&ldb.Config{
		Datadir:         "test",
		ErrorOnNotFound: false,
	})
	if err != nil {
		t.Error(err)
	}
	defer ldb.Close()

	//if result != sc.Expected {
	//t.Errorf("result is not as expected: wanted [%s], got [%s]", sc.Expected, result)
	//}

}
