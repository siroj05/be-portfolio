package dto

type ProfileDto struct {
	ID         string `json:"id"`
	UserID     int64  `json:"userId"`
	ImagePath  string `json:"imagePath"`
	FullName   string `json:"fullName"`
	JobTitle   string `json:"jobTitle"`
	Email      string `json:"email"`
	Linkedin   string `json:"linkedin"`
	Repository string `json:"repository"`
	About      string `json:"about"`
}
