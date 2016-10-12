# Clinical trials
Search for clinical trials from Health Canada. A go package and a command line
interface (CLI). 

# Example (cli) 
Search for all breast cancer studies: 

```
$ clinical-trials -medicalConditionName="breast cancer"
Found 69 studies:
	Drug name: 18F-FLUORODEOXYGLUCOSE
	CTA protocol title: IMPACT OF 18 F-FDG PET-CT VERSUS CONVENTIONAL STAGING IN THE MANAGEMENT OF PATIENTS PRESENTING WITH CLINICAL STAGE III BREAST CANCER (PETABC)
	Medical condition: BREAST CANCER
	Study population: ADULT FEMALE, ADULT MALE
	Trial status: PENDING
	URL: https://health-products.canada.ca/ctdb-bdec/info.do?submissionNo=194153

	Drug name: ABEMACICLIB (LY2835219)
	CTA protocol title: A PHASE 2 STUDY OF ABEMACICLIB IN PATIENTS WITH BRAIN METASTASES SECONDARY TO HORMONE RECEPTOR POSITIVE BREAST CANCER
	Medical condition: BREAST CANCER METASTATIC
	Study population: ADULT FEMALE
	Trial status: ONGOING
	URL: https://health-products.canada.ca/ctdb-bdec/info.do?submissionNo=181445
    
    ...
    
```

# Install 
- Install [go](http://golang.org) and set it up accordingly. 
- Get [goquery](https://github.com/PuerkitoBio/goquery): 'go get github.com/PuerkitoBio/goquery'
- Get [clinical-trials](github.com/fjukstad/clinical-trials) `go get github.com/fjukstad/clinical-trials`
- (optional) Install the CLI `cd $GOPATH/src/github.com/fjukstad/clinical/clinical-trials/ && go install `

# Usage (cli) 
```
Usage of clinical-trials:
  -controlNumber string
    	control number
  -drugName string
    	drug name
  -endDateFrom string
    	study end date from (yyyy-mm-dd)
  -endDateTo string
    	study end date to (yyyy-mm-dd)
  -medicalConditionName string
    	medical condition
  -nolDateFrom string
    	no objection letter date from (yyyy-mm-dd)
  -nolDateTo string
    	no ojection letter date to (yyyy-mm-dd)
  -protocolNo string
    	protocol number
  -protocolTitle string
    	protocol title
  -sponsorName string
    	sponsor name
  -startDateFrom string
    	study start from date (yyyy-mm-dd)
  -startDateTo string
    	study start to date (yyyy-mm-dd)
  -statusId string
    	trial status
  -studyPopulationId string
    	study population id
```
