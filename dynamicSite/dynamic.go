package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	var res string
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, close := chromedp.NewContext(allocCtx)
	defer close()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://magnit.ru/catalog/?categoryId=4893`),
		chromedp.Click(".city-leaving-cancel"), // choose city
		chromedp.Click(".cookies__button"), // cookies
		chromedp.Click(".paginate__more"), // more products
		chromedp.Text(".catalog-page__product-grid", &res),
	); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}