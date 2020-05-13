package lib

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andersfylling/disgord"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ShortStr truncates a string by n length.
func ShortStr(str string, n int, els ...string) string {
	runes := []rune(str)
	el := JoinString(els, " ")

	if el == "" {
		el = "..."
	}

	if len(runes) > n {
		return string(runes[:n]) + el
	}

	return str
}

// JoinString joins a string slice with a char, and removes the end char.
func JoinString(strs []string, char string) string {
	return strings.TrimRight(strings.Join(strs, char), char)
}

// JoinStringMap joins a string map with a char, and removes the end char.
func JoinStringMap(strs map[int]string, char string) string {
	// make sure map is sorted in order cause Go likes random orders
	// for some stupid fucking reason
	var keys []int

	for k := range strs {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var vals []string

	for _, k := range keys {
		vals = append(vals, strs[k])
	}

	return JoinString(vals, char)
}

// Ucwords capitalizes the first letter in each word. (Mirror's PHP's ucwords function)
func Ucwords(str string) string {
	return strings.Title(str)
}

// HumanNumber converts a number into a human readable one.
func HumanNumber(i interface{}) string {
	printer := message.NewPrinter(language.English)

	if reflect.TypeOf(i).String() == "string" {
		a := fmt.Sprintf("%v", i)
		i, _ = strconv.Atoi(a)
	}

	return printer.Sprintf("%d", i)
}

// GenerateMessageURL returns the URL for a specific Discord message.
func GenerateMessageURL(guildID disgord.Snowflake, msg *disgord.Message) string {
	return fmt.Sprintf("https://discordapp.com/channels/%s/%s/%s", guildID, msg.ChannelID, msg.ID)
}

// GetDiscordIDFromMention gets the snowflake id from a mention.
func GetDiscordIDFromMention(mention string) (ID disgord.Snowflake, err error) {
	if IsMention(mention) {
		ID = StrToSnowflake(strings.TrimLeft(strings.TrimLeft(strings.TrimRight(mention, ">"), "<@"), "!"))
	} else {
		err = errors.New("given string is not a Discord mention")
	}

	return
}

// SnowflakeToUInt64 returns a uint64 version of a snowflake.
func SnowflakeToUInt64(snowflake disgord.Snowflake) uint64 {
	did, _ := strconv.Atoi(snowflake.String())

	return uint64(did)
}

// StrToSnowflake returns a Snowflake from a string.
func StrToSnowflake(str string) disgord.Snowflake {
	did, _ := strconv.Atoi(str)

	return UInt64ToSnowflake(uint64(did))
}

// UInt64ToSnowflake converts a uint64 to a snowflake.
func UInt64ToSnowflake(i uint64) disgord.Snowflake {
	return disgord.NewSnowflake(i)
}

// IsMention checks if a string is a discord mention.
func IsMention(mention string) bool {
	return strings.HasPrefix(mention, "<@")
}

// HTMLParse parses an html string, and replaces anchor tags with bbcode links.
func HTMLParse(s string) string {
	// Strip tags that are not anchor tags
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	p.AllowAttrs("href").OnElements("a")
	s = p.Sanitize(s)

	// Search through text, get anchor tags, and
	// replace them with bbcode links instead
	sr := strings.NewReader(s)
	doc, _ := goquery.NewDocumentFromReader(sr)
	sel := doc.Find("a")

	for i := range sel.Nodes {
		single := sel.Eq(i)
		href, ok := single.Attr("href")

		if ok {
			text := single.Text()
			single.ReplaceWithHtml(fmt.Sprintf("[%s](%s)", text, href))
		}
	}

	return doc.Contents().Text()
}
