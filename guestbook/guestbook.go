package guestbook

import "fmt"

type GuestBookEntry struct {
	Id      int
	Email   string
	Title   string
	Content string
}

type GuestBook struct {
	GuestBookData []*GuestBookEntry
}

func NewGuestBook() *GuestBook {
	return &GuestBook{
		make([]*GuestBookEntry, 0),
	}
}

func (g *GuestBook) AddEntry(email, title, content string) int {
	length := len(g.GuestBookData)
	id := length

	entry := &GuestBookEntry{
		Email:   email,
		Title:   title,
		Content: content,
		Id:      id,
	}

	g.GuestBookData = append(g.GuestBookData, entry)
	return id
}

func (g *GuestBook) DeleteEntry(id int) error {
	if id >= 0 && id < len(g.GuestBookData) && g.GuestBookData[id] != nil {
		g.GuestBookData[id] = nil
	} else {
		return fmt.Errorf("invalid id")
	}
	return nil
}

func (g *GuestBook) GetEntry(id int) (*GuestBookEntry, error) {
	if g.GuestBookData[id] != nil {
		return g.GuestBookData[id], nil
	} else {
		return nil, fmt.Errorf("invalid id")
	}
}

func (g *GuestBook) GetAllEntries() []*GuestBookEntry {
	var entries []*GuestBookEntry

	for _, entry := range g.GuestBookData {
		if entry != nil {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (g *GuestBook) DeleteAllEntries() {
	g.GuestBookData = make([]*GuestBookEntry, 0)
}
