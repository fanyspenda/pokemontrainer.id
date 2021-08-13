package requests

// GymAdd add new Gym
type GymAdd struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
