package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"strconv"
	"strings"
)

func main() {
	//str, err :=
	readPdf("./feb.pdf")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(str)
}

//func readPdf(path string) (string, error) {
//	//reachedStartPoint := false
//	f, r, err := pdf.Open(path)
//	defer func() {
//		_ = f.Close()
//	}()
//	if err != nil {
//		return "", err
//	}
//	totalPage := r.NumPage()
//
//	var transactions models.Transactions
//
//	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
//		p := r.Page(pageIndex)
//		if p.V.IsNull() {
//			continue
//		}
//
//		var currentBalance float64
//		var prevBalance float64
//		//var prevDescriptionY float64
//		//var prevDescriptionX float64
//		rows, _ := p.GetTextByRow()
//		for _, row := range rows {
//			transactionTime, err := time.Parse("02-Jan-2006", row.Content[0].S)
//			if err != nil {
//				//if row.Content.Len() == 1 {
//				//	if math.Abs(prevDescriptionY-row.Content[0].Y) < 11 {
//				//		transactions.Txns[len(transactions.Txns)-1].Description += row.Content[0].S
//				//	}
//				//}
//				continue
//			} else {
//				currentBalance, err = BeautifyCommaNumber(row.Content[row.Content.Len()-1].S)
//				if err != nil {
//					fmt.Println(err)
//				} else {
//					if row.Content.Len() == 3 {
//						txnAmount, err := BeautifyCommaNumber(row.Content[row.Content.Len()-2].S)
//						if err != nil {
//							//prevDescriptionY = row.Content[row.Content.Len()-2].Y
//							//prevDescriptionX = row.Content[row.Content.Len()-2].X
//							transactions.Txns = append(transactions.Txns, models.Transaction{
//								Date:        transactionTime,
//								Description: row.Content[row.Content.Len()-2].S,
//								Credit:      0,
//								Debit:       0,
//								FinalAmount: currentBalance,
//							})
//						} else if prevBalance < currentBalance {
//							transactions.Txns = append(transactions.Txns, models.Transaction{
//								Date:        transactionTime,
//								Description: "",
//								Credit:      txnAmount,
//								Debit:       0,
//								FinalAmount: currentBalance,
//							})
//						} else {
//							transactions.Txns = append(transactions.Txns, models.Transaction{
//								Date:        transactionTime,
//								Description: "",
//								Credit:      0,
//								Debit:       txnAmount,
//								FinalAmount: currentBalance,
//							})
//						}
//					} else if row.Content.Len() == 4 {
//						txnAmount, err := BeautifyCommaNumber(row.Content[row.Content.Len()-2].S)
//						if err != nil {
//							fmt.Println(err)
//						} else if prevBalance < currentBalance {
//							//prevDescriptionY = row.Content[row.Content.Len()-3].Y
//							//prevDescriptionX = row.Content[row.Content.Len()-3].X
//							transactions.Txns = append(transactions.Txns, models.Transaction{
//								Date:        transactionTime,
//								Description: row.Content[row.Content.Len()-3].S,
//								Credit:      txnAmount,
//								Debit:       0,
//								FinalAmount: currentBalance,
//							})
//						} else if prevBalance > currentBalance {
//							fmt.Println("Debit")
//							//prevDescriptionY = row.Content[row.Content.Len()-3].Y
//							//prevDescriptionX = row.Content[row.Content.Len()-3].X
//							transactions.Txns = append(transactions.Txns, models.Transaction{
//								Date:        transactionTime,
//								Description: row.Content[row.Content.Len()-3].S,
//								Credit:      0,
//								Debit:       txnAmount,
//								FinalAmount: currentBalance,
//							})
//						}
//					}
//					prevBalance = currentBalance
//				}
//			}
//		}
//	}
//	for _, x := range transactions.Txns {
//		fmt.Println(x)
//	}
//	return "", nil
//}

func BeautifyCommaNumber(str string) (float64, error) {
	txnArr := strings.Split(str, ",")
	txn := ""
	for _, str := range txnArr {
		txn += str
	}
	return strconv.ParseFloat(txn, 64)
}

func readPdf(fileName string) {
	//err := api.DecryptFile(fileName, "act_feb", model.NewAESConfiguration("RAJA2712", "RAJA2712", 128))
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	err := api.ExtractContentFile("feb.pdf", "act_feb", nil, model.NewAESConfiguration("RAJA2712", "RAJA2712", 128))
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(ctx.Title)
	//fmt.Println(ctx.Valid)
	//fmt.Println(ctx.CreationDate)
	//fmt.Println(ctx.Pages())
	//fmt.Println(ctx.Keywords)
	//fmt.Println(ctx.OwnerPWNew)
	//fmt.Println(ctx.XRefTable.CurPage)
}
