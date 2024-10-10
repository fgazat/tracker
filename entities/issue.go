package entities

// Issue entity
// https://yandex.cloud/en/docs/tracker/concepts/issues/get-issue
type Issue struct {
	Self    string `json:"self,omitempty"`
	ID      string `json:"id,omitempty"`
	Key     string `json:"key,omitempty"`
	Version int64  `json:"version,omitempty"`

	LastCommentUpdatedAt string   `json:"lastCommentUpdatedAt,omitempty"`
	Parent               *Entity  `json:"parent,omitempty"`
	Queue                *Entity  `json:"queue,omitempty"`
	Summary              string   `json:"summary,omitempty"`
	Description          string   `json:"description,omitempty"`
	Status               *Entity  `json:"status,omitempty"`
	Type                 *Entity  `json:"type,omitempty"`
	Priority             *Entity  `json:"priority,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
	CreatedAt            string   `json:"createdAt,omitempty"`
	UpdatedAt            string   `json:"updatedAt,omitempty"`
	StartDate            string   `json:"start,omitempty"`
	DeadlineDate         string   `json:"deadline,omitempty"`

	Author    *User  `json:"author,omitempty"`
	Assignee  *User  `json:"assignee,omitempty"`
	UpdatedBy *User  `json:"updatedBy,omitempty"`
	CreatedBy *User  `json:"createdBy,omitempty"`
	Followers []User `json:"followers,omitempty"`
}

type Entity struct {
	Self    string `json:"self,omitempty"`
	ID      string `json:"id,omitempty"`
	Key     string `json:"key,omitempty"`
	Display string `json:"display,omitempty"`
}

type User struct {
	Self    string `json:"self,omitempty"`
	ID      string `json:"id,omitempty"`
	Display string `json:"display,omitempty"`
}
