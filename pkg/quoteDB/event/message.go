// Package event contains functions relating to IRC numeric events.
package event

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/n7st/quoteDB/pkg/quoteDB/helper"
	"github.com/n7st/quoteDB/pkg/quoteDB/model"

	"github.com/thoj/go-ircevent"
)

// callbackPrivmsg() runs when the bot receives a message. Every message is
// stored so it can be cycled through to build a quote.
func (q *EventFnProvider) callbackPrivmsg(e *irc.Event) {
	channel := e.Arguments[0]

	if e.Nick != q.qb.Config.Nickname {
		// Log message information for reading back later
		q.qb.History[channel] = append(q.qb.History[channel], map[string]string{
			"nick":      e.Nick,
			"message":   e.Message(),
			"timestamp": time.Now().String(),
		})

		maxLen := q.qb.Config.MaxQuoteSize
		currentLen := len(q.qb.History[channel])

		if currentLen > maxLen {
			// Trim messages past max length out of the slice
			newStart := currentLen - maxLen
			q.qb.History[channel] = q.qb.History[channel][newStart:]
		}
	}

	if q.isCommandAttempt(e.Message()) {
		q.handleCommand(e)
	}
}

// isCommandAttempt() checks if a given message looks like a bot command.
func (q *EventFnProvider) isCommandAttempt(message string) bool {
	if strings.HasPrefix(message, q.qb.Config.Trigger) {
		return true
	}

	return false
}

// handleCommand() organises arguments to a command and calls a function to run
// it.
func (q *EventFnProvider) handleCommand(e *irc.Event) {
	args := strings.Split(e.Message(), " ")

	command := strings.TrimPrefix(args[0], q.qb.Config.Trigger)
	argsWithoutCmd := strings.Join(args[1:], " ")

	if command == "addquote" {
		if argsWithoutCmd != "" {
			q.parseAddQuote(e, argsWithoutCmd)
		}
	} else if command == "quotehelp" {
		q.parseQuoteHelp(e)
	} else if command == "quotepage" {
		q.parseQuotePage(e)
	}
}

// parseQuoteHelp() is run by the "quotehelp" command and displays commands
// available to users of the bot.
func (q EventFnProvider) parseQuoteHelp(e *irc.Event) {
	q.qb.IRC.Privmsgf(e.Arguments[0], "Commands: %saddquote, %squotepage",
		q.qb.Config.Trigger, q.qb.Config.Trigger)
}

// parseQuotePage() runs the "quotepage" command which displays the URL for the
// web output page.
func (q EventFnProvider) parseQuotePage(e *irc.Event) {
	channel := url.PathEscape(e.Arguments[0])
	loc := fmt.Sprintf("%schannel/%s", q.qb.Config.BaseURL, channel)

	q.qb.IRC.Privmsgf(e.Arguments[0], "Quotes for this channel can be found at %s", loc)
}

// parseAddQuote() handles the "addquote" command and needs refactoring.
func (q EventFnProvider) parseAddQuote(e *irc.Event, argsWithoutCmd string) {
	channel := e.Arguments[0]
	options := helper.OptionsFromString(argsWithoutCmd)

	if len(options) < 1 || options[0] == "" || (len(options) == 2 && options[1] == "") {
		q.qb.IRC.Privmsg(channel, `The 'addquote' command requires one or two arguments ("start string" [and "end string"])`)
		return
	}

	if len(options) == 1 {
		// make the start and end the same, just select one line
		options = append(options, options[0])
	}

	lines := helper.LinesFromHistory(q.qb.History[channel], options)

	if len(lines) > q.qb.Config.MaxQuoteSize {
		q.qb.IRC.Privmsgf(channel, "That quote is too long (max %d lines)", q.qb.Config.MaxQuoteSize)
		return
	}

	if len(lines) != 0 {
		var createErrors []error

		tx := q.qb.DB.Begin()

		head := model.Head{Channel: channel}
		if headErr := tx.Create(&head).Error; headErr != nil {
			q.qb.IRC.Privmsg(channel, "An error occurred creating the quote")
			log.Println(headErr)
			return
		}

		for _, line := range lines {
			line := model.Line{
				Content: line["message"],
				Author:  line["nick"],
				Head:    head,
			}

			if err := tx.Create(&line).Error; err != nil {
				log.Println(err)

				createErrors = append(createErrors, err)
			}
		}

		if len(createErrors) == 0 {
			tx.Commit()
			q.qb.IRC.Privmsgf(channel, "Quote added - %sview/%d",
				q.qb.Config.BaseURL, head.ID)
		} else {
			tx.Rollback()
			q.qb.IRC.Privmsg(channel, "An error occurred creating the quote")
		}
	}
}
