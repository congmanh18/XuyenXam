package req

import "github.com/congmanh18/XuyenXam/ArtistService/entity"

type ArtistUpdateReq struct {
	Firstname    string         `json:"firstname"`
	Lastname     string         `json:"lastname"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Biography    string         `json:"biography"`
	Styles       []string       `json:"styles"`
	Availability []entity.Slot  `json:"availability"`
	Location     entity.Address `json:"location"`
	SocialLinks  []string       `json:"social"`
}
