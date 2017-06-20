// Package event contains functions relating to IRC numeric events.
package event

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/thoj/go-ircevent"
)

// callback001() runs when the bot connects to the IRC network.
func (q *EventFnProvider) callback001(e *irc.Event) {
	if q.qb.Recover && q.qb.Config.Password != "" {
		log.Println("Attempting to recover nickname")

		q.qb.IRC.Privmsgf("nickserv", "recover %s %s",
			q.qb.Config.Nickname, q.qb.Config.Password)

		q.qb.Recover = false
	}

	if q.qb.Config.Password != "" {
		q.qb.IRC.Privmsgf("nickserv", "identify %s", q.qb.Config.Password)
	} else {
		q.qb.JoinChannels(q.qb.Config.Channels)
	}

	q.qb.IRC.Mode(q.qb.IRC.GetNick(), q.qb.Config.Modes)
}

// callback433() runs when the bot encounters a ghost (or someone using its
// nickname).
func (q *EventFnProvider) callback433(e *irc.Event) {
	time.Sleep(2 * time.Second)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n\nNickname in use - attempt recovery? (Y/n)\n> ")

	text, err := reader.ReadString('\n')

	if err != nil {
		log.Println(err)
		return
	}

	text = strings.TrimSpace(text)

	if text == "" {
		text = "y"
	}

	if strings.ToLower(text) == "y" {
		// Recovery will be run on event 001
		log.Println("Recovery will be attempted when the bot finishes connecting")
		q.qb.Recover = true
	}
}

// callback900() runs when the bot receives a confirmation from nickserv that
// it has identified. There is another JoinChannels() here so it can join
// private channels.
func (q *EventFnProvider) callback900(e *irc.Event) {
	q.qb.JoinChannels(q.qb.Config.Channels)
}
