package base

import "github.com/lavab/api/utils"

// Resource is the base struct for every resource that needs to be saved to db and marshalled with json.
type Resource struct {
	// ID is the resources ID, used as a primary index by the db.
	ID string `json:"id" gorethink:"id"`

	// DateChanged is an RFC3339-encoded time string that lets clients poll whether a resource has changed.
	DateChanged string `json:"date_changed" gorethink:"date_changed"`

	// DateCreated is an RFC3339-encoded string that stores the creation date of a resource.
	DateCreated string `json:"date_created" gorethink:"date_created"`

	// Name is a human-friendly description of the resource.
	// Sometimes it can be essential to the resource, e.g. the `User.Name` field.
	Name string `json:"name" gorethink:"name,omitempty"`

	// UserID is the ID of the user to which this resource belongs.
	UserID string `json:"user_id" gorethink:"user_id"`
}

// MakeResource creates a new Resource object with sane defaults.
func MakeResource(name, userID string) Resource {
	timeString := utils.TimeNowString()
	return Resource{
		ID:          utils.UUID(),
		DateChanged: timeString,
		DateCreated: timeString,
		Name:        name,
		UserID:      userID,
	}
}