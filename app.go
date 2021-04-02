package roku

import (
	"sort"
	"strings"
)

// App represents an app installed on the device
type App struct {
	Text    string `xml:",chardata"`
	ID      int    `xml:"id,attr"`
	Type    string `xml:"type,attr"`
	Version string `xml:"version,attr"`
}

// Apps is a slice of App
type Apps struct {
	Apps []*App `xml:"app"`
}

// ActiveApp details the app that is currently active on the device
type ActiveApp struct {
	App         string `xml:"app"`
	Screensaver struct {
		Text    string `xml:",chardata"`
		ID      int    `xml:"id,attr"`
		Type    string `xml:"type,attr"`
		Version string `xml:"version,attr"`
	} `xml:"screensaver"`
}

// sort.Sort interface implementation
type sortByID Apps

func (a sortByID) Len() int           { return len(a.Apps) }
func (a sortByID) Swap(i, j int)      { a.Apps[i], a.Apps[j] = a.Apps[j], a.Apps[i] }
func (a sortByID) Less(i, j int) bool { return a.Apps[i].ID <= a.Apps[j].ID }

// FindById searches for an app by ID
// Returns the app and true if found
func (a *Apps) FindByID(id int) (*App, bool) {

	sort.Sort(sortByID(*a))

	// binary search for idx
	i := sort.Search(len(a.Apps), func(i int) bool {
		// >= in search
		return a.Apps[i].ID >= id
	})

	// found
	if i < len(a.Apps) && a.Apps[i].ID == id {
		return a.Apps[i], true
	}

	return &App{}, false
}

// FindByName searches for an app by name
// Returns the app and true if found
func (a *Apps) FindByName(text string) (*App, bool) {

	// sort the slice
	sort.Slice(a.Apps, func(i, j int) bool {
		return strings.ToLower(a.Apps[i].Text) <= strings.ToLower(a.Apps[j].Text)
	})

	text = strings.ToLower(text)

	// binary search >=
	i := sort.Search(len(a.Apps), func(i int) bool {
		return strings.ToLower(a.Apps[i].Text) >= text
	})

	// found
	if i < len(a.Apps) && strings.ToLower(a.Apps[i].Text) == text {
		return a.Apps[i], true
	}

	return &App{}, false
}
