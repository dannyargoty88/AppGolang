package pdf

import (
	"github.com/signintech/gopdf"
)

/**
files: array de String que contine la ruta del pdf
fileMerge: String que define la ruta del nuevo pdf
numPag: cantidad de paginas que contiene cada pdf

Ejemplo:

pdfMerge := []string{}
pdfMerge = append(facturasMerge, "./impresion/factura_01.pdf")
pdfMerge = append(facturasMerge, "./impresion/factura_02.pdf")
MergePdf(pdfMerge, "./impresion/merge/FacturaMerge.pdf", 2)
*/

func MergePdf(files []string, fileMerge string, numPag int) error {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	for i := 0; i < len(files); i++ {
		for j := 1; j <= numPag; j++ {
			pdf.AddPage()
			pdf.UseImportedTemplate(pdf.ImportPage(files[i], j, "/MediaBox"), 0, 0, 0, 0)
		}
	}

	err := pdf.WritePdf(fileMerge)
	if err != nil {
		return err
	}

	return nil
}
