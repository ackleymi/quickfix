package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityList msg type = y.
type SecurityList struct {
	message.Message
}

//SecurityListBuilder builds SecurityList messages.
type SecurityListBuilder struct {
	message.MessageBuilder
}

//CreateSecurityListBuilder returns an initialized SecurityListBuilder with specified required fields.
func CreateSecurityListBuilder() SecurityListBuilder {
	var builder SecurityListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//SecurityReqID is a non-required field for SecurityList.
func (m SecurityList) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a non-required field for SecurityList.
func (m SecurityList) SecurityResponseID() (field.SecurityResponseID, error) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a non-required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (field.SecurityRequestResult, error) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) TotNoRelatedSym() (field.TotNoRelatedSym, error) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for SecurityList.
func (m SecurityList) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//SecurityReportID is a non-required field for SecurityList.
func (m SecurityList) SecurityReportID() (field.SecurityReportID, error) {
	var f field.SecurityReportID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for SecurityList.
func (m SecurityList) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for SecurityList.
func (m SecurityList) MarketID() (field.MarketID, error) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityList.
func (m SecurityList) MarketSegmentID() (field.MarketSegmentID, error) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for SecurityList.
func (m SecurityList) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for SecurityList.
func (m SecurityList) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SecurityList.
func (m SecurityList) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for SecurityList.
func (m SecurityList) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListID is a non-required field for SecurityList.
func (m SecurityList) SecurityListID() (field.SecurityListID, error) {
	var f field.SecurityListID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListRefID is a non-required field for SecurityList.
func (m SecurityList) SecurityListRefID() (field.SecurityListRefID, error) {
	var f field.SecurityListRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListDesc is a non-required field for SecurityList.
func (m SecurityList) SecurityListDesc() (field.SecurityListDesc, error) {
	var f field.SecurityListDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityListDescLen is a non-required field for SecurityList.
func (m SecurityList) EncodedSecurityListDescLen() (field.EncodedSecurityListDescLen, error) {
	var f field.EncodedSecurityListDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityListDesc is a non-required field for SecurityList.
func (m SecurityList) EncodedSecurityListDesc() (field.EncodedSecurityListDesc, error) {
	var f field.EncodedSecurityListDesc
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListType is a non-required field for SecurityList.
func (m SecurityList) SecurityListType() (field.SecurityListType, error) {
	var f field.SecurityListType
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListTypeSource is a non-required field for SecurityList.
func (m SecurityList) SecurityListTypeSource() (field.SecurityListTypeSource, error) {
	var f field.SecurityListTypeSource
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for SecurityList.
func (m SecurityList) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}
