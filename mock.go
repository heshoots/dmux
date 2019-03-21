package dmux

import (
	"errors"
)

type MockGuild struct {
	SessionRoles           []Role
	SessionUsers           map[string][]string
	SessionUserPermissions map[string]Permissions
	SessionMessages        []Message
}

func (g MockGuild) ID() string {
	return "mockguild"
}

func (g MockGuild) Roles(s Session) ([]Role, error) {
	return g.SessionRoles, nil
}

type MockSession struct {
	Guild MockGuild
}

type MockRole struct {
	RoleName string
	RoleID   string
}

func (r MockRole) Name() string {
	return r.RoleName
}

func (r MockRole) ID() string {
	return r.RoleID
}

type MockMessage struct {
	id      string
	content string
	user    User
}

type MockUser struct {
	id    string
	name  string
	admin bool
}

func CreateMockMessage(id string, content string, u User) MockMessage {
	return MockMessage{id, content, u}
}

func CreateMockUser(id string, name string, admin bool) MockUser {
	return MockUser{id, name, admin}
}

func (u MockUser) Admin(s Session, c Channel) (bool, error) {
	return u.admin, nil
}

func (u MockUser) ID() string {
	return u.id
}

func (u MockUser) Name() string {
	return u.name
}

func (m MockMessage) ID() string {
	return m.id
}

func (m MockMessage) Content() string {
	return m.content
}

func (m MockMessage) Author() User {
	return m.user
}

func (m MockSession) MessageChannel(channel Channel, message Message) (Message, error) {
	m.Guild.SessionMessages = append(m.Guild.SessionMessages, message)
	return message, nil
}

func (m MockSession) GuildRoles(guild Guild) ([]Role, error) {
	return m.Guild.SessionRoles, nil
}

func (m MockSession) GuildMemberRoleAdd(guild Guild, user User, role Role) error {
	m.Guild.SessionUsers[user.ID()] = append(m.Guild.SessionUsers[user.ID()], role.ID())
	return nil
}

func (m MockSession) UserPermissions(user User, channel Channel) (Permissions, error) {
	if val, ok := m.Guild.SessionUserPermissions[user.ID()]; ok {
		return val, nil
	}
	return nil, errors.New("user does not have permissions")
}

func (m MockSession) AddHandler(h Handler) {
	return
}

func (m MockSession) Open() {
	return
}

func (m MockSession) Close() {
	return
}

func (m MockSession) GuildMemberRoleRemove(guild Guild, user User, roleremove Role) error {
	var roles []string
	for _, role := range m.Guild.SessionUsers[user.ID()] {
		if role != roleremove.Name() {
			roles = append(roles, role)
		}
	}
	return nil
}

type MockMessageContext struct {
	guildID   string
	channelID string
	message   string
	userID    string
	userName  string
	userAdmin bool
}

func (m MockMessageContext) GuildID() string {
	return m.guildID
}

func (m MockMessageContext) ChannelID() string {
	return m.channelID
}

func (m MockMessageContext) Message() string {
	return m.message
}

func (m MockMessageContext) UserID() string {
	return m.userID
}

func (m MockMessageContext) UserName() string {
	return m.userName
}

func (m MockMessageContext) UserAdmin() bool {
	return m.userAdmin
}
