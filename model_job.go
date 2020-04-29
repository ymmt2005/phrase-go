package phrase
import (
	"time"
)
// Job struct for Job
type Job struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Briefing string `json:"briefing,omitempty"`
	DueDate time.Time `json:"due_date,omitempty"`
	State string `json:"state,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}