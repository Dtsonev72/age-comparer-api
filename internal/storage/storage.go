package storage

type ReceiptStore interface {
	SaveReceipt(id string, receipt interface{})
	GetReceipt(id string) (interface{}, bool)
}
