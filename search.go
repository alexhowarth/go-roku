package roku

import (
	"github.com/koron/go-ssdp"
)

const searchType string = "roku:ecp"

// Search uses SSDP to discover all Roku devices on the local network
func Search(timeout int) (*[]ssdp.Service, error) {
	var found []ssdp.Service

	list, err := ssdp.Search(searchType, timeout, "")
	if err != nil {
		return &found, err
	}

	found = append(found, list...)

	return &found, nil
}
