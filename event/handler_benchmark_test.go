package event

import "testing"

func runHandlerBenchmarchSuite(b *testing.B, eventGenerator func() Eventer, h Handler) {
	b.Run("send lot of unread event", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = h.Send(eventGenerator())
		}
	})

	b.Run("register a lot of register", func(b *testing.B) {
		var r []Receiver
		for i := 0; i < b.N; i++ {
			r = append(r, h.Register())
		}
		b.ResetTimer()
		_ = h.Send(eventGenerator())
		for _, re := range r {
			re.Received()
		}
	})
	b.Run("send one event to lot of listener", func(b *testing.B) {
		var r []Receiver
		for i := 0; i < b.N; i++ {
			r = append(r, h.Register())
		}
		b.ResetTimer()
		_ = h.Send(eventGenerator())
		for _, re := range r {
			re.Received()
		}
	})
	b.Run("send lot event then read lot of listener", func(b *testing.B) {
		var r []Receiver
		for i := 0; i < b.N; i++ {
			r = append(r, h.Register())
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = h.Send(eventGenerator())
		}
		for _, re := range r {
			re.Received()
		}
	})
}