package rendering

// LoginResponse holds the standard response for a successfull or unsucessfull
// login attempt
type LoginResponse struct {
	Success bool
	Code    int
	Reason  string
}
