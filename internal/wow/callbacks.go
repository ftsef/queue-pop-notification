package wow

type PvPMode string

const (
	PvP_2v2          PvPMode = "2v2"
	PvP_3v3          PvPMode = "3v3"
	PvP_5v5          PvPMode = "5v5"
	PvP_Solo_Shuffle PvPMode = "solo-shuffle"
	PvP_BG_Blitz     PvPMode = "bg-blitz"
	PvP_BG_Random    PvPMode = "bg-random"
	PvP_Rated_BG     PvPMode = "rated-bg"
)

type EventCallbacks struct {
	OnPvPQueuePop func(mode PvPMode, details map[string]string)
}
