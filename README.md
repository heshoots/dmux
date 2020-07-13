# dmux
discord router 

This package adds a small wrapper to allow for easier extension of bwmarrin's discordgo library, It provides the following

- A net/http inspired router/handler interface for building out a simple bot command pattern
- An example middleware RegexMessageHandler, which uses a Regex to match bot commends
- Mock Servers, Role and Guilds to allow you to test functionallity

some example code for setting up a basic echo server

```
func echo(s dmux.Session, context dmux.RegexHandlerContext) {
	ok, ctx := context.MessageContext()
	if ok {
    message := context.Groups()["message"]
		SendMessage(s, ctx.Channel(), message)
		return
	}
}

fun main() {
	discordInstance, err := dmux.Router(authToken)
	if err != nil {
		panic(err)
		return
	}
  
  handler := &dmux.DiscordRegexMessageHandler{
			HandlerPattern:     `^!echo (?P<message>.*)$`,
			HandlerFn:          echo,
			HandlerName:        "!echo",
			RequiresAdmin:      false,
			HandlerDescription: "echo a message",
	},
  
  discordInstance.AddHandler(handler)
	discordInstance.Open()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discordInstance.Close()
}
```
