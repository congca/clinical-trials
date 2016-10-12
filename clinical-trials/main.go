package main

import (
	"flag"
	"fmt"
)

import "github.com/fjukstad/clinical"

func main() {

	medicalConditionName := flag.String("medicalConditionName", "", "medical condition")
	protocolTitle := flag.String("protocolTitle", "", "protocol title")
	drugName := flag.String("drugName", "", "drug name")
	sponsorName := flag.String("sponsorName", "", "sponsor name")
	studyPopulationId := flag.String("studyPopulationId", "", "study population id")
	statusId := flag.String("statusId", "", "trial status")
	nolDateFrom := flag.String("nolDateFrom", "", "no objection letter date from (yyyy-mm-dd)")
	nolDateTo := flag.String("nolDateTo", "", "no ojection letter date to (yyyy-mm-dd)")
	startDateFrom := flag.String("startDateFrom", "", "study start from date (yyyy-mm-dd)")
	startDateTo := flag.String("startDateTo", "", "study start to date (yyyy-mm-dd)")
	endDateFrom := flag.String("endDateFrom", "", "study end date from (yyyy-mm-dd)")
	endDateTo := flag.String("endDateTo", "", "study end date to (yyyy-mm-dd)")
	protocolNo := flag.String("protocolNo", "", "protocol number")
	controlNumber := flag.String("controlNumber", "", "control number")

	flag.Parse()

	studies := clinical.Search(*medicalConditionName,
		*protocolTitle,
		*drugName,
		*sponsorName,
		*studyPopulationId,
		*statusId,
		*nolDateFrom,
		*nolDateTo,
		*startDateFrom,
		*startDateTo,
		*endDateFrom,
		*endDateTo,
		*protocolNo,
		*controlNumber)

	fmt.Println("Found", len(studies), "studies:")

	for _, study := range studies {
		fmt.Println("\tDrug name:", study.DrugName)
		fmt.Println("\tCTA protocol title:", study.CTAProtocolTitle)
		fmt.Println("\tMedical condition:", study.MedicalCondition)
		fmt.Println("\tStudy population:", study.StudyPopulation)
		fmt.Println("\tTrial status:", study.TrialStatus)
		fmt.Println("\tURL:", study.Url)
		fmt.Println()
	}
}
