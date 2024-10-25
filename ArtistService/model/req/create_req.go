package req

type ArtistCreateReq struct {
	Firstname   string   `json:"firstname" validate:"firstname"`
	Lastname    string   `json:"lastname" validate:"lastname"`
	Email       string   `json:"email" validate:"email"`
	PhoneNumber string   `json:"phone_number" validate:"phone_number"`
	Biography   string   `json:"biography" validate:"biography"`
	SocialLinks []string `json:"social_links" validate:"social_links"`
}
