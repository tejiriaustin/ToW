package housekeeper

import (
	"github.com/robfig/cron"
)

type HouseKeeping interface {
	AddKeeper(delay string, job func()) *HouseKeeper
	StartHouseKeeping()
	StopHouseKeeping()
}

type HouseKeeper struct {
	cron *cron.Cron
}

func (h *HouseKeeper) StartHouseKeeping() {
	h.cron.Start()
}

func (h *HouseKeeper) StopHouseKeeping() {
	h.cron.Stop()
}

func (h *HouseKeeper) AddKeeper(delay string, job func()) *HouseKeeper {
	err := h.cron.AddFunc(delay, job)
	if err != nil {
		return nil
	}
	return h
}

func NewKeeper() *HouseKeeper {
	return &HouseKeeper{cron: cron.New()}
}

var _ HouseKeeping = (*HouseKeeper)(nil)
