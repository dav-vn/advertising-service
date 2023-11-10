package models

import "fmt"

type Advertisement struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
}

type AdvertisementStorage struct {
	Advertisements []Advertisement
}

func (s *AdvertisementStorage) GetAdvertisements() []Advertisement {
	return s.Advertisements
}

func (s *AdvertisementStorage) GetAdvertisementByID(id int) (*Advertisement, error) {
	for _, ad := range s.Advertisements {
		if ad.ID == id {
			return &ad, nil
		}
	}

	return nil, fmt.Errorf("advertisemnt not found")
}

func (s *AdvertisementStorage) CreateAdvertisement(advertisement Advertisement) (*Advertisement, error) {
	if len(advertisement.Photos) > 3 {
		return nil, fmt.Errorf("too many photos")
	}
	if len(advertisement.Description) > 1000 {
		return nil, fmt.Errorf("too long decription")
	}
	if len(advertisement.Title) > 200 {
		return nil, fmt.Errorf("too long Title")
	}

	advertisement.ID = len(s.Advertisements) + 1

	return &advertisement, nil
}
