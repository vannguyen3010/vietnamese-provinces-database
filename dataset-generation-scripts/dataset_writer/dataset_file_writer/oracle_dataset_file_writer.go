package dataset_writer

import (
	"bufio"
	"fmt"
	vn_common "github.com/thanglequoc-vn-provinces/v2/common"
	"log"
	"os"
	"time"
)

const insertProvinceOracleTemplate string = "\tINTO provinces(code,name,name_en,full_name,full_name_en,code_name,administrative_unit_id,administrative_region_id) VALUES('%s','%s','%s','%s','%s','%s',%d,%d)"

type OracleDatasetFileWriter struct {
	OutputFilePath string
}

func (w *OracleDatasetFileWriter) WriteToFile(
	regions []vn_common.AdministrativeRegion,
	administrativeUnits []vn_common.AdministrativeUnit,
	provinces []vn_common.Province,
	districts []vn_common.District,
	wards []vn_common.Ward) error {

	fileTimeSuffix := getFileTimeSuffix()
	outputFilePath := fmt.Sprintf(w.OutputFilePath, fileTimeSuffix)

	file, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Unable to write to file", err)
		panic(err)
	}

	dataWriter := bufio.NewWriter(file)
	dataWriter.WriteString("/* === Vietnamese Provinces Database Dataset for Oracle === */\n")
	dataWriter.WriteString(fmt.Sprintf("/* Created at:  %s */\n", time.Now().Format(time.RFC1123Z)))
	dataWriter.WriteString("/* Reference: https://github.com/ThangLeQuoc/vietnamese-provinces-database */\n")
	dataWriter.WriteString("/* =============================================== */\n\n")

	dataWriter.WriteString("-- DATA for administrative_regions --\n")
	for _, r := range regions {
		insertLine := fmt.Sprintf(insertAdministrativeRegionTemplate,
			r.Id, r.Name, r.NameEn, r.CodeName, r.CodeNameEn)
		dataWriter.WriteString(insertLine + "\n")
	}
	dataWriter.WriteString("-- ----------------------------------\n\n")

	dataWriter.WriteString("-- DATA for administrative_units --\n")
	for _, u := range administrativeUnits {
		insertLine := fmt.Sprintf(insertAdministrativeUnitTemplate,
			u.Id, u.FullName, u.FullNameEn, u.ShortName, u.ShortNameEn, u.CodeName, u.CodeNameEn)
		dataWriter.WriteString(insertLine + "\n")
	}
	dataWriter.WriteString("-- ----------------------------------\n\n")

	// variable to generate batch insert statement
	counter := 0
	isAppending := false

	dataWriter.WriteString("-- DATA for provinces --\n")

	for i, p := range provinces {
		if !isAppending {
			dataWriter.WriteString("INSERT ALL\n")
		}
		dataWriter.WriteString(
			fmt.Sprintf(insertProvinceOracleTemplate, p.Code, escapeSingleQuote(p.Name), escapeSingleQuote(p.NameEn), escapeSingleQuote(p.FullName),
				escapeSingleQuote(p.FullNameEn), p.CodeName, p.AdministrativeUnitId, p.AdministrativeRegionId))
		counter++

		// the batch insert statement batch reach limit, break and create a new batch insert statement
		// In oracle, the last insert batch statement require a dummy select after multiple INSERT ALL INTO statements
		if counter == batchInsertItemSize || i == len(provinces)-1 {
			isAppending = false
			dataWriter.WriteString("\n\tSELECT 1 FROM DUAL;\n\n")
			counter = 0 // reset counter
		} else {
			dataWriter.WriteString("\n")
			isAppending = true
		}
	}
	dataWriter.WriteString("-- ----------------------------------\n\n")

	const insertDistrictOracleTemplate string = "\tINTO districts(code,name,name_en,full_name,full_name_en,code_name,province_code,administrative_unit_id) VALUES('%s','%s','%s','%s','%s','%s','%s',%d)"
	dataWriter.WriteString("-- DATA for districts --\n")
	counter = 0
	isAppending = false
	for i, d := range districts {
		if !isAppending {
			dataWriter.WriteString("INSERT ALL\n")
		}
		dataWriter.WriteString(
			fmt.Sprintf(insertDistrictOracleTemplate, d.Code, escapeSingleQuote(d.Name), escapeSingleQuote(d.NameEn), escapeSingleQuote(d.FullName),
				escapeSingleQuote(d.FullNameEn), d.CodeName, d.ProvinceCode, d.AdministrativeUnitId))
		counter++

		// the batch insert statement batch reach limit, break and create a new batch insert statement
		// In oracle, the last insert batch statement require a dummy select after multiple INSERT ALL INTO statements
		if counter == batchInsertItemSize || i == len(districts)-1 {
			isAppending = false
			dataWriter.WriteString("\n\tSELECT 1 FROM DUAL;\n\n")
			counter = 0 // reset counter
		} else {
			dataWriter.WriteString("\n")
			isAppending = true
		}
	}
	dataWriter.WriteString("-- ----------------------------------\n\n")

	// ward insert statement
	const insertWardOracleTemplate string = "\tINTO wards(code,name,name_en,full_name,full_name_en,code_name,district_code,administrative_unit_id) VALUES('%s','%s','%s','%s','%s','%s','%s',%d)"

	dataWriter.WriteString("-- DATA for wards --\n")
	counter = 0
	isAppending = false
	for i, d := range wards {
		if !isAppending {
			dataWriter.WriteString("INSERT ALL\n")
		}
		dataWriter.WriteString(
			fmt.Sprintf(insertWardOracleTemplate, d.Code, escapeSingleQuote(d.Name), escapeSingleQuote(d.NameEn), escapeSingleQuote(d.FullName),
				escapeSingleQuote(d.FullNameEn), d.CodeName, d.DistrictCode, d.AdministrativeUnitId))
		counter++

		// the batch insert statement batch reach limit, break and create a new batch insert statement
		// In oracle, the last insert batch statement require a dummy select after multiple INSERT ALL INTO statements
		if counter == batchInsertItemSize || i == len(wards)-1 {
			isAppending = false
			dataWriter.WriteString("\n\tSELECT 1 FROM DUAL;\n\n")
			counter = 0 // reset counter
		} else {
			dataWriter.WriteString("\n")
			isAppending = true
		}
	}
	dataWriter.WriteString("-- ----------------------------------\n\n")

	dataWriter.WriteString("-- END OF SCRIPT FILE --\n")
	dataWriter.Flush()
	file.Close()
	return nil
}
