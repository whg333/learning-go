package calendar

type Event struct {
	title string
	Date
}

func NewEvent(title string) *Event {
	e := Event{}
	e.SetTitle(title)
	return &e
}

func (e *Event) SetTitle(title string) {
	e.title = title
}
