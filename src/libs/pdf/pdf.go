package pdf

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"text/template"

	models "app/src/models"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func CreatePdf(data_pdf models.Pdf) error {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// construct your html
	body := data_pdf.BodyPdf

	t, err := template.ParseFiles(data_pdf.BodyTemplate)
	if err != nil {
		log.Panic(err)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, body)
	if err != nil {
		log.Panic(err)
	}

	html := buf.String()

	err = chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}

			return page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx)
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			return ioutil.WriteFile(data_pdf.PrintPath, buf, 0644)
		}),
	)

	if err != nil {
		return err
	}

	return nil
}
