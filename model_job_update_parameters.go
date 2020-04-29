package phrase
import (
	"time"
)
// JobUpdateParameters struct for JobUpdateParameters
type JobUpdateParameters struct {
	// specify the branch to use
	Branch string `json:"branch,omitempty"`
	// Job name
	Name string `json:"name,omitempty"`
	// Briefing for the translators
	Briefing string `json:"briefing,omitempty"`
	// Date the job should be finished
	DueDate time.Time `json:"due_date,omitempty"`
}