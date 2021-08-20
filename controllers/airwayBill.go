package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jun2900/digiFormTest/database"
	"github.com/jun2900/digiFormTest/models"
	"github.com/xuri/excelize/v2"
	"gorm.io/datatypes"
)

type Output struct {
	Lokasi     models.Lokasi     `json:"lokasi"`
	AirwayBill models.AirwayBill `json:"airway_bill"`
}

func insertLokasi(xlsxFile *excelize.File, i int) models.Lokasi {
	provinsiAsal, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("B%d", i))
	tipeAsal, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("C%d", i))
	kotaAsal, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("D%d", i))
	kecamatanAsal, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("E%d", i))
	kodeposAsal, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("F%d", i))
	alamatAsal, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("G%d", i))
	provinsiTujuan, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("H%d", i))
	tipeTujuan, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("I%d", i))
	kotaTujuan, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("J%d", i))
	kecamatanTujuan, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("K%d", i))
	kodeposTujuan, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("L%d", i))
	alamatTujuan, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("M%d", i))
	service, _ := xlsxFile.GetCellValue("Sheet1", fmt.Sprintf("R%d", i))

	lokasiRow := models.Lokasi{
		ProvinsiAsal:    provinsiAsal,
		TipeAsal:        tipeAsal,
		KotaAsal:        kotaAsal,
		KecamatanAsal:   kecamatanAsal,
		KodeposAsal:     kodeposAsal,
		AlamatAsal:      alamatAsal,
		ProvinsiTujuan:  provinsiTujuan,
		TipeTujuan:      tipeTujuan,
		KotaTujuan:      kotaTujuan,
		KecamatanTujuan: kecamatanTujuan,
		KodeposTujuan:   kodeposTujuan,
		AlamatTujuan:    alamatTujuan,
		Service:         service,
	}
	return lokasiRow
}

func insertAirwayBill(c *fiber.Ctx, service string, kodeposTujuan string, t time.Time) models.AirwayBill {
	airwayBillRow := models.AirwayBill{
		SendDate:      datatypes.Date(t),
		Service:       service,
		Origin:        c.FormValue("origin"),
		KodeposTujuan: kodeposTujuan,
	}

	return airwayBillRow
}

func AirwayBill(c *fiber.Ctx) error {
	db := database.DBConn

	file, err := c.FormFile("document")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))

	xlsxFile, err := excelize.OpenFile(fmt.Sprintf("./%s", file.Filename))
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	var result = []Output{}

	layoutISO := "2006-01-02"
	t, _ := time.Parse(layoutISO, c.FormValue("send_date"))
	for i := 4; i < 14; i++ {
		lokasiRow := insertLokasi(xlsxFile, i)
		airwayBillRow := insertAirwayBill(c, lokasiRow.Service, lokasiRow.KodeposTujuan, t)

		db.Create(&lokasiRow)
		db.Create(&airwayBillRow)
		result = append(result, Output{Lokasi: lokasiRow, AirwayBill: airwayBillRow})
	}

	return c.JSON(result)
}
