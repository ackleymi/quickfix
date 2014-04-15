package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//DerivativeSecurityListRequest msg type = z.
type DerivativeSecurityListRequest struct {
	message.Message
}

//DerivativeSecurityListRequestBuilder builds DerivativeSecurityListRequest messages.
type DerivativeSecurityListRequestBuilder struct {
	message.MessageBuilder
}

//CreateDerivativeSecurityListRequestBuilder returns an initialized DerivativeSecurityListRequestBuilder with specified required fields.
func CreateDerivativeSecurityListRequestBuilder(
	securityreqid field.SecurityReqID,
	securitylistrequesttype field.SecurityListRequestType) DerivativeSecurityListRequestBuilder {
	var builder DerivativeSecurityListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securitylistrequesttype)
	return builder
}

//SecurityReqID is a required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListRequestType is a required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecurityListRequestType() (field.SecurityListRequestType, error) {
	var f field.SecurityListRequestType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbol is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSymbol() (field.UnderlyingSymbol, error) {
	var f field.UnderlyingSymbol
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSymbolSfx() (field.UnderlyingSymbolSfx, error) {
	var f field.UnderlyingSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityID() (field.UnderlyingSecurityID, error) {
	var f field.UnderlyingSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityIDSource() (field.UnderlyingSecurityIDSource, error) {
	var f field.UnderlyingSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUnderlyingSecurityAltID() (field.NoUnderlyingSecurityAltID, error) {
	var f field.NoUnderlyingSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingProduct is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingProduct() (field.UnderlyingProduct, error) {
	var f field.UnderlyingProduct
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCFICode is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCFICode() (field.UnderlyingCFICode, error) {
	var f field.UnderlyingCFICode
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityType() (field.UnderlyingSecurityType, error) {
	var f field.UnderlyingSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecuritySubType() (field.UnderlyingSecuritySubType, error) {
	var f field.UnderlyingSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityMonthYear() (field.UnderlyingMaturityMonthYear, error) {
	var f field.UnderlyingMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityDate() (field.UnderlyingMaturityDate, error) {
	var f field.UnderlyingMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponPaymentDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCouponPaymentDate() (field.UnderlyingCouponPaymentDate, error) {
	var f field.UnderlyingCouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingIssueDate() (field.UnderlyingIssueDate, error) {
	var f field.UnderlyingIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepoCollateralSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepoCollateralSecurityType() (field.UnderlyingRepoCollateralSecurityType, error) {
	var f field.UnderlyingRepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseTerm is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepurchaseTerm() (field.UnderlyingRepurchaseTerm, error) {
	var f field.UnderlyingRepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRepurchaseRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRepurchaseRate() (field.UnderlyingRepurchaseRate, error) {
	var f field.UnderlyingRepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFactor is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFactor() (field.UnderlyingFactor, error) {
	var f field.UnderlyingFactor
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCreditRating is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCreditRating() (field.UnderlyingCreditRating, error) {
	var f field.UnderlyingCreditRating
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingInstrRegistry() (field.UnderlyingInstrRegistry, error) {
	var f field.UnderlyingInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCountryOfIssue() (field.UnderlyingCountryOfIssue, error) {
	var f field.UnderlyingCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStateOrProvinceOfIssue() (field.UnderlyingStateOrProvinceOfIssue, error) {
	var f field.UnderlyingStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingLocaleOfIssue() (field.UnderlyingLocaleOfIssue, error) {
	var f field.UnderlyingLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRedemptionDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRedemptionDate() (field.UnderlyingRedemptionDate, error) {
	var f field.UnderlyingRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStrikePrice() (field.UnderlyingStrikePrice, error) {
	var f field.UnderlyingStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStrikeCurrency() (field.UnderlyingStrikeCurrency, error) {
	var f field.UnderlyingStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingOptAttribute() (field.UnderlyingOptAttribute, error) {
	var f field.UnderlyingOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingContractMultiplier() (field.UnderlyingContractMultiplier, error) {
	var f field.UnderlyingContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCouponRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCouponRate() (field.UnderlyingCouponRate, error) {
	var f field.UnderlyingCouponRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityExchange() (field.UnderlyingSecurityExchange, error) {
	var f field.UnderlyingSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingIssuer() (field.UnderlyingIssuer, error) {
	var f field.UnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingIssuerLen() (field.EncodedUnderlyingIssuerLen, error) {
	var f field.EncodedUnderlyingIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingIssuer() (field.EncodedUnderlyingIssuer, error) {
	var f field.EncodedUnderlyingIssuer
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSecurityDesc() (field.UnderlyingSecurityDesc, error) {
	var f field.UnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingSecurityDescLen() (field.EncodedUnderlyingSecurityDescLen, error) {
	var f field.EncodedUnderlyingSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedUnderlyingSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedUnderlyingSecurityDesc() (field.EncodedUnderlyingSecurityDesc, error) {
	var f field.EncodedUnderlyingSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPProgram is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCPProgram() (field.UnderlyingCPProgram, error) {
	var f field.UnderlyingCPProgram
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCPRegType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCPRegType() (field.UnderlyingCPRegType, error) {
	var f field.UnderlyingCPRegType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCurrency() (field.UnderlyingCurrency, error) {
	var f field.UnderlyingCurrency
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingQty() (field.UnderlyingQty, error) {
	var f field.UnderlyingQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPx() (field.UnderlyingPx, error) {
	var f field.UnderlyingPx
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDirtyPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingDirtyPrice() (field.UnderlyingDirtyPrice, error) {
	var f field.UnderlyingDirtyPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingEndPrice() (field.UnderlyingEndPrice, error) {
	var f field.UnderlyingEndPrice
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingStartValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingStartValue() (field.UnderlyingStartValue, error) {
	var f field.UnderlyingStartValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCurrentValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCurrentValue() (field.UnderlyingCurrentValue, error) {
	var f field.UnderlyingCurrentValue
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingEndValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingEndValue() (field.UnderlyingEndValue, error) {
	var f field.UnderlyingEndValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyingStips is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUnderlyingStips() (field.NoUnderlyingStips, error) {
	var f field.NoUnderlyingStips
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAllocationPercent is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingAllocationPercent() (field.UnderlyingAllocationPercent, error) {
	var f field.UnderlyingAllocationPercent
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSettlementType() (field.UnderlyingSettlementType, error) {
	var f field.UnderlyingSettlementType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCashAmount() (field.UnderlyingCashAmount, error) {
	var f field.UnderlyingCashAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCashType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCashType() (field.UnderlyingCashType, error) {
	var f field.UnderlyingCashType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingUnitOfMeasure() (field.UnderlyingUnitOfMeasure, error) {
	var f field.UnderlyingUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingTimeUnit() (field.UnderlyingTimeUnit, error) {
	var f field.UnderlyingTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingCapValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingCapValue() (field.UnderlyingCapValue, error) {
	var f field.UnderlyingCapValue
	err := m.Body.Get(&f)
	return f, err
}

//NoUndlyInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoUndlyInstrumentParties() (field.NoUndlyInstrumentParties, error) {
	var f field.NoUndlyInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSettlMethod() (field.UnderlyingSettlMethod, error) {
	var f field.UnderlyingSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAdjustedQuantity is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingAdjustedQuantity() (field.UnderlyingAdjustedQuantity, error) {
	var f field.UnderlyingAdjustedQuantity
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFXRate() (field.UnderlyingFXRate, error) {
	var f field.UnderlyingFXRate
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFXRateCalc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFXRateCalc() (field.UnderlyingFXRateCalc, error) {
	var f field.UnderlyingFXRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingMaturityTime() (field.UnderlyingMaturityTime, error) {
	var f field.UnderlyingMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPutOrCall() (field.UnderlyingPutOrCall, error) {
	var f field.UnderlyingPutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingExerciseStyle() (field.UnderlyingExerciseStyle, error) {
	var f field.UnderlyingExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingUnitOfMeasureQty() (field.UnderlyingUnitOfMeasureQty, error) {
	var f field.UnderlyingUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasure() (field.UnderlyingPriceUnitOfMeasure, error) {
	var f field.UnderlyingPriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingPriceUnitOfMeasureQty() (field.UnderlyingPriceUnitOfMeasureQty, error) {
	var f field.UnderlyingPriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingContractMultiplierUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingContractMultiplierUnit() (field.UnderlyingContractMultiplierUnit, error) {
	var f field.UnderlyingContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingFlowScheduleType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingFlowScheduleType() (field.UnderlyingFlowScheduleType, error) {
	var f field.UnderlyingFlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingRestructuringType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingRestructuringType() (field.UnderlyingRestructuringType, error) {
	var f field.UnderlyingRestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSeniority is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingSeniority() (field.UnderlyingSeniority, error) {
	var f field.UnderlyingSeniority
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingNotionalPercentageOutstanding is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingNotionalPercentageOutstanding() (field.UnderlyingNotionalPercentageOutstanding, error) {
	var f field.UnderlyingNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingOriginalNotionalPercentageOutstanding is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingOriginalNotionalPercentageOutstanding() (field.UnderlyingOriginalNotionalPercentageOutstanding, error) {
	var f field.UnderlyingOriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingAttachmentPoint is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingAttachmentPoint() (field.UnderlyingAttachmentPoint, error) {
	var f field.UnderlyingAttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingDetachmentPoint is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) UnderlyingDetachmentPoint() (field.UnderlyingDetachmentPoint, error) {
	var f field.UnderlyingDetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) MarketID() (field.MarketID, error) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) MarketSegmentID() (field.MarketSegmentID, error) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbol is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSymbol() (field.DerivativeSymbol, error) {
	var f field.DerivativeSymbol
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSymbolSfx is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSymbolSfx() (field.DerivativeSymbolSfx, error) {
	var f field.DerivativeSymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityID() (field.DerivativeSecurityID, error) {
	var f field.DerivativeSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityIDSource is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityIDSource() (field.DerivativeSecurityIDSource, error) {
	var f field.DerivativeSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeSecurityAltID is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeSecurityAltID() (field.NoDerivativeSecurityAltID, error) {
	var f field.NoDerivativeSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProduct is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeProduct() (field.DerivativeProduct, error) {
	var f field.DerivativeProduct
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeProductComplex is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeProductComplex() (field.DerivativeProductComplex, error) {
	var f field.DerivativeProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivFlexProductEligibilityIndicator() (field.DerivFlexProductEligibilityIndicator, error) {
	var f field.DerivFlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityGroup is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityGroup() (field.DerivativeSecurityGroup, error) {
	var f field.DerivativeSecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCFICode is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCFICode() (field.DerivativeCFICode, error) {
	var f field.DerivativeCFICode
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityType() (field.DerivativeSecurityType, error) {
	var f field.DerivativeSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecuritySubType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecuritySubType() (field.DerivativeSecuritySubType, error) {
	var f field.DerivativeSecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityMonthYear is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityMonthYear() (field.DerivativeMaturityMonthYear, error) {
	var f field.DerivativeMaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityDate() (field.DerivativeMaturityDate, error) {
	var f field.DerivativeMaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMaturityTime is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMaturityTime() (field.DerivativeMaturityTime, error) {
	var f field.DerivativeMaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettleOnOpenFlag is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSettleOnOpenFlag() (field.DerivativeSettleOnOpenFlag, error) {
	var f field.DerivativeSettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeInstrmtAssignmentMethod() (field.DerivativeInstrmtAssignmentMethod, error) {
	var f field.DerivativeInstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityStatus is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityStatus() (field.DerivativeSecurityStatus, error) {
	var f field.DerivativeSecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssueDate is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeIssueDate() (field.DerivativeIssueDate, error) {
	var f field.DerivativeIssueDate
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeInstrRegistry is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeInstrRegistry() (field.DerivativeInstrRegistry, error) {
	var f field.DerivativeInstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCountryOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCountryOfIssue() (field.DerivativeCountryOfIssue, error) {
	var f field.DerivativeCountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStateOrProvinceOfIssue() (field.DerivativeStateOrProvinceOfIssue, error) {
	var f field.DerivativeStateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikePrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikePrice() (field.DerivativeStrikePrice, error) {
	var f field.DerivativeStrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeLocaleOfIssue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeLocaleOfIssue() (field.DerivativeLocaleOfIssue, error) {
	var f field.DerivativeLocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeCurrency is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeCurrency() (field.DerivativeStrikeCurrency, error) {
	var f field.DerivativeStrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeMultiplier() (field.DerivativeStrikeMultiplier, error) {
	var f field.DerivativeStrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeStrikeValue is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeStrikeValue() (field.DerivativeStrikeValue, error) {
	var f field.DerivativeStrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptAttribute is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeOptAttribute() (field.DerivativeOptAttribute, error) {
	var f field.DerivativeOptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractMultiplier is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeContractMultiplier() (field.DerivativeContractMultiplier, error) {
	var f field.DerivativeContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrement is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMinPriceIncrement() (field.DerivativeMinPriceIncrement, error) {
	var f field.DerivativeMinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeMinPriceIncrementAmount() (field.DerivativeMinPriceIncrementAmount, error) {
	var f field.DerivativeMinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeUnitOfMeasure() (field.DerivativeUnitOfMeasure, error) {
	var f field.DerivativeUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeUnitOfMeasureQty() (field.DerivativeUnitOfMeasureQty, error) {
	var f field.DerivativeUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasure is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceUnitOfMeasure() (field.DerivativePriceUnitOfMeasure, error) {
	var f field.DerivativePriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceUnitOfMeasureQty() (field.DerivativePriceUnitOfMeasureQty, error) {
	var f field.DerivativePriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeExerciseStyle is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeExerciseStyle() (field.DerivativeExerciseStyle, error) {
	var f field.DerivativeExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeOptPayAmount is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeOptPayAmount() (field.DerivativeOptPayAmount, error) {
	var f field.DerivativeOptPayAmount
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeTimeUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeTimeUnit() (field.DerivativeTimeUnit, error) {
	var f field.DerivativeTimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityExchange is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityExchange() (field.DerivativeSecurityExchange, error) {
	var f field.DerivativeSecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePositionLimit() (field.DerivativePositionLimit, error) {
	var f field.DerivativePositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeNTPositionLimit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeNTPositionLimit() (field.DerivativeNTPositionLimit, error) {
	var f field.DerivativeNTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeIssuer() (field.DerivativeIssuer, error) {
	var f field.DerivativeIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuerLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedIssuerLen() (field.DerivativeEncodedIssuerLen, error) {
	var f field.DerivativeEncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedIssuer is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedIssuer() (field.DerivativeEncodedIssuer, error) {
	var f field.DerivativeEncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityDesc() (field.DerivativeSecurityDesc, error) {
	var f field.DerivativeSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedSecurityDescLen() (field.DerivativeEncodedSecurityDescLen, error) {
	var f field.DerivativeEncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeEncodedSecurityDesc is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeEncodedSecurityDesc() (field.DerivativeEncodedSecurityDesc, error) {
	var f field.DerivativeEncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractSettlMonth is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeContractSettlMonth() (field.DerivativeContractSettlMonth, error) {
	var f field.DerivativeContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeEvents is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeEvents() (field.NoDerivativeEvents, error) {
	var f field.NoDerivativeEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoDerivativeInstrumentParties is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) NoDerivativeInstrumentParties() (field.NoDerivativeInstrumentParties, error) {
	var f field.NoDerivativeInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSettlMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSettlMethod() (field.DerivativeSettlMethod, error) {
	var f field.DerivativeSettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePriceQuoteMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePriceQuoteMethod() (field.DerivativePriceQuoteMethod, error) {
	var f field.DerivativePriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeValuationMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeValuationMethod() (field.DerivativeValuationMethod, error) {
	var f field.DerivativeValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeListMethod is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeListMethod() (field.DerivativeListMethod, error) {
	var f field.DerivativeListMethod
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeCapPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeCapPrice() (field.DerivativeCapPrice, error) {
	var f field.DerivativeCapPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFloorPrice is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeFloorPrice() (field.DerivativeFloorPrice, error) {
	var f field.DerivativeFloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//DerivativePutOrCall is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativePutOrCall() (field.DerivativePutOrCall, error) {
	var f field.DerivativePutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLLen is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXMLLen() (field.DerivativeSecurityXMLLen, error) {
	var f field.DerivativeSecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXML is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXML() (field.DerivativeSecurityXML, error) {
	var f field.DerivativeSecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeSecurityXMLSchema is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeSecurityXMLSchema() (field.DerivativeSecurityXMLSchema, error) {
	var f field.DerivativeSecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeContractMultiplierUnit is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeContractMultiplierUnit() (field.DerivativeContractMultiplierUnit, error) {
	var f field.DerivativeContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//DerivativeFlowScheduleType is a non-required field for DerivativeSecurityListRequest.
func (m DerivativeSecurityListRequest) DerivativeFlowScheduleType() (field.DerivativeFlowScheduleType, error) {
	var f field.DerivativeFlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}
