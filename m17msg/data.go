package main

import (
	"fyne.io/fyne/v2"
)

type appData struct {
	servers []*server
}

type server struct {
	id           string
	name         string
	iconURL      string
	iconResource fyne.Resource
	channels     []*channel
	service      service
	// users         map[string]*user
}

func (s *server) icon() fyne.Resource {
	if s.iconResource != nil {
		return s.iconResource
	}

	icon, err := fyne.LoadResourceFromURLString(s.iconURL)
	if err != nil {
		fyne.LogError("Failed to read icon "+s.iconURL, err)
		return nil
	}

	s.iconResource = icon
	return icon
}

type channel struct {
	id       string
	name     string
	messages []*message
	server   *server
}

type message struct {
	content string
	user    *user
}

type user struct {
	name, username, avatarURL string
}

func findChan(d *appData, sID, cID string) *channel {
	for _, s := range d.servers {
		if s.id == sID {
			if c := findServerChan(s, cID); c != nil {
				return c
			}
		}
	}
	return nil
}

func findServerChan(s *server, cID string) *channel {
	for _, c := range s.channels {
		if c.id == cID {
			return c
		}
	}
	return nil
}
