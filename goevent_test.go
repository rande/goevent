package goevent

import (
	"github.com/stretchr/testify/assert"
	"testing"
	// "log"
)

func Test_Event(t *testing.T) {

	e := NewEvent()

	e.Set("hello", "world")

	assert.Equal(t, e.Get("hello").(string), "world")
}

func Test_Dispatcher(t *testing.T) {
    d := NewEventDispatcher()

    d.Attach("app.load", func(e *Event) *Event {
        e.Set("hello", "world")
        e.StopPropagation()

        return e
    })

    e := d.Dispatch("app.load", NewEvent())

    assert.True(t, e.IsPropagationStopped())
    assert.Equal(t, e.Get("hello").(string), "world")
}
