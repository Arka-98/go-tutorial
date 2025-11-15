package utils

import "rsc.io/quote/v4"

func GetQuote() string {
	return quote.Go()
}
