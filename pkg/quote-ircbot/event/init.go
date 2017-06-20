// Package event contains functions relating to IRC numeric events.
package event

import (
	"github.com/n7st/quoteDB/pkg/quote-ircbot"

	"github.com/thoj/go-ircevent"
)

type EventFnProvider struct {
	qb *quote_ircbot.QuoteBot
}

func Initialise(qb *quote_ircbot.QuoteBot) {
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
