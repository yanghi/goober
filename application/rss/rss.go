package rss

type Feed struct {
	Title       string   `json:"title"`
	FeedType    string   `json:"feedType"`
	Link        string   `json:"link"`
	LinkType    string   `json:"linkType"`
	Description string   `json:"description"`
	Author      []string `json:"author"`
}
type FeedAuthor struct {
}
type FeedItem struct {
	GUID        string   `json:"guid"`
	PublishTime string   `json:"publishTime"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Category    []string `json:"category"`
	Author      []string `json:"author"`
}
