package front

//AssigneeURL item which represents the teammate link
type AssigneeURL struct {
	Links Links `json:"assignee"`
}

//Conversation item
type Conversation struct {
	ID       string      `json:"id"`
	Subject  string      `json:"subject"`
	Status   string      `json:"status"`
	Teammate AssigneeURL `json:"assignee"`
}

//ConversationList list of inbox conversations
type ConversationList struct {
	Links   Links          `json:"_links"`
	Results []Conversation `json:"_results"`
}
