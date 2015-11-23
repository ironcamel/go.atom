package atom

import (
	"encoding/xml"
	"io"
)

type XMLFeed struct {
	XMLName      xml.Name         `xml:"feed"`
	Namespace    *string          `xml:"xmlns,attr"`
	Id           *string          `xml:"id"`
	Title        *XMLTitle        `xml:"title"`
	Updated      *string          `xml:"updated"`
	Author       *XMLAuthor       `xml:"author"`
	Entries      []XMLEntry       `xml:"entry"`
	Links        []XMLLink        `xml:"link"`
	Categories   []XMLCategory    `xml:"category"`
	Contributors []XMLContributor `xml:"contributor"`
	Generator    *XMLGenerator    `xml:"generator"`
	Icon         *string          `xml:"icon"`
	Logo         *string          `xml:"logo"`
	Rights       *XMLRights       `xml:"rights"`
	Subtitle     *XMLSubtitle     `xml:"subtitle"`
}

type XMLEntry struct {
	XMLName      xml.Name         `xml:"entry"`
	Namespace    *string          `xml:"xmlns,attr"`
	Id           *string          `xml:"id"`
	Title        *XMLTitle        `xml:"title"`
	Updated      *string          `xml:"updated"`
	Author       *XMLAuthor       `xml:"author"`
	Content      *XMLEntryContent `xml:"content"`
	Links        []XMLLink        `xml:"link"`
	Summary      *string          `xml:"summary"`
	Categories   []XMLCategory    `xml:"category"`
	Contributors []XMLContributor `xml:"contributor"`
	Published    *string          `xml:"published"`
}

type XMLEntryContent struct {
	XMLName xml.Name `xml:"content"`
	Type    *string  `xml:"type,attr"`
	Raw     string   `xml:",innerxml"`
}

type XMLAuthor struct {
	XMLName xml.Name `xml:"author"`
	Name    *string  `xml:"name"`
	Email   *string  `xml:"email"`
	URI     *string  `xml:"uri"`
}

type XMLContributor struct {
	XMLName xml.Name `xml:"contributor"`
	Name    *string  `xml:"name"`
}

type XMLGenerator struct {
	XMLName xml.Name `xml:"generator"`
	URI     *string  `xml:"uri,attr"`
	Version *string  `xml:"version,attr"`
	Text    *string  `xml:",chardata"`
}

type XMLLink struct {
	XMLName  xml.Name `xml:"link"`
	Href     *string  `xml:"href,attr"`
	Rel      *string  `xml:"rel,attr"`
	Hreflang *string  `xml:"hreflang,attr"`
	Title    *string  `xml:"title,attr"`
	Length   *string  `xml:"length,attr"`
}

type XMLCategory struct {
	XMLName xml.Name `xml:"category"`
	Term    *string  `xml:"term,attr"`
}

type XMLTitle struct {
	XMLName xml.Name `xml:"title"`
	Type    *string  `xml:"type,attr"`
	Raw     string   `xml:",innerxml"`
}

type XMLSummary struct {
	XMLName xml.Name `xml:"summary"`
	Type    *string  `xml:"type,attr"`
	Raw     string   `xml:",innerxml"`
}

type XMLRights struct {
	XMLName xml.Name `xml:"rights"`
	Type    *string  `xml:"type,attr"`
	Raw     string   `xml:",innerxml"`
}

type XMLSubtitle struct {
	XMLName xml.Name `xml:"subtitle"`
	Type    *string  `xml:"type,attr"`
	Raw     string   `xml:",innerxml"`
}

type XMLMultiType struct {
	Type *string `xml:"type,attr"`
	Raw  string  `xml:",innerxml"`
}

func DecodeEntry(r io.Reader) (*XMLEntry, error) {
	var entry XMLEntry
	err := xml.NewDecoder(r).Decode(&entry)
	return &entry, err
}
