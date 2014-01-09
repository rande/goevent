package goevent

type Event struct {
	data               map[string]interface{}
	propagationStopped bool
}

type EventListenerFunc func(e *Event) *Event

func (e *Event) StopPropagation() {
	e.propagationStopped = true
}

func (e *Event) IsPropagationStopped() bool {
	return e.propagationStopped
}

func (e *Event) Has(name string) bool {
	if _, ok := e.data[name]; ok {
		return true
	}

	return false
}

func (e *Event) Get(name string) interface{} {
	if value, ok := e.data[name]; ok {
		return value
	}

	return nil
}

func (e *Event) Set(name string, value interface{}) *Event {
	e.data[name] = value

	return e
}

func NewEvent() *Event {
	e := Event{}
	e.data = make(map[string]interface{})

	return &e
}

func NewEventDispatcher() *EventDispatcher{
	d := EventDispatcher{}
	d.listeners = make(map[string][]EventListenerFunc)

	return &d
}

type EventDispatcher struct {
	listeners map[string][]EventListenerFunc
}

func (d *EventDispatcher) Dispatch(name string, e *Event) *Event {

	if _, ok := d.listeners[name]; !ok {
		return e
	}

	for pos := range d.listeners[name] {
		e = d.listeners[name][pos](e)

		if e.IsPropagationStopped() {
			return e
		}
	}

	return e
}

func (d *EventDispatcher) Attach(name string, f EventListenerFunc) *EventDispatcher {
	if _, ok := d.listeners[name]; !ok {
		d.listeners[name] = make([]EventListenerFunc, 0)
	}

	d.listeners[name] = append(d.listeners[name], f)

	return d
}
