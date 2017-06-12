package event

import "github.com/thoj/go-ircevent"

// connected
func (q *EventFnProvider) callback001(e *irc.Event) {
	if q.qb.Config.Password != "" {
		q.qb.IRC.Privmsgf("nickserv", "identify %s", q.qb.Config.Password)
	}

	q.qb.JoinChannels(q.qb.Config.Channels)

	q.qb.IRC.Mode(q.qb.IRC.GetNick(), q.qb.Config.Modes)
}

// nickname in use
func (q *EventFnProvider) callback443(e *irc.Event) {
}

// identified
func (q *EventFnProvider) callback900(e *irc.Event) {
}
