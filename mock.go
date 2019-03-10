package dmux

import (
	"errors"
)

type MockSession struct {
	SessionMessages        []Message
	SessionRoles           []Role
	SessionUsers           map[string][]string
	SessionUserPermissions map[string]Permissions
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
	string
}

func (m MockSession) MessageChannel(channelID, message string) (Message, error) {
	m.SessionMessages = append(m.SessionMessages, message)
	return MockMessage{message}, nil
}

func (m MockSession) GuildRoles(guildID string) ([]Role, error) {
	return m.SessionRoles, nil
}

func (m MockSession) GuildMemberRoleAdd(guildID, userID, roleID string) error {
	m.SessionUsers[userID] = append(m.SessionUsers[userID], roleID)
	return nil
}

func (m MockSession) UserChannelPermissions(userID, channelID string) (Permissions, error) {
	if val, ok := m.SessionUserPermissions[userID]; ok {
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

func (m MockSession) GuildMemberRoleRemove(guildID, userID, roleID string) error {
	var roles []string
	for _, role := range m.SessionUsers[userID] {
		if role != roleID {
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
