package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityListUpdateReport msg type = BK.
type SecurityListUpdateReport struct {
	message.Message
}

//SecurityListUpdateReportBuilder builds SecurityListUpdateReport messages.
type SecurityListUpdateReportBuilder struct {
	message.MessageBuilder
}

//CreateSecurityListUpdateReportBuilder returns an initialized SecurityListUpdateReportBuilder with specified required fields.
func CreateSecurityListUpdateReportBuilder() SecurityListUpdateReportBuilder {
	var builder SecurityListUpdateReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//SecurityReportID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityReportID() (field.SecurityReportID, error) {
	var f field.SecurityReportID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityReqID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityResponseID() (field.SecurityResponseID, error) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityRequestResult() (field.SecurityRequestResult, error) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) TotNoRelatedSym() (field.TotNoRelatedSym, error) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityUpdateAction is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityUpdateAction() (field.SecurityUpdateAction, error) {
	var f field.SecurityUpdateAction
	err := m.Body.Get(&f)
	return f, err
}

//CorporateAction is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) CorporateAction() (field.CorporateAction, error) {
	var f field.CorporateAction
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) MarketID() (field.MarketID, error) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) MarketSegmentID() (field.MarketSegmentID, error) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListID() (field.SecurityListID, error) {
	var f field.SecurityListID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListRefID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListRefID() (field.SecurityListRefID, error) {
	var f field.SecurityListRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListDesc is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListDesc() (field.SecurityListDesc, error) {
	var f field.SecurityListDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityListDescLen is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) EncodedSecurityListDescLen() (field.EncodedSecurityListDescLen, error) {
	var f field.EncodedSecurityListDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityListDesc is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) EncodedSecurityListDesc() (field.EncodedSecurityListDesc, error) {
	var f field.EncodedSecurityListDesc
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListType is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListType() (field.SecurityListType, error) {
	var f field.SecurityListType
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListTypeSource is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListTypeSource() (field.SecurityListTypeSource, error) {
	var f field.SecurityListTypeSource
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}
