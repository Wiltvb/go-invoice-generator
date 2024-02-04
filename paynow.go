package generator

import (
	"bytes"
	b64 "encoding/base64"
	"github.com/go-pdf/fpdf"
	"image"
)

type PayNow struct {
	UEN    string `json:"uen,omitempty"`
	Image  []byte `json:"logo,omitempty"` // Logo byte array
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (p *PayNow) appendPayNowTODoc(x float64,
	y float64,
	doc *Document) float64 {
	// Logo
	if p.Image != nil {
		// Create filename
		fileName := b64.StdEncoding.EncodeToString([]byte(p.UEN))

		// Create reader from logo bytes
		ioReader := bytes.NewReader(p.Image)

		// Get image format
		_, format, _ := image.DecodeConfig(bytes.NewReader(p.Image))

		// Register image in pdf
		imageInfo := doc.pdf.RegisterImageOptionsReader(fileName, fpdf.ImageOptions{
			ImageType: format,
		}, ioReader)

		if imageInfo != nil {
			var imageOpt fpdf.ImageOptions
			imageOpt.ImageType = format

			currentY := doc.pdf.GetY() + float64(p.Height) - 10

			doc.pdf.ImageOptions(fileName, doc.pdf.GetX(), currentY, 0, float64(p.Height), false, imageOpt, 0, "")

		}
	}

	return doc.pdf.GetY()
}

func (p *PayNow) appendPaynowToDoc(doc *Document) float64 {
	x, y, _, _ := doc.pdf.GetMargins()
	return p.appendPayNowTODoc(x, y, doc)
}
