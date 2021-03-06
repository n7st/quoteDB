// Package event contains functions relating to IRC numeric events.
package event

import (
	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB"

	"github.com/thoj/go-ircevent"
)

type EventFnProvider struct {
	qb *quoteDB.QuoteBot
}

func Initialise(qb *quoteDB.QuoteBot) {
	provider := EventFnProvider{qb: qb}

	for name, fn := range ircEvents(provider) {
		qb.IRC.AddCallback(name, fn)
	}
}

func ircEvents(p EventFnProvider) map[string]func(e *irc.Event) {
	return map[string]func(e *irc.Event){
		// basic.go
		"001": p.callback001,
		"433": p.callback433,
		"900": p.callback900,

		// message.go
		"PRIVMSG": p.callbackPrivmsg,
	}
}
