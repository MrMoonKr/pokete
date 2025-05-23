// DO NOT EDIT!
// This code was auto generated by the `protoc-gen-pokete-resources` plugin,
// part of the pokete project, by <lxgr@protonmail.com>

package resources

import ()

type TrainerPokeArgs struct {
	Name string `json:"name"`
	Xp   int32  `json:"xp"`
}

type TrainerArgs struct {
	Name      string   `json:"name"`
	Gender    string   `json:"gender"`
	Texts     []string `json:"texts"`
	LoseTexts []string `json:"lose_texts"`
	WinTexts  []string `json:"win_texts"`
	X         int32    `json:"x"`
	Y         int32    `json:"y"`
}

type Trainer struct {
	Pokes []TrainerPokeArgs `json:"pokes"`
	Args  TrainerArgs       `json:"args"`
}
