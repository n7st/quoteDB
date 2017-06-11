package event

import (
	"regexp"
	"strings"
	"time"

	"github.com/thoj/go-ircevent"
	"github.com/n7st/quoteDB/model"
)

func (q *EventFnProvider) callbackPrivmsg(e *irc.Event) {
	channel := e.Arguments[0]

	// TODO strip bot commands from logged information?
	q.qb.History[channel] = append(q.qb.History[channel], map[string]string{
		"nick":      e.Nick,
		"message":   e.Message(),
		"timestamp": time.Now().String(),
	})

	if q.isCommandAttempt(e) {
		q.handleCommand(e)
	}
}

func (q *EventFnProvider) isCommandAttempt(e *irc.Event) bool {
	if strings.HasPrefix(e.Message(), q.qb.Config.Trigger) {
		return true
	}

	return false
}

func (q *EventFnProvider) handleCommand(e *irc.Event) {
	args := strings.Split(e.Message(), " ")

	command := strings.TrimPrefix(args[0], q.qb.Config.Trigger)
	argsWithoutCmd := strings.Join(args[1:], " ")

	if command == "addquote" {
		if argsWithoutCmd != "" {
			q.parseAddQuote(e, argsWithoutCmd)
		}
	}
}

// parseAddQuote() handles the "addquote" command and needs refactoring.
func (q EventFnProvider) parseAddQuote(e *irc.Event, argsWithoutCmd string) {
	var (
		lines   []map[string]string
		options []string
	)

	channel := e.Arguments[0]
	matched := false

	// TODO https://golang.org/pkg/flag/#FlagSet
	r := regexp.MustCompile("([^\"]*)")
	matches := r.FindAllStringSubmatch(argsWithoutCmd, -1)

	for i, v := range matches {
		// Skip non-options (i.e. strings between options).
		if i % 2 == 0 {
			continue
		}

		options = append(options, v[0])
	}

	if len(options) < 1 || options[0] == "" || (len(options) == 2 && options[1] == "") {
		q.qb.IRC.Privmsg(channel, `The 'addquote' command requires one or two arguments ("start string" [and "end string"])`)
		return
	}

	if len(options) == 1 {
		// make the start and end the same, just select one line
		options = append(options, options[0])
	}

	for _, line := range q.qb.History[channel] {
		if strings.HasPrefix(line["message"], options[0]) {
			matched = true
		}

		if matched {
			lines = append(lines, line)
		}

		if strings.HasPrefix(line["message"], options[1]) {
			matched = false
			continue // ?
		}
	}

	if len(lines) != 0 {
		head := model.Head{Channel: channel}
		q.qb.DB.Create(&head)

		for _, line := range lines {
			line := model.Line{
				Content: line["message"],
				Author:  line["nick"],
				Head:    head,
			}

			q.qb.DB.Create(&line)
		}

		q.qb.IRC.Privmsgf(channel, "Added quote contains %d lines (%s/%d)", len(lines), channel, head.ID)
	}
}
