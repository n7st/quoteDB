// Package event contains functions relating to IRC numeric events.
package event

import "github.com/thoj/go-ircevent"

// callback001() runs when the bot connects to the IRC network.
func (q *EventFnProvider) callback001(e *irc.Event) {
	if q.qb.Config.Password != "" {
		q.qb.IRC.Privmsgf("nickserv", "identify %s", q.qb.Config.Password)
	} else {
		q.qb.JoinChannels(q.qb.Config.Channels)
	}

	q.qb.IRC.Mode(q.qb.IRC.GetNick(), q.qb.Config.Modes)
}

// callback443() runs when the bot encounters a ghost (or someone using its
// nickname).
func (q *EventFnProvider) callback443(e *irc.Event) {
	// TODO: Gracefully change nickname or attempt to recover and release
	// nickname.
}

// callback900() runs when the bot receives a confirmation from nickserv that
// it has identified. There is another JoinChannels() here so it can join
// private channels.
func (q *EventFnProvider) callback900(e *irc.Event) {
	q.qb.JoinChannels(q.qb.Config.Channels)
}
