package store

import . "github.com/agencyrevolution/go-microservices-example/models"

type Engine struct {
	State *State
}

type State struct {
	Users []*User
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) InitializeState() {
	s := &State{
		Users: []*User{
			&User{
				Name:        "agencyrevolution",
				ID:          1,
				Description: "Help insurance agencies create Meaningful Relationships with their customers and marketplace.",
			},
			&User{
				Name:        "facebook",
				ID:          69631,
				Description: "We work hard to contribute our work back to the web, mobile, big data, & infrastructure communities. NB: members must have two-factor auth.",
			},
		},
	}

	e.State = s
}
