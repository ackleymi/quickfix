package fix50

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
