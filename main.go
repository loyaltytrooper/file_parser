package main

import (
	"fmt"
	"github.com/ledongthuc/pdf"
)

func main() {
	str, err := readPdf("./Jan.pdf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

func readPdf(path string) (string, error) {
	//reachedStartPoint := false
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
				fmt.Println(word.Y)
			}
		}
	}
	return "", nil
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
//	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
//		p := r.Page(pageIndex)
//		if p.V.IsNull() {
//			continue
//		}
//
//		rows, _ := p.GetTextByColumn()
//		for _, row := range rows {
//			println(">>>> row: ", row.Position)
//			for _, word := range row.Content {
//				fmt.Println(word.Y)
//			}
//		}
//	}
//	return "", nil
//}

//func readPdf(path string) (string, error) {
//	f, r, err := pdf.Open(path)
//	// remember close file
//	defer f.Close()
//	if err != nil {
//		return "", err
//	}
//	var buf bytes.Buffer
//	b, err := r.GetPlainText()
//	if err != nil {
//		return "", err
//	}
//	buf.ReadFrom(b)
//	return buf.String(), nil
//}
