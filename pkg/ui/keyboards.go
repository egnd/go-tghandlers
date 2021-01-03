// Package ui ...
package ui

import (
	"errors"
	"fmt"
	"math"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func NewInlineButton(title string, data string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(title, data),
		),
	)
}

type KeybItem struct {
	Key   string
	Value string
}

// NewInlineKeyboard is ...
// @TODO: fix nolint.
func NewInlineKeyboard( //nolint:funlen
	command string,
	curPage int,
	pageSize int,
	lineSize int,
	items []KeybItem,
) (keyb tgbotapi.InlineKeyboardMarkup, err error) {
	itemsCnt := len(items)
	if itemsCnt == 0 {
		err = errors.New("empty items")

		return
	}

	var pagesCount int
	if pageSize > 0 {
		pagesCount = int(math.Ceil(float64(itemsCnt) / float64(pageSize)))
		if curPage < 1 || curPage > pagesCount {
			err = fmt.Errorf("invalid page number - %d", curPage)

			return
		}
	}

	var btns [][]tgbotapi.InlineKeyboardButton

	for i := 0; i < pageSize; i++ {
		if itemsCnt <= curPage*pageSize-pageSize+i {
			break
		}

		if i%lineSize == 0 {
			btns = append(btns, []tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData(
					items[curPage*pageSize-pageSize+i].Key,
					fmt.Sprintf("%s;%s", command, items[curPage*pageSize-pageSize+i].Value),
				),
			})
		} else {
			btns[int(math.Floor(float64(i)/float64(lineSize)))] = append(btns[int(math.Floor(float64(i)/float64(lineSize)))],
				tgbotapi.NewInlineKeyboardButtonData(
					items[curPage*pageSize-pageSize+i].Key,
					fmt.Sprintf("%s;%s", command, items[curPage*pageSize-pageSize+i].Value),
				),
			)
		}
	}

	if pageSize > 0 {
		switch curPage {
		case 1:
			btns = append(btns, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("→", fmt.Sprintf("%s-page;%d", command, curPage+1)),
			))
		case pagesCount:
			btns = append(btns, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("←", fmt.Sprintf("%s-page;%d", command, curPage-1)),
			))
		default:
			btns = append(btns, tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("←", fmt.Sprintf("%s-page;%d", command, curPage-1)),
				tgbotapi.NewInlineKeyboardButtonData("→", fmt.Sprintf("%s-page;%d", command, curPage+1)),
			))
		}
	}

	keyb = tgbotapi.NewInlineKeyboardMarkup(btns...)

	return keyb, err
}
