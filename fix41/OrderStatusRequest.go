package fix41

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderStatusRequest msg type = H.
type OrderStatusRequest struct {
	message.Message
}

//OrderStatusRequestBuilder builds OrderStatusRequest messages.
type OrderStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderStatusRequestBuilder returns an initialized OrderStatusRequestBuilder with specified required fields.
func CreateOrderStatusRequestBuilder(
	clordid field.ClOrdID,
	symbol field.Symbol,
	side field.Side) OrderStatusRequestBuilder {
	var builder OrderStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clordid)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	return builder
}

//OrderID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for OrderStatusRequest.
func (m OrderStatusRequest) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClientID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ClientID() (field.ClientID, error) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) ExecBroker() (field.ExecBroker, error) {
	var f field.ExecBroker
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for OrderStatusRequest.
func (m OrderStatusRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) IDSource() (field.IDSource, error) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDay is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) MaturityDay() (field.MaturityDay, error) {
	var f field.MaturityDay
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for OrderStatusRequest.
func (m OrderStatusRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for OrderStatusRequest.
func (m OrderStatusRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}
