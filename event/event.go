package event

type Event struct {
	Name string
	Data interface{}
}
type EventListener func(Event)

type EventDispacher struct {
	listener map[string][]EventListener
}

func (d *EventDispacher) AddEventListener(eventName string, listener EventListener) {
	d.listener[eventName] = append(d.listener[eventName], listener)
}
func (d *EventDispacher) Dispatch(event Event) {
	for _, listeners := range d.listener[event.Name] {
		listeners(event)
	}
}
