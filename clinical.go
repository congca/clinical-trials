package clinical

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

type Study struct {
	Url              string
	DrugName         string
	CTAProtocolTitle string
	MedicalCondition string
	StudyPopulation  string
	TrialStatus      string
}

func Search(medicalConditionName, protocolTitle, drugName, sponsorName,
	studyPopulationId, statusId, nolDateFrom, nolDateTo, startDateFrom, startDateTo,
	endDateFrom, endDateTo, protocolNo, controlNumber string) []Study {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{Jar: jar}

	URL := "https://health-products.canada.ca/ctdb-bdec/search-recherche.do"

	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.PostForm(URL,
		url.Values{"medicalConditionName": {medicalConditionName},
			"protocolTitle":     {protocolTitle},
			"drugName":          {drugName},
			"sponsorName":       {sponsorName},
			"studyPopulationId": {studyPopulationId},
			"statusId":          {statusId},
			"nolDateFrom":       {nolDateFrom},
			"nolDateTo":         {nolDateTo},
			"startDateFrom":     {startDateFrom},
			"startDateTo":       {startDateTo},
			"endDateFrom":       {endDateFrom},
			"endDateTo":         {endDateTo},
			"protocolNo":        {protocolNo},
			"controlNumber":     {controlNumber},
			"action":            {"Search"},
		})

	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	studies := []Study{}

	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		rows := s.Find("tr")
		rows.Each(func(j int, se *goquery.Selection) {
			entry := se.Find("td")
			study := Study{}
			entry.Each(func(k int, sel *goquery.Selection) {
				data := sel.Text()
				data = strings.TrimSpace(data)
				if k == 0 {
					link := sel.Find("a")
					href, _ := link.Attr("href")
					study.Url = "https://health-products.canada.ca" + href
					study.DrugName = data
				} else if k == 1 {
					study.CTAProtocolTitle = data
				} else if k == 2 {
					study.MedicalCondition = data
				} else if k == 3 {
					study.StudyPopulation = data
				} else if k == 4 {
					study.TrialStatus = data
				}
			})
			studies = append(studies, study)
		})
		return
	})

	if len(studies) > 0 {
		// first row is always empty in the table
		studies = studies[1:len(studies)]
	}
	return studies
}
