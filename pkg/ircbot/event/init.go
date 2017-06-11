package event

import (
	"github.com/thoj/go-ircevent"
	"github.com/n7st/quoteDB/pkg/ircbot"
)


type EventFnProvider struct {
	qb *ircbot.QuoteBot
}

func Initialise(qb *ircbot.QuoteBot) {
	provider := EventFnProvider{qb: qb}

	for name, fn := range ircEvents(provider) {
		qb.IRC.AddCallback(name, fn)
	}
}

func ircEvents(p EventFnProvider) map[string]func(e *irc.Event) {
	return map[string]func(e *irc.Event){
		// basic.go
		"001": p.callback001,
		"443": p.callback443,
		"900": p.callback900,

		// message.go
		"PRIVMSG": p.callbackPrivmsg,
	}
}
