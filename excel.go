package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func (a *API) generateExcel(mallId int64) (*excelize.File, error) {
	customers, err := a.s.GetCustomersByMallId(mallId)
	if err != nil {
		return nil, err
	}
	f := excelize.NewFile()
	sheet := "Sheet1"
	// make titles
	f.SetCellValue(sheet, "A1", "user_id")
	f.SetCellValue(sheet, "B1", "user_name")
	f.SetCellValue(sheet, "C1", "user_location")
	f.SetCellValue(sheet, "D1", "mall_id")
	f.SetCellValue(sheet, "E1", "item_id")
	f.SetCellValue(sheet, "F1", "buy_count")
	for i, customer := range customers {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), customer.Id)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), customer.Username)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), customer.Location)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), customer.MallId)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), customer.ItemId)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", i+2), customer.BuyCount)
	}
	return f, nil
}
