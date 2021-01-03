package ui_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/egnd/go-tghandlers/pkg/ui"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
)

func Test_NewInlineButton(t *testing.T) {
	assert.EqualValues(t,
		tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("title", "data"),
			),
		),
		ui.NewInlineButton("title", "data"),
	)
}

func Test_NewInlineKeyboard(t *testing.T) {
	cases := []struct {
		name     string
		command  string
		curPage  int
		pageSize int
		lineSize int
		items    []ui.KeybItem
		keyb     tgbotapi.InlineKeyboardMarkup
		err      error
	}{
		{
			name:     "page first",
			command:  "testcmd",
			curPage:  1,
			pageSize: 4,
			lineSize: 2,
			items: []ui.KeybItem{
				{"lang 1", "l1"}, {"lang 2", "l2"}, {"lang 3", "l3"}, {"lang 4", "l4"},
				{},
			},
			keyb: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 1", "testcmd;l1"),
					tgbotapi.NewInlineKeyboardButtonData("lang 2", "testcmd;l2"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 3", "testcmd;l3"),
					tgbotapi.NewInlineKeyboardButtonData("lang 4", "testcmd;l4"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("→", fmt.Sprintf("testcmd-page;%d", 2)),
				),
			),
		},
		{
			name:     "page first 2",
			command:  "testcmd",
			curPage:  1,
			pageSize: 2,
			lineSize: 1,
			items: []ui.KeybItem{
				{"lang 1", "l1"},
				{"lang 3", "l3"},
				{},
			},
			keyb: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 1", "testcmd;l1"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 3", "testcmd;l3"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("→", fmt.Sprintf("testcmd-page;%d", 2)),
				),
			),
		},
		{
			name:     "page 2",
			command:  "testcmd",
			curPage:  2,
			pageSize: 4,
			lineSize: 2,
			items: []ui.KeybItem{
				{}, {}, {}, {},
				{"lang 1", "l1"}, {"lang 2", "l2"}, {"lang 3", "l3"}, {"lang 4", "l4"},
				{},
			},
			keyb: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 1", "testcmd;l1"),
					tgbotapi.NewInlineKeyboardButtonData("lang 2", "testcmd;l2"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 3", "testcmd;l3"),
					tgbotapi.NewInlineKeyboardButtonData("lang 4", "testcmd;l4"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("←", fmt.Sprintf("testcmd-page;%d", 1)),
					tgbotapi.NewInlineKeyboardButtonData("→", fmt.Sprintf("testcmd-page;%d", 3)),
				),
			),
		},
		{
			name:     "page last",
			command:  "testcmd",
			curPage:  2,
			pageSize: 4,
			lineSize: 2,
			items: []ui.KeybItem{
				{}, {}, {}, {},
				{"lang 1", "l1"}, {"lang 2", "l2"}, {"lang 3", "l3"},
			},
			keyb: tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 1", "testcmd;l1"),
					tgbotapi.NewInlineKeyboardButtonData("lang 2", "testcmd;l2"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("lang 3", "testcmd;l3"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("←", fmt.Sprintf("testcmd-page;%d", 1)),
				),
			),
		},
		{
			name: "empty items",
			err:  errors.New("empty items"),
		},
		{
			name:     "empty items",
			pageSize: 1,
			items:    []ui.KeybItem{{}},
			err:      errors.New("invalid page number - 0"),
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.name, func(tt *testing.T) {
			keyb, err := ui.NewInlineKeyboard(tcase.command, tcase.curPage, tcase.pageSize, tcase.lineSize, tcase.items)
			assert.EqualValues(tt, tcase.err, err)
			if err == nil {
				assert.EqualValues(tt, tcase.keyb, keyb)
			}
		})
	}
}
