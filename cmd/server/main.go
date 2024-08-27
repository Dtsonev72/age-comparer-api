// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	//	"github.com/Dtsonev72/receipt-processor-api/internal/receipt"
)

func main() {
	receiptService := receipt.NewService()

	http.HandleFunc("/receipts/process", receiptService.ProcessReceipt)
	http.HandleFunc("/receipts/", receiptService.GetPoints)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
