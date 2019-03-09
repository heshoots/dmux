package dmux

import (
	"errors"
)

type MockSession struct {
	messages        []Message
	roles           []Role
	users           map[string][]string
	userPermissions map[string]Permissions
}

type MockRole struct {
	name string
	id   string
}

func (r MockRole) Name() string {
	return r.name
}

func (r MockRole) ID() string {
	return r.id
}

type MockMessage struct {
	string
}

func (m MockSession) MessageChannel(channelID, message string) (Message, error) {
	m.messages = append(m.messages, message)
	return MockMessage{message}, nil
}

func (m MockSession) GuildRoles(guildID string) ([]Role, error) {
	return m.roles, nil
}

func (m MockSession) GuildMemberRoleAdd(guildID, userID, roleID string) error {
	m.users[userID] = append(m.users[userID], roleID)
	return nil
}

func (m MockSession) UserChannelPermissions(userID, channelID string) (Permissions, error) {
	if val, ok := m.userPermissions[userID]; ok {
		return val, nil
	}
	return nil, errors.New("user does not have permissions")
}

func (m MockSession) GuildMemberRoleRemove(guildID, userID, roleID string) error {
	var roles []string
	for _, role := range m.users[userID] {
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
