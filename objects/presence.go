package objects

import "github.com/rxdn/gdl/utils"

type Presence struct {
	User         *User
	Roles        utils.Uint64StringSlice `json:",string"`
	Game         Activity
	GuildId      uint64 `json:",string"`
	Status       string
	Activities   []*Activity
	ClientStatus ClientStatus
}

type Activity struct {
	Name          string
	Type          int
	Url           string
	Timestamps    Timestamp
	ApplicationId string
	Details       string
	State         string
	Party         Party
	Assets        Asset
	Secret        Secret
	Instance      bool
	Flags         int
}

type Timestamp struct {
	Start int
	End   int
}

type Party struct {
	Id   string
	Size []int
}

type Secret struct {
	Join     string
	Spectate string
	Match    string
}

type Asset struct {
	LargeImage string
	LargeText  string
	SmallImage string
	SmallText  string
}

type ClientStatus struct {
	Desktop string
	Mobile  string
	Web     string
}
