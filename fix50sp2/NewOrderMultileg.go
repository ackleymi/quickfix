package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderMultileg msg type = AB.
type NewOrderMultileg struct {
	message.Message
}

//NewOrderMultilegBuilder builds NewOrderMultileg messages.
type NewOrderMultilegBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderMultilegBuilder returns an initialized NewOrderMultilegBuilder with specified required fields.
func CreateNewOrderMultilegBuilder(
	clordid field.ClOrdID,
	side field.Side,
	nolegs field.NoLegs,
	transacttime field.TransactTime,
	ordtype field.OrdType) NewOrderMultilegBuilder {
	var builder NewOrderMultilegBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderMultileg.
func (m NewOrderMultileg) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ClOrdLinkID() (field.ClOrdLinkID, error) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TradeOriginationDate() (field.TradeOriginationDate, error) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//DayBookingInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DayBookingInst() (field.DayBookingInst, error) {
	var f field.DayBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//BookingUnit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) BookingUnit() (field.BookingUnit, error) {
	var f field.BookingUnit
	err := m.Body.Get(&f)
	return f, err
}

//PreallocMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PreallocMethod() (field.PreallocMethod, error) {
	var f field.PreallocMethod
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//CashMargin is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CashMargin() (field.CashMargin, error) {
	var f field.CashMargin
	err := m.Body.Get(&f)
	return f, err
}

//ClearingFeeIndicator is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ClearingFeeIndicator() (field.ClearingFeeIndicator, error) {
	var f field.ClearingFeeIndicator
	err := m.Body.Get(&f)
	return f, err
}

//HandlInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) HandlInst() (field.HandlInst, error) {
	var f field.HandlInst
	err := m.Body.Get(&f)
	return f, err
}

//ExecInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExecInst() (field.ExecInst, error) {
	var f field.ExecInst
	err := m.Body.Get(&f)
	return f, err
}

//MinQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MinQty() (field.MinQty, error) {
	var f field.MinQty
	err := m.Body.Get(&f)
	return f, err
}

//MaxFloor is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxFloor() (field.MaxFloor, error) {
	var f field.MaxFloor
	err := m.Body.Get(&f)
	return f, err
}

//ExDestination is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExDestination() (field.ExDestination, error) {
	var f field.ExDestination
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for NewOrderMultileg.
func (m NewOrderMultileg) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//PrevClosePx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PrevClosePx() (field.PrevClosePx, error) {
	var f field.PrevClosePx
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a required field for NewOrderMultileg.
func (m NewOrderMultileg) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//LocateReqd is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) LocateReqd() (field.LocateReqd, error) {
	var f field.LocateReqd
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for NewOrderMultileg.
func (m NewOrderMultileg) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//OrderQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderQty() (field.OrderQty, error) {
	var f field.OrderQty
	err := m.Body.Get(&f)
	return f, err
}

//CashOrderQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CashOrderQty() (field.CashOrderQty, error) {
	var f field.CashOrderQty
	err := m.Body.Get(&f)
	return f, err
}

//OrderPercent is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderPercent() (field.OrderPercent, error) {
	var f field.OrderPercent
	err := m.Body.Get(&f)
	return f, err
}

//RoundingDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RoundingDirection() (field.RoundingDirection, error) {
	var f field.RoundingDirection
	err := m.Body.Get(&f)
	return f, err
}

//RoundingModulus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RoundingModulus() (field.RoundingModulus, error) {
	var f field.RoundingModulus
	err := m.Body.Get(&f)
	return f, err
}

//OrdType is a required field for NewOrderMultileg.
func (m NewOrderMultileg) OrdType() (field.OrdType, error) {
	var f field.OrdType
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//Price is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Price() (field.Price, error) {
	var f field.Price
	err := m.Body.Get(&f)
	return f, err
}

//StopPx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StopPx() (field.StopPx, error) {
	var f field.StopPx
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//ComplianceID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ComplianceID() (field.ComplianceID, error) {
	var f field.ComplianceID
	err := m.Body.Get(&f)
	return f, err
}

//SolicitedFlag is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SolicitedFlag() (field.SolicitedFlag, error) {
	var f field.SolicitedFlag
	err := m.Body.Get(&f)
	return f, err
}

//IOIID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) IOIID() (field.IOIID, error) {
	var f field.IOIID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) QuoteID() (field.QuoteID, error) {
	var f field.QuoteID
	err := m.Body.Get(&f)
	return f, err
}

//TimeInForce is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TimeInForce() (field.TimeInForce, error) {
	var f field.TimeInForce
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExpireDate() (field.ExpireDate, error) {
	var f field.ExpireDate
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//GTBookingInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) GTBookingInst() (field.GTBookingInst, error) {
	var f field.GTBookingInst
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CommCurrency() (field.CommCurrency, error) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FundRenewWaiv() (field.FundRenewWaiv, error) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}

//CustOrderCapacity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CustOrderCapacity() (field.CustOrderCapacity, error) {
	var f field.CustOrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//ForexReq is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ForexReq() (field.ForexReq, error) {
	var f field.ForexReq
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BookingType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) BookingType() (field.BookingType, error) {
	var f field.BookingType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PositionEffect() (field.PositionEffect, error) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//CoveredOrUncovered is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CoveredOrUncovered() (field.CoveredOrUncovered, error) {
	var f field.CoveredOrUncovered
	err := m.Body.Get(&f)
	return f, err
}

//MaxShow is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxShow() (field.MaxShow, error) {
	var f field.MaxShow
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetValue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegOffsetValue() (field.PegOffsetValue, error) {
	var f field.PegOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//PegMoveType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegMoveType() (field.PegMoveType, error) {
	var f field.PegMoveType
	err := m.Body.Get(&f)
	return f, err
}

//PegOffsetType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegOffsetType() (field.PegOffsetType, error) {
	var f field.PegOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//PegLimitType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegLimitType() (field.PegLimitType, error) {
	var f field.PegLimitType
	err := m.Body.Get(&f)
	return f, err
}

//PegRoundDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegRoundDirection() (field.PegRoundDirection, error) {
	var f field.PegRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//PegScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegScope() (field.PegScope, error) {
	var f field.PegScope
	err := m.Body.Get(&f)
	return f, err
}

//PegPriceType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegPriceType() (field.PegPriceType, error) {
	var f field.PegPriceType
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSecurityIDSource() (field.PegSecurityIDSource, error) {
	var f field.PegSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSecurityID() (field.PegSecurityID, error) {
	var f field.PegSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//PegSymbol is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSymbol() (field.PegSymbol, error) {
	var f field.PegSymbol
	err := m.Body.Get(&f)
	return f, err
}

//PegSecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSecurityDesc() (field.PegSecurityDesc, error) {
	var f field.PegSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionInst() (field.DiscretionInst, error) {
	var f field.DiscretionInst
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetValue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionOffsetValue() (field.DiscretionOffsetValue, error) {
	var f field.DiscretionOffsetValue
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionMoveType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionMoveType() (field.DiscretionMoveType, error) {
	var f field.DiscretionMoveType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionOffsetType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionOffsetType() (field.DiscretionOffsetType, error) {
	var f field.DiscretionOffsetType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionLimitType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionLimitType() (field.DiscretionLimitType, error) {
	var f field.DiscretionLimitType
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionRoundDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionRoundDirection() (field.DiscretionRoundDirection, error) {
	var f field.DiscretionRoundDirection
	err := m.Body.Get(&f)
	return f, err
}

//DiscretionScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionScope() (field.DiscretionScope, error) {
	var f field.DiscretionScope
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategy is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TargetStrategy() (field.TargetStrategy, error) {
	var f field.TargetStrategy
	err := m.Body.Get(&f)
	return f, err
}

//TargetStrategyParameters is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TargetStrategyParameters() (field.TargetStrategyParameters, error) {
	var f field.TargetStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//ParticipationRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ParticipationRate() (field.ParticipationRate, error) {
	var f field.ParticipationRate
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//Designation is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Designation() (field.Designation, error) {
	var f field.Designation
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegRptTypeReq is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MultiLegRptTypeReq() (field.MultiLegRptTypeReq, error) {
	var f field.MultiLegRptTypeReq
	err := m.Body.Get(&f)
	return f, err
}

//NoStrategyParameters is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoStrategyParameters() (field.NoStrategyParameters, error) {
	var f field.NoStrategyParameters
	err := m.Body.Get(&f)
	return f, err
}

//SwapPoints is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SwapPoints() (field.SwapPoints, error) {
	var f field.SwapPoints
	err := m.Body.Get(&f)
	return f, err
}

//MatchIncrement is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MatchIncrement() (field.MatchIncrement, error) {
	var f field.MatchIncrement
	err := m.Body.Get(&f)
	return f, err
}

//MaxPriceLevels is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxPriceLevels() (field.MaxPriceLevels, error) {
	var f field.MaxPriceLevels
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryDisplayQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecondaryDisplayQty() (field.SecondaryDisplayQty, error) {
	var f field.SecondaryDisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayWhen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayWhen() (field.DisplayWhen, error) {
	var f field.DisplayWhen
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayMethod() (field.DisplayMethod, error) {
	var f field.DisplayMethod
	err := m.Body.Get(&f)
	return f, err
}

//DisplayLowQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayLowQty() (field.DisplayLowQty, error) {
	var f field.DisplayLowQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayHighQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayHighQty() (field.DisplayHighQty, error) {
	var f field.DisplayHighQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayMinIncr is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayMinIncr() (field.DisplayMinIncr, error) {
	var f field.DisplayMinIncr
	err := m.Body.Get(&f)
	return f, err
}

//RefreshQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RefreshQty() (field.RefreshQty, error) {
	var f field.RefreshQty
	err := m.Body.Get(&f)
	return f, err
}

//DisplayQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayQty() (field.DisplayQty, error) {
	var f field.DisplayQty
	err := m.Body.Get(&f)
	return f, err
}

//PriceProtectionScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceProtectionScope() (field.PriceProtectionScope, error) {
	var f field.PriceProtectionScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerType() (field.TriggerType, error) {
	var f field.TriggerType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerAction is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerAction() (field.TriggerAction, error) {
	var f field.TriggerAction
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPrice() (field.TriggerPrice, error) {
	var f field.TriggerPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSymbol is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSymbol() (field.TriggerSymbol, error) {
	var f field.TriggerSymbol
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSecurityID() (field.TriggerSecurityID, error) {
	var f field.TriggerSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSecurityIDSource() (field.TriggerSecurityIDSource, error) {
	var f field.TriggerSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//TriggerSecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSecurityDesc() (field.TriggerSecurityDesc, error) {
	var f field.TriggerSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPriceType() (field.TriggerPriceType, error) {
	var f field.TriggerPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceTypeScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPriceTypeScope() (field.TriggerPriceTypeScope, error) {
	var f field.TriggerPriceTypeScope
	err := m.Body.Get(&f)
	return f, err
}

//TriggerPriceDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPriceDirection() (field.TriggerPriceDirection, error) {
	var f field.TriggerPriceDirection
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewPrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerNewPrice() (field.TriggerNewPrice, error) {
	var f field.TriggerNewPrice
	err := m.Body.Get(&f)
	return f, err
}

//TriggerOrderType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerOrderType() (field.TriggerOrderType, error) {
	var f field.TriggerOrderType
	err := m.Body.Get(&f)
	return f, err
}

//TriggerNewQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerNewQty() (field.TriggerNewQty, error) {
	var f field.TriggerNewQty
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerTradingSessionID() (field.TriggerTradingSessionID, error) {
	var f field.TriggerTradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TriggerTradingSessionSubID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerTradingSessionSubID() (field.TriggerTradingSessionSubID, error) {
	var f field.TriggerTradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//RefOrderID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RefOrderID() (field.RefOrderID, error) {
	var f field.RefOrderID
	err := m.Body.Get(&f)
	return f, err
}

//RefOrderIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RefOrderIDSource() (field.RefOrderIDSource, error) {
	var f field.RefOrderIDSource
	err := m.Body.Get(&f)
	return f, err
}

//PreTradeAnonymity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PreTradeAnonymity() (field.PreTradeAnonymity, error) {
	var f field.PreTradeAnonymity
	err := m.Body.Get(&f)
	return f, err
}

//ExDestinationIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExDestinationIDSource() (field.ExDestinationIDSource, error) {
	var f field.ExDestinationIDSource
	err := m.Body.Get(&f)
	return f, err
}

//MultilegModel is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MultilegModel() (field.MultilegModel, error) {
	var f field.MultilegModel
	err := m.Body.Get(&f)
	return f, err
}

//MultilegPriceMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MultilegPriceMethod() (field.MultilegPriceMethod, error) {
	var f field.MultilegPriceMethod
	err := m.Body.Get(&f)
	return f, err
}

//RiskFreeRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RiskFreeRate() (field.RiskFreeRate, error) {
	var f field.RiskFreeRate
	err := m.Body.Get(&f)
	return f, err
}
