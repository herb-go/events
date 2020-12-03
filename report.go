package events

//ReportBody event body
type ReportBody interface {
	//EventReportBody return event report body and any error if raised.
	EventReportBody() (map[string]string, error)
}

//Map map event report body struct
type Map map[string]string

//EventReportBody return event report body and any error if raised.
func (m Map) EventReportBody() (map[string]string, error) {
	return m, nil
}

//Report event report struct
type Report struct {
	//ID report iD
	ID string
	//Target event target
	Target string
	//Type event type
	Type string
	//Timestamp unix timestamp in second
	Timestamp int64
	//Body report body
	Body map[string]string
}

//NewReport create new report
func NewReport() *Report {
	return &Report{}
}

//CreateReport create reprot with given event,id and timestamp.
//Event data should be ReportBody
//Return report created and an error if raised
func CreateReport(e *Event, id string, timestamp int64) (*Report, error) {
	if e.Data == nil {
		return nil, ErrNotReportableEvent
	}
	b, ok := e.Data.(ReportBody)
	if !ok {
		return nil, ErrNotReportableEvent
	}
	body, err := b.EventReportBody()
	if err != nil {
		return nil, err
	}
	report := NewReport()
	report.Target = e.Target
	report.Type = string(e.Type)
	report.Body = body
	report.ID = id
	report.Timestamp = timestamp
	return report, nil
}
