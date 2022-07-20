package main

import (
	"github.com/xuri/excelize/v2"
	"log"
)

type Fulldata struct {
	Part_number      string                    `json:"part_number" bson:"_id"`
	Freq_min         string                    `json:"freq_min" bson:"freq___min"`
	Freq_max         string                    `json:"freq_max" bson:"freq___max"`
	Z_Min            map[string]string         `json:"Z Min,omitempty" bson:"z___min"`
	Frequency_values []Equivalent_circuit_data `json:"frequency_values" bson:"frequency___values"`
}

type Equivalent_circuit_data struct {
	Freq_MHZ string `json:"Freq (MHz)" bson:"freq___mhz"`
	Z_Ohm    string `json:"Z (Ohm)" bson:"z___ohm"`
	Re_Z_Ohm string `json:"Re[Z] (Ohm)" bson:"re___z___ohm"`
	IM_Z_Ohm string `json:"Im[Z] (Ohm)" bson:"im___z___ohm"`
	Cap_F    string `json:"Cap (F)" bson:"cap___f"`
	Phase    string `json:"Phase (º)" bson:"phase"`
	L_H      string `json:"L (H),omitempty" bson:"l___h"`
}

func returnExcelData(filePath string, sheetIndex int) Fulldata {
	f, sheetName := openExcel(filePath, sheetIndex)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return mappingColumn(rows)
}

func openExcel(filePath string, sheetIndex int) (*excelize.File, string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 取得sheet name
	sheetName := f.GetSheetName(sheetIndex)

	return f, sheetName
}

func mappingColumn(rows [][]string) Fulldata {

	ecd := Equivalent_circuit_data{
		Freq_MHZ: "",
		Z_Ohm:    "",
		Re_Z_Ohm: "",
		IM_Z_Ohm: "",
		Cap_F:    "",
		Phase:    "",
		L_H:      "",
	}

	fd := Fulldata{
		Part_number:      "A767EB226M1VLAE050",
		Freq_min:         "0.0001",
		Freq_max:         "10",
		Z_Min:            map[string]string{"Freq (MHz)": "5.04E-01", "Z (Ohm)": "2.67E-02"},
		Frequency_values: []Equivalent_circuit_data{},
	}

	for rowIndex, row := range rows {
		if rowIndex >= 4 {
			for columnIndex, colCell := range row {
				switch columnIndex {
				case 0:
					ecd.Freq_MHZ = colCell
				case 1:
					ecd.Z_Ohm = colCell
				case 2:
					ecd.Re_Z_Ohm = colCell
				case 3:
					ecd.IM_Z_Ohm = colCell
				case 4:
					ecd.Cap_F = colCell
				case 5:
					ecd.Phase = colCell
				case 6:
					if colCell != "" {
						ecd.L_H = colCell
					} else {
						ecd.L_H = ""
					}
				}
			}
			fd.Frequency_values = append(fd.Frequency_values, ecd)
		}
	}
	return fd
}
