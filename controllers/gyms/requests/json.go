package requests

// GymAdd add new Gym
type GymAdd struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GymUpdate update gym request
type GymUpdate struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
