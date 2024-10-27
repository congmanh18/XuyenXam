package mapper

import (
	"github.com/congmanh18/XuyenXam/ArtistService/entity"
)

func MergeArtistEntity(existing *entity.Artist, new *entity.Artist) *entity.Artist {
	if new.Firstname != nil {
		existing.Firstname = new.Firstname
	}
	if new.Lastname != nil {
		existing.Lastname = new.Lastname
	}
	if new.Username != nil {
		existing.Username = new.Username
	}
	if new.Email != nil {
		existing.Email = new.Email
	}
	if new.Biography != nil {
		existing.Biography = new.Biography
	}
	if new.Styles != nil {
		existing.Styles = new.Styles
	}
	if new.Availability != nil {
		existing.Availability = new.Availability
	}
	if new.Location != nil {
		existing.Location = new.Location
	}
	if new.SocialLinks != nil {
		existing.SocialLinks = new.SocialLinks
	}
	return existing
}
