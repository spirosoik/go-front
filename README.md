# Golang Front API

This is an implementation in Go for [Front API](https://dev.frontapp.com/)

# Implemented resources

- [x] Token Identity
- [x] Teams
- [x] Teammates
- [x] Inboxes
- [x] Channels
- [x] Conversations
- [ ] Contacts
- [ ] Topics
- [ ] Webhooks
- [ ] Analytics
- [ ] Events
- [ ] Contact groups
- [ ] Contact handles
- [ ] Contact notes
- [ ] Tags
- [ ] Rules
- [ ] Exports
- [ ] Attachments

The rest endpoints are coming soon!!! Stay tuned

# Installation

```
go get -v github.com/spirosoik/go-front
```

# Example

```
import (
	"fmt"

	"github.com/spirosoik/go-front/front"
)

func main() {
	cfg := front.Config{
		APIToken: "test_token",
		BaseURL:  "https://api2.frontapp.com",
	}

	c, err := front.New(&cfg)
	if err != nil {
		fmt.Print(err)
	}

	data, err := c.Inbox.List()
	fmt.Print(data.Results)
}
```

# License

The librady is under MIT [LICENSE](./LICENSE)