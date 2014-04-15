package fix40

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Advertisement msg type = 7.
type Advertisement struct {
	message.Message
}

//AdvertisementBuilder builds Advertisement messages.
type AdvertisementBuilder struct {
	message.MessageBuilder
}

//CreateAdvertisementBuilder returns an initialized AdvertisementBuilder with specified required fields.
func CreateAdvertisementBuilder(
	advid field.AdvId,
	advtranstype field.AdvTransType,
	symbol field.Symbol,
	advside field.AdvSide,
	shares field.Shares) AdvertisementBuilder {
	var builder AdvertisementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(advid)
	builder.Body.Set(advtranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(advside)
	builder.Body.Set(shares)
	return builder
}

//AdvId is a required field for Advertisement.
func (m Advertisement) AdvId() (field.AdvId, error) {
	var f field.AdvId
	err := m.Body.Get(&f)
	return f, err
}

//AdvTransType is a required field for Advertisement.
func (m Advertisement) AdvTransType() (field.AdvTransType, error) {
	var f field.AdvTransType
	err := m.Body.Get(&f)
	return f, err
}

//AdvRefID is a non-required field for Advertisement.
func (m Advertisement) AdvRefID() (field.AdvRefID, error) {
	var f field.AdvRefID
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a required field for Advertisement.
func (m Advertisement) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for Advertisement.
func (m Advertisement) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for Advertisement.
func (m Advertisement) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//IDSource is a non-required field for Advertisement.
func (m Advertisement) IDSource() (field.IDSource, error) {
	var f field.IDSource
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for Advertisement.
func (m Advertisement) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for Advertisement.
func (m Advertisement) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//AdvSide is a required field for Advertisement.
func (m Advertisement) AdvSide() (field.AdvSide, error) {
	var f field.AdvSide
	err := m.Body.Get(&f)
	return f, err
}

//Shares is a required field for Advertisement.
func (m Advertisement) Shares() (field.Shares, error) {
	var f field.Shares
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for Advertisement.
func (m Advertisement) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for Advertisement.
func (m Advertisement) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for Advertisement.
func (m Advertisement) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Advertisement.
func (m Advertisement) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
