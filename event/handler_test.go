package event

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func runHandlerTestSuite(t *testing.T, eventGenerator func() Eventer, h Handler) {
	t.Run("Send", func(t *testing.T) {
		runTestSendHandler(t, eventGenerator, h)
	})
	t.Run("Register", func(t *testing.T) {
		runTestRegisterHandler(t, eventGenerator, h)
	})
}

func runTestSendHandler(t *testing.T, eventGenerator func() Eventer, h Handler) {
	t.Run("simple send", func(t *testing.T) {
		require.NoError(t, h.Send(eventGenerator()))
	})
	t.Run("send wrong event type", func(t *testing.T) {
		require.Error(t, h.Send(NewEvent("wrong", "", "", nil)))
	})
}


func runTestRegisterHandler(t *testing.T, eventGenerator func() Eventer, h Handler) {
	t.Run("simple register and no event emitted", func(t *testing.T) {
		r := h.Register()
		require.NotNil(t, r)
		events := r.Received()
		require.Equal(t, 0, len(events))
	})
	t.Run("simple register and event emitted", func(t *testing.T) {
		r := h.Register()
		require.NotNil(t, r)
		e := eventGenerator()
		h.Send(e)
		events := r.Received()
		require.Equal(t, 1, len(events))
		require.Equal(t, e.ID(), events[0].ID())
	})
	t.Run("simple register and event emitted multiple read", func(t *testing.T) {
		r := h.Register()
		require.NotNil(t, r)
		e := eventGenerator()
		_ = h.Send(e)
		r.Received()
		events := r.Received()
		require.Equal(t, 0, len(events))
	})
	t.Run("multiple register and event emitted multiple read", func(t *testing.T) {
		r1 := h.Register()
		r2 := h.Register()
		e := eventGenerator()
		_ = h.Send(e)
		events := r1.Received()
		require.Equal(t, 1, len(events))
		_ = h.Send(e)
		require.Equal(t, 2, len(r2.Received()))
		require.Equal(t, 1, len(r1.Received()))
	})
	t.Run("simple register and unregister", func(t *testing.T) {
		r := h.Register()
		require.NotNil(t, r)
		h.Unregister(r)
		e := eventGenerator()
		_ = h.Send(e)
		events := r.Received()
		require.Equal(t, 0, len(events))
	})
	t.Run("simple register and consume on unregister", func(t *testing.T) {
		r := h.Register()
		require.NotNil(t, r)
		e := eventGenerator()
		_ = h.Send(e)
		h.Unregister(r)
		events := r.Received()
		require.Equal(t, 0, len(events))
	})
}
