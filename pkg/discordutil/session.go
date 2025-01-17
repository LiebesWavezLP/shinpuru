package discordutil

import (
	"image"
	"io"
	"time"

	"github.com/bwmarrin/discordgo"
	. "github.com/bwmarrin/discordgo"
)

// WrapHandler takes a handler function which uses ISession for
// the Discord session instance and T as event playload and returns
// a valid handler function using *discordgo.Session for the session
// and T as event payload.
//
// This can be used for handler functions which are unit tested
// and therefore need to use ISession to pass in mocked session
// instances.
func WrapHandler[T any](f func(s ISession, e T)) func(s *discordgo.Session, h T) {
	return func(s *discordgo.Session, h T) {
		f(s, h)
	}
}

type ISession interface {
	AddHandler(handler interface{}) func()
	AddHandlerOnce(handler interface{}) func()
	Application(appID string) (st *Application, err error)
	Applications() (st []*Application, err error)
	ApplicationCreate(ap *Application) (st *Application, err error)
	ApplicationUpdate(appID string, ap *Application) (st *Application, err error)
	ApplicationDelete(appID string) (err error)
	ApplicationAssets(appID string) (ass []*Asset, err error)
	ApplicationBotCreate(appID string) (st *User, err error)
	Request(method, urlStr string, data interface{}) (response []byte, err error)
	RequestWithBucketID(method, urlStr string, data interface{}, bucketID string) (response []byte, err error)
	RequestWithLockedBucket(method, urlStr, contentType string, b []byte, bucket *Bucket, sequence int) (response []byte, err error)
	User(userID string) (st *User, err error)
	UserAvatar(userID string) (img image.Image, err error)
	UserAvatarDecode(u *User) (img image.Image, err error)
	UserUpdate(username, avatar string) (st *User, err error)
	UserConnections() (conn []*UserConnection, err error)
	UserChannelCreate(recipientID string) (st *Channel, err error)
	UserGuilds(limit int, beforeID, afterID string) (st []*UserGuild, err error)
	UserChannelPermissions(userID, channelID string) (apermissions int64, err error)
	Guild(guildID string) (st *Guild, err error)
	GuildWithCounts(guildID string) (st *Guild, err error)
	GuildPreview(guildID string) (st *GuildPreview, err error)
	GuildCreate(name string) (st *Guild, err error)
	GuildEdit(guildID string, g GuildParams) (st *Guild, err error)
	GuildDelete(guildID string) (st *Guild, err error)
	GuildLeave(guildID string) (err error)
	GuildBans(guildID string, limit int, beforeID, afterID string) (st []*GuildBan, err error)
	GuildBanCreate(guildID, userID string, days int) (err error)
	GuildBan(guildID, userID string) (st *GuildBan, err error)
	GuildBanCreateWithReason(guildID, userID, reason string, days int) (err error)
	GuildBanDelete(guildID, userID string) (err error)
	GuildMembers(guildID string, after string, limit int) (st []*Member, err error)
	GuildMembersSearch(guildID, query string, limit int) (st []*Member, err error)
	GuildMember(guildID, userID string) (st *Member, err error)
	GuildMemberAdd(accessToken, guildID, userID, nick string, roles []string, mute, deaf bool) (err error)
	GuildMemberDelete(guildID, userID string) (err error)
	GuildMemberDeleteWithReason(guildID, userID, reason string) (err error)
	GuildMemberEdit(guildID, userID string, roles []string) (err error)
	GuildMemberEditComplex(guildID, userID string, data GuildMemberParams) (st *Member, err error)
	GuildMemberMove(guildID string, userID string, channelID *string) (err error)
	GuildMemberNickname(guildID, userID, nickname string) (err error)
	GuildMemberMute(guildID string, userID string, mute bool) (err error)
	GuildMemberTimeout(guildID string, userID string, until *time.Time) (err error)
	GuildMemberDeafen(guildID string, userID string, deaf bool) (err error)
	GuildMemberRoleAdd(guildID, userID, roleID string) (err error)
	GuildMemberRoleRemove(guildID, userID, roleID string) (err error)
	GuildChannels(guildID string) (st []*Channel, err error)
	GuildChannelCreateComplex(guildID string, data GuildChannelCreateData) (st *Channel, err error)
	GuildChannelCreate(guildID, name string, ctype ChannelType) (st *Channel, err error)
	GuildChannelsReorder(guildID string, channels []*Channel) (err error)
	GuildInvites(guildID string) (st []*Invite, err error)
	GuildRoles(guildID string) (st []*Role, err error)
	GuildRoleCreate(guildID string) (st *Role, err error)
	GuildRoleEdit(guildID, roleID, name string, color int, hoist bool, perm int64, mention bool) (st *Role, err error)
	GuildRoleReorder(guildID string, roles []*Role) (st []*Role, err error)
	GuildRoleDelete(guildID, roleID string) (err error)
	GuildPruneCount(guildID string, days uint32) (count uint32, err error)
	GuildPrune(guildID string, days uint32) (count uint32, err error)
	GuildIntegrations(guildID string) (st []*Integration, err error)
	GuildIntegrationCreate(guildID, integrationType, integrationID string) (err error)
	GuildIntegrationEdit(guildID, integrationID string, expireBehavior, expireGracePeriod int, enableEmoticons bool) (err error)
	GuildIntegrationDelete(guildID, integrationID string) (err error)
	GuildIcon(guildID string) (img image.Image, err error)
	GuildSplash(guildID string) (img image.Image, err error)
	GuildEmbed(guildID string) (st *GuildEmbed, err error)
	GuildEmbedEdit(guildID string, enabled bool, channelID string) (err error)
	GuildAuditLog(guildID, userID, beforeID string, actionType, limit int) (st *GuildAuditLog, err error)
	GuildEmojis(guildID string) (emoji []*Emoji, err error)
	GuildEmoji(guildID, emojiID string) (emoji *Emoji, err error)
	GuildEmojiCreate(guildID, name, image string, roles []string) (emoji *Emoji, err error)
	GuildEmojiEdit(guildID, emojiID, name string, roles []string) (emoji *Emoji, err error)
	GuildEmojiDelete(guildID, emojiID string) (err error)
	GuildTemplate(templateCode string) (st *GuildTemplate, err error)
	GuildCreateWithTemplate(templateCode, name, icon string) (st *Guild, err error)
	GuildTemplates(guildID string) (st []*GuildTemplate, err error)
	GuildTemplateCreate(guildID, name, description string) (st *GuildTemplate)
	GuildTemplateSync(guildID, templateCode string) (err error)
	GuildTemplateEdit(guildID, templateCode, name, description string) (st *GuildTemplate, err error)
	GuildTemplateDelete(guildID, templateCode string) (err error)
	Channel(channelID string) (st *Channel, err error)
	ChannelEdit(channelID, name string) (*Channel, error)
	ChannelEditComplex(channelID string, data *ChannelEdit) (st *Channel, err error)
	ChannelDelete(channelID string) (st *Channel, err error)
	ChannelTyping(channelID string) (err error)
	ChannelMessages(channelID string, limit int, beforeID, afterID, aroundID string) (st []*Message, err error)
	ChannelMessage(channelID, messageID string) (st *Message, err error)
	ChannelMessageSend(channelID string, content string) (*Message, error)
	ChannelMessageSendComplex(channelID string, data *MessageSend) (st *Message, err error)
	ChannelMessageSendTTS(channelID string, content string) (*Message, error)
	ChannelMessageSendEmbed(channelID string, embed *MessageEmbed) (*Message, error)
	ChannelMessageSendEmbeds(channelID string, embeds []*MessageEmbed) (*Message, error)
	ChannelMessageSendReply(channelID string, content string, reference *MessageReference) (*Message, error)
	// ChannelMessageSendEmbedReply(channelID string, embed *MessageEmbed, reference *MessageReference) (*Message, error)
	// ChannelMessageSendEmbedsReply(channelID string, embeds []*MessageEmbed, reference *MessageReference) (*Message, error)
	ChannelMessageEdit(channelID, messageID, content string) (*Message, error)
	ChannelMessageEditComplex(m *MessageEdit) (st *Message, err error)
	ChannelMessageEditEmbed(channelID, messageID string, embed *MessageEmbed) (*Message, error)
	ChannelMessageEditEmbeds(channelID, messageID string, embeds []*MessageEmbed) (*Message, error)
	ChannelMessageDelete(channelID, messageID string) (err error)
	ChannelMessagesBulkDelete(channelID string, messages []string) (err error)
	ChannelMessagePin(channelID, messageID string) (err error)
	ChannelMessageUnpin(channelID, messageID string) (err error)
	ChannelMessagesPinned(channelID string) (st []*Message, err error)
	ChannelFileSend(channelID, name string, r io.Reader) (*Message, error)
	ChannelFileSendWithMessage(channelID, content string, name string, r io.Reader) (*Message, error)
	ChannelInvites(channelID string) (st []*Invite, err error)
	ChannelInviteCreate(channelID string, i Invite) (st *Invite, err error)
	ChannelPermissionSet(channelID, targetID string, targetType PermissionOverwriteType, allow, deny int64) (err error)
	ChannelPermissionDelete(channelID, targetID string) (err error)
	ChannelMessageCrosspost(channelID, messageID string) (st *Message, err error)
	ChannelNewsFollow(channelID, targetID string) (st *ChannelFollow, err error)
	Invite(inviteID string) (st *Invite, err error)
	InviteWithCounts(inviteID string) (st *Invite, err error)
	InviteComplex(inviteID, guildScheduledEventID string, withCounts, withExpiration bool) (st *Invite, err error)
	InviteDelete(inviteID string) (st *Invite, err error)
	InviteAccept(inviteID string) (st *Invite, err error)
	VoiceRegions() (st []*VoiceRegion, err error)
	Gateway() (gateway string, err error)
	GatewayBot() (st *GatewayBotResponse, err error)
	WebhookCreate(channelID, name, avatar string) (st *Webhook, err error)
	ChannelWebhooks(channelID string) (st []*Webhook, err error)
	GuildWebhooks(guildID string) (st []*Webhook, err error)
	Webhook(webhookID string) (st *Webhook, err error)
	WebhookWithToken(webhookID, token string) (st *Webhook, err error)
	WebhookEdit(webhookID, name, avatar, channelID string) (st *Role, err error)
	WebhookEditWithToken(webhookID, token, name, avatar string) (st *Role, err error)
	WebhookDelete(webhookID string) (err error)
	WebhookDeleteWithToken(webhookID, token string) (st *Webhook, err error)
	WebhookExecute(webhookID, token string, wait bool, data *WebhookParams) (st *Message, err error)
	WebhookThreadExecute(webhookID, token string, wait bool, threadID string, data *WebhookParams) (st *Message, err error)
	WebhookMessage(webhookID, token, messageID string) (message *Message, err error)
	WebhookMessageEdit(webhookID, token, messageID string, data *WebhookEdit) (st *Message, err error)
	WebhookMessageDelete(webhookID, token, messageID string) (err error)
	MessageReactionAdd(channelID, messageID, emojiID string) error
	MessageReactionRemove(channelID, messageID, emojiID, userID string) error
	MessageReactionsRemoveAll(channelID, messageID string) error
	MessageReactionsRemoveEmoji(channelID, messageID, emojiID string) error
	MessageReactions(channelID, messageID, emojiID string, limit int, beforeID, afterID string) (st []*User, err error)
	MessageThreadStartComplex(channelID, messageID string, data *ThreadStart) (ch *Channel, err error)
	MessageThreadStart(channelID, messageID string, name string, archiveDuration int) (ch *Channel, err error)
	ThreadStartComplex(channelID string, data *ThreadStart) (ch *Channel, err error)
	ThreadStart(channelID, name string, typ ChannelType, archiveDuration int) (ch *Channel, err error)
	ThreadJoin(id string) error
	ThreadLeave(id string) error
	ThreadMemberAdd(threadID, memberID string) error
	ThreadMemberRemove(threadID, memberID string) error
	ThreadMember(threadID, memberID string) (member *ThreadMember, err error)
	ThreadMembers(threadID string) (members []*ThreadMember, err error)
	ThreadsActive(channelID string) (threads *ThreadsList, err error)
	GuildThreadsActive(guildID string) (threads *ThreadsList, err error)
	ThreadsArchived(channelID string, before *time.Time, limit int) (threads *ThreadsList, err error)
	ThreadsPrivateArchived(channelID string, before *time.Time, limit int) (threads *ThreadsList, err error)
	ThreadsPrivateJoinedArchived(channelID string, before *time.Time, limit int) (threads *ThreadsList, err error)
	ApplicationCommandCreate(appID string, guildID string, cmd *ApplicationCommand) (ccmd *ApplicationCommand, err error)
	ApplicationCommandEdit(appID, guildID, cmdID string, cmd *ApplicationCommand) (updated *ApplicationCommand, err error)
	ApplicationCommandBulkOverwrite(appID string, guildID string, commands []*ApplicationCommand) (createdCommands []*ApplicationCommand, err error)
	ApplicationCommandDelete(appID, guildID, cmdID string) error
	ApplicationCommand(appID, guildID, cmdID string) (cmd *ApplicationCommand, err error)
	ApplicationCommands(appID, guildID string) (cmd []*ApplicationCommand, err error)
	GuildApplicationCommandsPermissions(appID, guildID string) (permissions []*GuildApplicationCommandPermissions, err error)
	ApplicationCommandPermissions(appID, guildID, cmdID string) (permissions *GuildApplicationCommandPermissions, err error)
	ApplicationCommandPermissionsEdit(appID, guildID, cmdID string, permissions *ApplicationCommandPermissionsList) (err error)
	ApplicationCommandPermissionsBatchEdit(appID, guildID string, permissions []*GuildApplicationCommandPermissions) (err error)
	InteractionRespond(interaction *Interaction, resp *InteractionResponse) error
	InteractionResponse(interaction *Interaction) (*Message, error)
	InteractionResponseEdit(interaction *Interaction, newresp *WebhookEdit) (*Message, error)
	InteractionResponseDelete(interaction *Interaction) error
	FollowupMessageCreate(interaction *Interaction, wait bool, data *WebhookParams) (*Message, error)
	FollowupMessageEdit(interaction *Interaction, messageID string, data *WebhookEdit) (*Message, error)
	FollowupMessageDelete(interaction *Interaction, messageID string) error
	StageInstanceCreate(data *StageInstanceParams) (si *StageInstance, err error)
	StageInstance(channelID string) (si *StageInstance, err error)
	StageInstanceEdit(channelID string, data *StageInstanceParams) (si *StageInstance, err error)
	StageInstanceDelete(channelID string) (err error)
	GuildScheduledEvents(guildID string, userCount bool) (st []*GuildScheduledEvent, err error)
	GuildScheduledEvent(guildID, eventID string, userCount bool) (st *GuildScheduledEvent, err error)
	GuildScheduledEventCreate(guildID string, event *GuildScheduledEventParams) (st *GuildScheduledEvent, err error)
	GuildScheduledEventEdit(guildID, eventID string, event *GuildScheduledEventParams) (st *GuildScheduledEvent, err error)
	GuildScheduledEventDelete(guildID, eventID string) (err error)
	GuildScheduledEventUsers(guildID, eventID string, limit int, withMember bool, beforeID, afterID string) (st []*GuildScheduledEventUser, err error)
	// AutoModerationRules(guildID string) (st []*AutoModerationRule, err error)
	// AutoModerationRule(guildID, ruleID string) (st *AutoModerationRule, err error)
	// AutoModerationRuleCreate(guildID string, rule *AutoModerationRule) (st *AutoModerationRule, err error)
	// AutoModerationRuleEdit(guildID, ruleID string, rule *AutoModerationRule) (st *AutoModerationRule, err error)
	// AutoModerationRuleDelete(guildID, ruleID string) (err error)
	Open() error
	HeartbeatLatency() time.Duration
	UpdateGameStatus(idle int, name string) (err error)
	UpdateStreamingStatus(idle int, name string, url string) (err error)
	UpdateListeningStatus(name string) (err error)
	UpdateStatusComplex(usd UpdateStatusData) (err error)
	RequestGuildMembers(guildID, query string, limit int, nonce string, presences bool) error
	RequestGuildMembersList(guildID string, userIDs []string, limit int, nonce string, presences bool) error
	RequestGuildMembersBatch(guildIDs []string, query string, limit int, nonce string, presences bool) (err error)
	RequestGuildMembersBatchList(guildIDs []string, userIDs []string, limit int, nonce string, presences bool) (err error)
	ChannelVoiceJoin(gID, cID string, mute, deaf bool) (voice *VoiceConnection, err error)
	ChannelVoiceJoinManual(gID, cID string, mute, deaf bool) (err error)
	Close() error
	CloseWithCode(closeCode int) (err error)
}

// Verify *discordgo.Session matches ISession
var _ ISession = (*discordgo.Session)(nil)
