package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReport msg type = AE.
type TradeCaptureReport struct {
	message.Message
}

//TradeCaptureReportBuilder builds TradeCaptureReport messages.
type TradeCaptureReportBuilder struct {
	message.MessageBuilder
}

//CreateTradeCaptureReportBuilder returns an initialized TradeCaptureReportBuilder with specified required fields.
func CreateTradeCaptureReportBuilder(
	lastqty field.LastQty,
	lastpx field.LastPx,
	nosides field.NoSides) TradeCaptureReportBuilder {
	var builder TradeCaptureReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(lastqty)
	builder.Body.Set(lastpx)
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportID() (field.TradeReportID, error) {
	var f field.TradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportTransType() (field.TradeReportTransType, error) {
	var f field.TradeReportTransType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportType() (field.TradeReportType, error) {
	var f field.TradeReportType
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeRequestID() (field.TradeRequestID, error) {
	var f field.TradeRequestID
	err := m.Body.Get(&f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdType() (field.TrdType, error) {
	var f field.TrdType
	err := m.Body.Get(&f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdSubType() (field.TrdSubType, error) {
	var f field.TrdSubType
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTrdType() (field.SecondaryTrdType, error) {
	var f field.SecondaryTrdType
	err := m.Body.Get(&f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TransferReason() (field.TransferReason, error) {
	var f field.TransferReason
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecType() (field.ExecType, error) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//TotNumTradeReports is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TotNumTradeReports() (field.TotNumTradeReports, error) {
	var f field.TotNumTradeReports
	err := m.Body.Get(&f)
	return f, err
}

//LastRptRequested is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastRptRequested() (field.LastRptRequested, error) {
	var f field.LastRptRequested
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnsolicitedIndicator() (field.UnsolicitedIndicator, error) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeReportRefID() (field.TradeReportRefID, error) {
	var f field.TradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportRefID() (field.SecondaryTradeReportRefID, error) {
	var f field.SecondaryTradeReportRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeReportID() (field.SecondaryTradeReportID, error) {
	var f field.SecondaryTradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLinkID() (field.TradeLinkID, error) {
	var f field.TradeLinkID
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdMatchID() (field.TrdMatchID, error) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryExecID() (field.SecondaryExecID, error) {
	var f field.SecondaryExecID
	err := m.Body.Get(&f)
	return f, err
}

//ExecRestatementReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExecRestatementReason() (field.ExecRestatementReason, error) {
	var f field.ExecRestatementReason
	err := m.Body.Get(&f)
	return f, err
}

//PreviouslyReported is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PreviouslyReported() (field.PreviouslyReported, error) {
	var f field.PreviouslyReported
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionID() (field.UnderlyingTradingSessionID, error) {
	var f field.UnderlyingTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingTradingSessionSubID() (field.UnderlyingTradingSessionSubID, error) {
	var f field.UnderlyingTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//LastQty is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastQty() (field.LastQty, error) {
	var f field.LastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastPx is a required field for TradeCaptureReport.
func (m TradeCaptureReport) LastPx() (field.LastPx, error) {
	var f field.LastPx
	err := m.Body.Get(&f)
	return f, err
}

//LastParPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastParPx() (field.LastParPx, error) {
	var f field.LastParPx
	err := m.Body.Get(&f)
	return f, err
}

//LastSpotRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastSpotRate() (field.LastSpotRate, error) {
	var f field.LastSpotRate
	err := m.Body.Get(&f)
	return f, err
}

//LastForwardPoints is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastForwardPoints() (field.LastForwardPoints, error) {
	var f field.LastForwardPoints
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastMkt() (field.LastMkt, error) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AvgPxIndicator() (field.AvgPxIndicator, error) {
	var f field.AvgPxIndicator
	err := m.Body.Get(&f)
	return f, err
}

//NoPosAmt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoPosAmt() (field.NoPosAmt, error) {
	var f field.NoPosAmt
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//TradeLegRefID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeLegRefID() (field.TradeLegRefID, error) {
	var f field.TradeLegRefID
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchStatus() (field.MatchStatus, error) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MatchType() (field.MatchType, error) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//NoSides is a required field for TradeCaptureReport.
func (m TradeCaptureReport) NoSides() (field.NoSides, error) {
	var f field.NoSides
	err := m.Body.Get(&f)
	return f, err
}

//CopyMsgIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CopyMsgIndicator() (field.CopyMsgIndicator, error) {
	var f field.CopyMsgIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PublishTrdIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) PublishTrdIndicator() (field.PublishTrdIndicator, error) {
	var f field.PublishTrdIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ShortSaleReason is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ShortSaleReason() (field.ShortSaleReason, error) {
	var f field.ShortSaleReason
	err := m.Body.Get(&f)
	return f, err
}

//TrdRptStatus is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TrdRptStatus() (field.TrdRptStatus, error) {
	var f field.TrdRptStatus
	err := m.Body.Get(&f)
	return f, err
}

//AsOfIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) AsOfIndicator() (field.AsOfIndicator, error) {
	var f field.AsOfIndicator
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//TierCode is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TierCode() (field.TierCode, error) {
	var f field.TierCode
	err := m.Body.Get(&f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MessageEventSource() (field.MessageEventSource, error) {
	var f field.MessageEventSource
	err := m.Body.Get(&f)
	return f, err
}

//LastUpdateTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastUpdateTime() (field.LastUpdateTime, error) {
	var f field.LastUpdateTime
	err := m.Body.Get(&f)
	return f, err
}

//RndPx is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RndPx() (field.RndPx, error) {
	var f field.RndPx
	err := m.Body.Get(&f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeID() (field.TradeID, error) {
	var f field.TradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryTradeID() (field.SecondaryTradeID, error) {
	var f field.SecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FirmTradeID() (field.FirmTradeID, error) {
	var f field.FirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SecondaryFirmTradeID() (field.SecondaryFirmTradeID, error) {
	var f field.SecondaryFirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CalculatedCcyLastQty() (field.CalculatedCcyLastQty, error) {
	var f field.CalculatedCcyLastQty
	err := m.Body.Get(&f)
	return f, err
}

//LastSwapPoints is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) LastSwapPoints() (field.LastSwapPoints, error) {
	var f field.LastSwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingSettlementDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) UnderlyingSettlementDate() (field.UnderlyingSettlementDate, error) {
	var f field.UnderlyingSettlementDate
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) GrossTradeAmt() (field.GrossTradeAmt, error) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoRootPartyIDs() (field.NoRootPartyIDs, error) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradeHandlingInstr() (field.TradeHandlingInstr, error) {
	var f field.TradeHandlingInstr
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeHandlingInstr() (field.OrigTradeHandlingInstr, error) {
	var f field.OrigTradeHandlingInstr
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeDate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeDate() (field.OrigTradeDate, error) {
	var f field.OrigTradeDate
	err := m.Body.Get(&f)
	return f, err
}

//OrigTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigTradeID() (field.OrigTradeID, error) {
	var f field.OrigTradeID
	err := m.Body.Get(&f)
	return f, err
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) OrigSecondaryTradeID() (field.OrigSecondaryTradeID, error) {
	var f field.OrigSecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//TZTransactTime is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TZTransactTime() (field.TZTransactTime, error) {
	var f field.TZTransactTime
	err := m.Body.Get(&f)
	return f, err
}

//ReportedPxDiff is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ReportedPxDiff() (field.ReportedPxDiff, error) {
	var f field.ReportedPxDiff
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//RejectText is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RejectText() (field.RejectText, error) {
	var f field.RejectText
	err := m.Body.Get(&f)
	return f, err
}

//FeeMultiplier is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) FeeMultiplier() (field.FeeMultiplier, error) {
	var f field.FeeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//Volatility is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) Volatility() (field.Volatility, error) {
	var f field.Volatility
	err := m.Body.Get(&f)
	return f, err
}

//DividendYield is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) DividendYield() (field.DividendYield, error) {
	var f field.DividendYield
	err := m.Body.Get(&f)
	return f, err
}

//RiskFreeRate is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) RiskFreeRate() (field.RiskFreeRate, error) {
	var f field.RiskFreeRate
	err := m.Body.Get(&f)
	return f, err
}

//CurrencyRatio is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) CurrencyRatio() (field.CurrencyRatio, error) {
	var f field.CurrencyRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRepIndicators is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) NoTrdRepIndicators() (field.NoTrdRepIndicators, error) {
	var f field.NoTrdRepIndicators
	err := m.Body.Get(&f)
	return f, err
}

//TradePublishIndicator is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) TradePublishIndicator() (field.TradePublishIndicator, error) {
	var f field.TradePublishIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//VenueType is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) VenueType() (field.VenueType, error) {
	var f field.VenueType
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarketSegmentID() (field.MarketSegmentID, error) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for TradeCaptureReport.
func (m TradeCaptureReport) MarketID() (field.MarketID, error) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}
