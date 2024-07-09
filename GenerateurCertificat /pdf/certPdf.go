package pdf

import (
	"GenerateurCertificat/cert"
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
)

// manipulation of gopdf


type PdfSaver struct {
	OutputDir string
}
// creation of pdf 
func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}
	// initialize pdf saver
 p = &PdfSaver{
	OutputDir: outputdir,
 }
 return p, nil
}

// implementation of the saver interface of our cert package
func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")

	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	// Background
	background(pdf)

	//header
	header(pdf, &cert)
	pdf.Ln(30)
	//body
	pdf.SetFont("Helvetica","I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	//body -Student name
	pdf.SetFont("Times","B", 40)
	pdf.WriteAligned(0,50, cert.Name, "C")
	pdf.Ln(30)

	//body - participation
	pdf.SetFont("Helvetica","I",20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// body -Date
	pdf.SetFont("Helvetica","I",15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")
	
	// footer
	footer(pdf)

	//save file
	filename:= fmt.Sprintf("%v.pdf", cert.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err 
}
fmt.Printf("Saved certificate to %v\n", path)
return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()
	pdf.ImageOptions("img/cadre.png",
	0, 0,
	pageWidth, pageHeight,
	false, opts, 0,"")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType:"png",
	}
	// gopher R/L
	margin := 30.0
	x:= 0.0
	imageWidth := 50.0
	filename := "img/gopher.png"
	pdf.ImageOptions(filename,
		x+margin, 20,
		imageWidth, 0,
		false, opts, 0,"")

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename,
		x-margin, 20,
		imageWidth, 0,
		false, opts, 0,"")	

	pdf.SetFont("Helvetica", "",40)
	pdf.WriteAligned(0, 70, c.LabelCompletion, "C") // pour centrer "C" sinon L = left, R = right
}

func footer(pdf *gofpdf.Fpdf){
	opts := gofpdf.ImageOptions{
		ImageType:"png",
	}
	imageWidth :=50.0
	filename := "img/tanpom.png"
	pageWidth, pageHeight := pdf.GetPageSize()
	x := pageWidth - imageWidth - 25.0
	y := pageHeight - imageWidth - 30.0
	pdf.ImageOptions(filename, 
		x, y, 
		imageWidth, 0, 
		false, opts, 0, "")
		// 200, 130, 50, 0, false, opts, 0,"")
	
}