package state

import "sync"

type Stats struct {
	Memory    string `json:"memory"`
	Disk      string `json:"disk"`
	CPU       string `json:"cpu"`
	Bandwidth string `json:"bandwidth"`
}

var SystemStats struct {
	Stat Stats
	Mu   sync.RWMutex
}
