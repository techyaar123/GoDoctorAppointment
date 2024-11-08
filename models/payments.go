package models

type Payment struct {
    ID            int     `json:"id"`
    UserID        int     `json:"user_id"`
    OrderID       int     `json:"order_id"`
    Amount        float64 `json:"amount"`
    PaymentMethod string  `json:"payment_method"`
    TransactionID string  `json:"transaction_id"`
    Status        string  `json:"status"`
}