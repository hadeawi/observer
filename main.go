package main

import "log"

func main() {
	var p Player

	p = newPlayer()

	p.broadcast("hello")

	s := newHP("100")
	p.addHP(s)
	p.broadcast("you're alive")

	p.removeHP(s.id())
	p.broadcast("you're dead")

}

type Player interface {
	addHP(hp HP)
	removeHP(subId string)
	broadcast(msg string)
}

type HP interface {
	id() string
	react(msg string)
}

type player struct {
	hp map[string]HP
}

func newPlayer() player {
	return player{hp: make(map[string]HP)}
}

func (p player) addHP(hp HP) {
	p.hp[hp.id()] = hp
}

func (p player) removeHP(subId string) {
	delete(p.hp, subId)
}
func (p player) broadcast(msg string) {
	for _, hp := range p.hp {
		hp.react(msg)
	}
}

type hp struct {
	subId string
}

func newHP(subId string) hp {
	return hp{subId: subId}
}
func (s hp) id() string {
	return s.subId
}
func (s hp) react(msg string) {
	log.Printf("Player %s - received: %s", s.subId, msg)
}
