package res

type ArtistReadRes struct {
	Firstname    string     `json:"firstname"`
	Lastname     string     `json:"lastname"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	PhoneNumber  string     `json:"phone_number"`
	Biography    string     `json:"biography"`
	Styles       []string   `json:"styles"`
	Rating       string     `json:"rating"`
	Availability []struct{} `json:"availability"`
	Gallery      []struct{} `json:"gallery"`
	Location     struct{}   `json:"location"`
	SocialLinks  []string   `json:"social_links"`
}
