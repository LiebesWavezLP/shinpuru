package listeners

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/sarulabs/di/v2"
	"github.com/zekroTJA/shinpuru/internal/middleware"
	"github.com/zekroTJA/shinpuru/internal/services/database"
	"github.com/zekroTJA/shinpuru/internal/util"
	"github.com/zekroTJA/shinpuru/internal/util/static"

	"github.com/bwmarrin/discordgo"
)

var (
	rxInvLink = regexp.MustCompile(`(?i)(?:https?:\/\/)?(?:www\.)?(?:discord\.gg|discord(?:app)?\.com\/invite)\/(.*)`)
	rxGenLink = regexp.MustCompile(`(?i)(https?:\/\/)?(www\.)?([\w-\S]+\.)+\w{1,10}\/?[\S]+`)
)

type ListenerInviteBlock struct {
	db  database.Database
	pmw *middleware.PermissionsMiddleware
}

func NewListenerInviteBlock(container di.Container) *ListenerInviteBlock {
	return &ListenerInviteBlock{
		db:  container.Get(static.DiDatabase).(database.Database),
		pmw: container.Get(static.DiPermissionMiddleware).(*middleware.PermissionsMiddleware),
	}
}

func (l *ListenerInviteBlock) HandlerMessageSend(s *discordgo.Session, e *discordgo.MessageCreate) {
	l.invokeCheck(s, e.Message)
}

func (l *ListenerInviteBlock) HandlerMessageEdit(s *discordgo.Session, e *discordgo.MessageUpdate) {
	l.invokeCheck(s, e.Message)
}

func (l *ListenerInviteBlock) invokeCheck(s *discordgo.Session, msg *discordgo.Message) {
	cont := msg.Content

	ok, matches := l.checkForInviteLink(cont)
	if ok {
		l.detected(s, msg, matches)
		return
	}

	link := rxGenLink.FindString(cont)
	if link != "" {
		ok, matches, err := l.followLink(link)
		if err != nil {
			util.Log.Error("Failed following link: ", err)
			return
		}
		if ok {
			l.detected(s, msg, matches)
		}
	}
}

func (l *ListenerInviteBlock) checkForInviteLink(cont string) (bool, [][]string) {
	matches := rxInvLink.FindAllStringSubmatch(cont, -1)
	return matches != nil, matches
}

func (l *ListenerInviteBlock) followLink(link string) (bool, [][]string, error) {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	}

	resp, err := http.DefaultClient.Get(link)
	if err != nil {
		return false, nil, nil
	}

	ok, matches := l.checkForInviteLink(resp.Request.URL.String())
	return ok, matches, nil
}

func (l *ListenerInviteBlock) detected(s *discordgo.Session, e *discordgo.Message, matches [][]string) error {
	fmt.Println(matches)

	enabled, err := l.db.GetGuildInviteBlock(e.GuildID)
	if database.IsErrDatabaseNotFound(err) {
		return nil
	}
	if err != nil || enabled == "" {
		return err
	}

	ok, override, err := l.pmw.CheckPermissions(s, e.GuildID, e.Author.ID, "!sp.guild.mod.inviteblock.send")
	if err != nil || ok || override {
		return err
	}

	if invites, err := s.GuildInvites(e.GuildID); err == nil {
		inviteCode := matches[0][1]
		for _, inv := range invites {
			if inv.Code == inviteCode {
				return nil
			}
		}
	} else {
		util.Log.Error(err)
	}

	return s.ChannelMessageDelete(e.ChannelID, e.ID)
}
