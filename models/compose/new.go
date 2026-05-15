package compose

import tea "charm.land/bubbletea/v2"

type Option func(*Model)

func New(opts ...Option) *Model {
	m := &Model{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func WithPrimary(model tea.Model) Option {
	return func(m *Model) { m.primary = model }
}

func WithSecondary(model tea.Model) Option {
	return func(m *Model) { m.secondary = model }
}

func WithOrientation(orientation Orientation) Option {
	return func(m *Model) { m.Orientation = orientation }
}

func WithReverseOrder(reverseOrder bool) Option {
	return func(m *Model) { m.ReverseOrder = reverseOrder }
}

func WithSizeResolver(sizeResolver SizeResolver) Option {
	return func(m *Model) { m.SizeResolver = sizeResolver }
}

func WithFocussed(focussed bool) Option {
	return func(m *Model) { m.focussed = focussed }
}

func WithInitialFocus(focus Focus) Option {
	return func(m *Model) { m.focus = focus }
}

func WithMsgRoute(msgRoute MsgRoute) Option {
	return func(m *Model) { m.MsgRoutes = append(m.MsgRoutes, msgRoute) }
}

func WithMsgTypeRoute[T any](destination Focus) Option {
	return WithMsgRoute(func(msg tea.Msg) (ok bool, focus Focus) {
		_, ok = msg.(T)
		focus = destination
		return
	})
}
