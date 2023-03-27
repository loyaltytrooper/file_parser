package helper

import (
	"strconv"
	"strings"
)

func SimplifyCommaNumber(str string) (float64, error) {
	txnArr := strings.Split(str, ",")
	txn := ""
	for _, str := range txnArr {
		txn += str
	}
	return strconv.ParseFloat(txn, 64)
}
