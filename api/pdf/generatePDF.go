package pdf

import (
	"fmt"
	"log"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//PDF generator
func PDF(rID string) {
	// Create new PDF generator
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Grayscale.Set(false)

	// Create a new input page from an URL
	fmt.Println(rID)

	url := "http://localhost:3001/report/"
	page := wkhtml.NewPage(url + rID)

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(1.30)

	// Add to document
	pdfg.AddPage(page)
	//fmt.Println(pdfg)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("api/pdf/report" + rID + ".pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done

}
