package models

import (
	"encoding/xml"
)

type Url struct {
	Url xml.Name `xml:"url"`
	Loc string   `xml:"loc"`
}

type UrlSet struct {
	XMLUrlSet xml.Name `xml:"urlSet"`
	Urls      []Url    `xml:"url"`
}
