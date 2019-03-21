package dmux

type MessageContext interface {
	Guild() Guild
	Channel() Channel
	Message() Message
	User() User
}

type HandlerContext interface {
	MessageContext() (bool, MessageContext)
}

type Handler interface {
	Handle(c Session, context HandlerContext)
	HandlerType() HandlerType
	Name() string
	Description() string
	Pattern(HandlerContext) bool
}

type Message interface {
	ID() string
	Content() string
	Author() User
}

type User interface {
	ID() string
	Name() string
	Admin(s Session, c Channel) (bool, error)
}

type Channel interface {
	ID() string
}

type Guild interface {
	ID() string
	Roles(s Session) ([]Role, error)
}

type Role interface {
	Name() string
	ID() string
}

type Permissions interface {
	isAdmin() bool
}

type Session interface {
	MessageChannel(c Channel, m Message) (Message, error)
	GuildRoles(g Guild) ([]Role, error)
	GuildMemberRoleAdd(g Guild, u User, r Role) error
	GuildMemberRoleRemove(g Guild, u User, r Role) error
	UserPermissions(u User, c Channel) (Permissions, error)
	AddHandler(Handler)
	Open()
	Close()
}
