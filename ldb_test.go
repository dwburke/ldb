package ldb_test

import (
	"os"
	"testing"

	"github.com/dwburke/ldb"
)

func TestOpen(t *testing.T) {

	ldb, err := ldb.NewLdb(&ldb.Config{
		Datadir:         "test_state",
		ErrorOnNotFound: false,
	})
	if err != nil {
		t.Error(err)
	}
	defer ldb.Close()
	defer os.RemoveAll("test_state")

	//if result != sc.Expected {
	//t.Errorf("result is not as expected: wanted [%s], got [%s]", sc.Expected, result)
	//}

}
