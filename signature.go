package generator

import (
	"bytes"
	b64 "encoding/base64"
	"github.com/go-pdf/fpdf"
	"image"
)

type Signature struct {
	SignatureName     string `json:"signature_name,omitempty"`
	SignatureJobTitle string `json:"job_title,omitempty"`
	Signature         []byte `json:"signature,omitempty"`
	SignatureWidth    int    `json:"width,omitempty"`
	SignatureHeight   int    `json:"height,omitempty"`
}

func (s *Signature) appendSignature(x float64,
	y float64,
	doc *Document) float64 {
	// Logo
	if s.Signature != nil {
		// Create filename
		fileName := b64.StdEncoding.EncodeToString([]byte("signature"))

		// Create reader from logo bytes
		ioReader := bytes.NewReader(s.Signature)

		// Get image format
		_, format, _ := image.DecodeConfig(bytes.NewReader(s.Signature))

		// Register image in pdf
		imageInfo := doc.pdf.RegisterImageOptionsReader(fileName, fpdf.ImageOptions{
			ImageType: format,
		}, ioReader)

		if imageInfo != nil {
			var imageOpt fpdf.ImageOptions
			imageOpt.ImageType = format

			currentY := doc.pdf.GetY() + 30

			doc.pdf.ImageOptions(fileName, doc.pdf.GetX()+10, currentY, 0, float64(s.SignatureHeight), false, imageOpt, 0, "")

		}
	}

	// Draw rect
	doc.pdf.SetFillColor(105, 105, 105)
	doc.pdf.Rect(doc.pdf.GetX(), doc.pdf.GetY()+40, 55, 0.5, "F")

	doc.pdf.SetY(doc.pdf.GetY() + 45)
	doc.pdf.SetX(10)
	doc.pdf.CellFormat(80, 4, doc.encodeString(s.SignatureName), "0", 0, "L", false, 0, "")

	doc.pdf.SetY(doc.pdf.GetY() + 5)
	doc.pdf.SetX(10)
	doc.pdf.CellFormat(80, 4, doc.encodeString(s.SignatureJobTitle), "0", 0, "L", false, 0, "")

	doc.pdf.SetY(doc.pdf.GetY() + 5)
	doc.pdf.SetX(10)
	doc.pdf.CellFormat(80, 4, doc.encodeString("WILT VENTURE BUILDER PTE. LTD."), "0", 0, "L", false, 0, "")

	return doc.pdf.GetY()
}

func (s *Signature) appendSignatureToDoc(doc *Document) float64 {
	x, y, _, _ := doc.pdf.GetMargins()
	return s.appendSignature(x, y, doc)
}
