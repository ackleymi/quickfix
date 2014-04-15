package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MarketDataRequest msg type = V.
type MarketDataRequest struct {
	message.Message
}

//MarketDataRequestBuilder builds MarketDataRequest messages.
type MarketDataRequestBuilder struct {
	message.MessageBuilder
}

//CreateMarketDataRequestBuilder returns an initialized MarketDataRequestBuilder with specified required fields.
func CreateMarketDataRequestBuilder(
	mdreqid field.MDReqID,
	subscriptionrequesttype field.SubscriptionRequestType,
	marketdepth field.MarketDepth,
	nomdentrytypes field.NoMDEntryTypes,
	norelatedsym field.NoRelatedSym) MarketDataRequestBuilder {
	var builder MarketDataRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(mdreqid)
	builder.Body.Set(subscriptionrequesttype)
	builder.Body.Set(marketdepth)
	builder.Body.Set(nomdentrytypes)
	builder.Body.Set(norelatedsym)
	return builder
}

//MDReqID is a required field for MarketDataRequest.
func (m MarketDataRequest) MDReqID() (field.MDReqID, error) {
	var f field.MDReqID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a required field for MarketDataRequest.
func (m MarketDataRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MarketDepth is a required field for MarketDataRequest.
func (m MarketDataRequest) MarketDepth() (field.MarketDepth, error) {
	var f field.MarketDepth
	err := m.Body.Get(&f)
	return f, err
}

//MDUpdateType is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDUpdateType() (field.MDUpdateType, error) {
	var f field.MDUpdateType
	err := m.Body.Get(&f)
	return f, err
}

//AggregatedBook is a non-required field for MarketDataRequest.
func (m MarketDataRequest) AggregatedBook() (field.AggregatedBook, error) {
	var f field.AggregatedBook
	err := m.Body.Get(&f)
	return f, err
}

//OpenCloseSettlFlag is a non-required field for MarketDataRequest.
func (m MarketDataRequest) OpenCloseSettlFlag() (field.OpenCloseSettlFlag, error) {
	var f field.OpenCloseSettlFlag
	err := m.Body.Get(&f)
	return f, err
}

//Scope is a non-required field for MarketDataRequest.
func (m MarketDataRequest) Scope() (field.Scope, error) {
	var f field.Scope
	err := m.Body.Get(&f)
	return f, err
}

//MDImplicitDelete is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDImplicitDelete() (field.MDImplicitDelete, error) {
	var f field.MDImplicitDelete
	err := m.Body.Get(&f)
	return f, err
}

//NoMDEntryTypes is a required field for MarketDataRequest.
func (m MarketDataRequest) NoMDEntryTypes() (field.NoMDEntryTypes, error) {
	var f field.NoMDEntryTypes
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a required field for MarketDataRequest.
func (m MarketDataRequest) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for MarketDataRequest.
func (m MarketDataRequest) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ApplQueueAction is a non-required field for MarketDataRequest.
func (m MarketDataRequest) ApplQueueAction() (field.ApplQueueAction, error) {
	var f field.ApplQueueAction
	err := m.Body.Get(&f)
	return f, err
}

//ApplQueueMax is a non-required field for MarketDataRequest.
func (m MarketDataRequest) ApplQueueMax() (field.ApplQueueMax, error) {
	var f field.ApplQueueMax
	err := m.Body.Get(&f)
	return f, err
}

//MDQuoteType is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDQuoteType() (field.MDQuoteType, error) {
	var f field.MDQuoteType
	err := m.Body.Get(&f)
	return f, err
}
