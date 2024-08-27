package receipt

import (
	"github.com/Dtsonev72/receipt-processor-api/internal/storage"
)

type Service struct {
	store storage.ReceiptStore
}
