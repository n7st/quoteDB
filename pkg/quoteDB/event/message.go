// Package event contains functions relating to IRC numeric events.
package event

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB/helper"
	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB/model"

	"github.com/thoj/go-ircevent"
)

const (
	// Bot commands
	quoteLastCmd = "quotelast"
	addQuoteCmd  = "addquote"
	quoteHelpCmd = "quotehelp"
	quotePageCmd = "quotepage"
)

// callbackPrivmsg() runs when the bot receives a message. Every message is
// stored so it can be cycled through to build a quote. Command attempts are
// not logged into history.
func (q *EventFnProvider) callbackPrivmsg(e *irc.Event) {
	channel := e.Arguments[0]

	if q.isCommandAttempt(e.Message()) {
		q.handleCommand(e)
	} else if e.Nick != q.qb.Config.Nickname {
		// Log message information for reading back later
		q.qb.History[channel] = append(q.qb.History[channel], map[string]string{
			"nick":      e.Nick,
			"message":   e.Message(),
			"timestamp": time.Now().String(),
		})

		maxLen := q.qb.Config.MaxHistorySize
		currentLen := len(q.qb.History[channel])

		if currentLen > maxLen {
			// Trim messages past max length out of the slice
			newStart := currentLen - maxLen
			q.qb.History[channel] = q.qb.History[channel][newStart:]
		}
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

	switch command {
	case addQuoteCmd:
		if argsWithoutCmd != "" {
			q.parseAddQuote(e, argsWithoutCmd)
		}
	case quoteLastCmd:
		if argsWithoutCmd != "" {
			q.parseQuoteLast(e, argsWithoutCmd)
		}
	case quoteHelpCmd:
		q.parseQuoteHelp(e)
	case quotePageCmd:
		q.parseQuotePage(e)
	}
}

func (q *EventFnProvider) parseQuoteLast(e *irc.Event, argsWithoutCmd string) {
	histLen, err := strconv.Atoi(argsWithoutCmd)

	if err == nil && histLen > 0 {
		channel := e.Arguments[0]
		lines := helper.LastNLinesFromHistory(q.qb.History[channel], histLen)

		q.addQuoteLines(channel, lines)
	} else {
		q.qb.IRC.Privmsgf(e.Arguments[0], "%s is not a positive integer", argsWithoutCmd)
	}
}

// parseQuoteHelp() is run by the "quotehelp" command and displays commands
// available to users of the bot.
func (q *EventFnProvider) parseQuoteHelp(e *irc.Event) {
	q.qb.IRC.Privmsgf(e.Arguments[0], "Commands: %saddquote, %squotepage",
		q.qb.Config.Trigger, q.qb.Config.Trigger)
}

// parseQuotePage() runs the "quotepage" command which displays the URL for the
// web output page.
func (q *EventFnProvider) parseQuotePage(e *irc.Event) {
	channel := url.PathEscape(e.Arguments[0])
	loc := fmt.Sprintf("%schannel/%s", q.qb.Config.BaseURL, channel)

	q.qb.IRC.Privmsgf(e.Arguments[0], "Quotes for this channel can be found at %s", loc)
}

// parseAddQuote() handles the "addquote" command.
func (q *EventFnProvider) parseAddQuote(e *irc.Event, argsWithoutCmd string) {
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

	q.addQuoteLines(channel, lines)
}

// addQuoteLines() runs a database transaction which adds a quote against a
// channel in the database.
// TODO: this might be better as part of a "repository" package.
func (q *EventFnProvider) addQuoteLines(channel string, lines []map[string]string) {
	if len(lines) > q.qb.Config.MaxQuoteSize {
		q.qb.IRC.Privmsgf(channel, "That quote is too long (max %d lines)", q.qb.Config.MaxQuoteSize)
		return
	}

	if len(lines) != 0 {
		var createErrors []error

		tx := q.qb.DB.Begin()

		mChannel := model.Channel{Name: channel}
		tx.Find(&mChannel)

		if mChannel.ID == 0 {
			if mChannelErr := tx.Create(&mChannel).Error; mChannelErr != nil {
				q.qb.IRC.Privmsg(channel, "An error occurred creating the quote")
				log.Println(mChannelErr)
				return
			}
		}

		head := model.Head{Channel: mChannel}
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
	} else {
		q.qb.IRC.Privmsgf(channel, "No lines were collected (maximum size: %d)",
			q.qb.Config.MaxQuoteSize)
	}
}
