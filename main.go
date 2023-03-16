package main

import (
	"fmt"
	"github.com/ledongthuc/pdf"
)

func main() {
	readPdf("./Jan.pdf")
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			println(">>>> row: ", row.Position)
			for _, word := range row.Content {
				fmt.Println(word.S)
			}
		}
	}
	return "", nil
}
