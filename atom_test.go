package atom

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntry(t *testing.T) {
	entryStr := "<entry><title>foo</title><content>bar</content></entry>"
	entry, err := DecodeEntry(strings.NewReader(entryStr))
	require.Nil(t, err)
	require.NotNil(t, entry)
	assert.Equal(t, entry.Content.Raw, "bar", "entry.content")
	assert.Equal(t, entry.Title.Raw, "foo", "entry.title")
	assert.Nil(t, entry.Author, "entry.author")
}

func TestFeed(t *testing.T) {
	feedStr := `
		<feed xmlns="http://www.w3.org/2005/Atom">

			<title>Example Feed</title>
			<subtitle>A subtitle.</subtitle>
			<link href="http://example.org/feed/" rel="self" />
			<link href="http://example.org/" />
			<id>urn:uuid:60a76c80-d399-11d9-b91C-0003939e0af6</id>
			<updated>2003-12-13T18:30:02Z</updated>
			
			
			<entry>
				<title>Atom-Powered Robots Run Amok</title>
				<link href="http://example.org/2003/12/13/atom03" />
				<link rel="alternate" type="text/html" href="http://example.org/2003/12/13/atom03.html"/>
				<link rel="edit" href="http://example.org/2003/12/13/atom03/edit"/>
				<id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
				<updated>2003-12-13T18:30:02Z</updated>
				<summary>Some text.</summary>
				<content type="xhtml">
					<div xmlns="http://www.w3.org/1999/xhtml">
						<p>This is the entry content.</p>
					</div>
				</content>
				<author>
					<name>John Doe</name>
					<email>johndoe@example.com</email>
				</author>
			</entry>

		</feed>`
	feed, err := DecodeFeed(strings.NewReader(feedStr))
	require.Nil(t, err)
	require.NotNil(t, feed)
	assert.Equal(t, feed.Title.Raw, "Example Feed", "feed title")
	assert.Equal(t, feed.Subtitle.Raw, "A subtitle.", "feed subtitle")
	assert.Equal(t, *feed.Updated, "2003-12-13T18:30:02Z", "feed updated")
	assert.Equal(t, *feed.Id, "urn:uuid:60a76c80-d399-11d9-b91C-0003939e0af6",
		"feed id")
	assert.Nil(t, feed.Rights, "feed rights")
	assert.Equal(t, len(feed.Links), 2, "feed num links")
	require.NotNil(t, feed.Entries, "feed entries is not nil")
	require.Equal(t, len(feed.Entries), 1, "feed num entries")

	entry := feed.Entries[0]
	assert.Equal(t, len(entry.Links), 3, "num links")
	assert.Equal(t, entry.Title.Raw, "Atom-Powered Robots Run Amok",
		"entry.title")
	assert.Contains(t, entry.Content.Raw, "<p>This is the entry content.</p>",
		"entry.content")
	assert.Equal(t, *entry.Author.Name, "John Doe", "entry author name")
	assert.Equal(t, *entry.Author.Email, "johndoe@example.com",
		"entry author email")
	assert.Equal(t, *entry.Summary, "Some text.", "entry summary")
	assert.Equal(t, *entry.Updated, "2003-12-13T18:30:02Z", "entry updated")
}
