package main

import "fmt"

type roda interface {
	getRoda() string
}

type rodaKaret struct{}

func (r rodaKaret) getRoda() string {
	return "Roda Karet"
}

type rodaKayu struct{}

func (r rodaKayu) getRoda() string {
	return "Roda Kayu"
}

type rodaBesi struct{}

func (r rodaBesi) getRoda() string {
	return "Roda Besi"
}

type pintu interface {
	getBunyi() string
	getBunyiBuka() string
}

type pintuKananType struct{}

func (p pintuKananType) getBunyi() string {
	return "Knock"
}

func (p pintuKananType) getBunyiBuka() string {
	return "Beep"
}

type pintuKiriType struct{}

func (p pintuKiriType) getBunyi() string {
	return "Beep"
}

func (p pintuKiriType) getBunyiBuka() string {
	return "Knock"
}

type kendaraan struct {
	roda  []roda
	pintu pintu
}

func (k kendaraan) getRoda() string {
	return k.roda[0].getRoda()
}

func (k kendaraan) getPintu() string {
	return k.pintu.getBunyi()
}

func (k kendaraan) getPintuBuka() string {
	return k.pintu.getBunyiBuka()
}

func main() {
	karet := rodaKaret{}
	kayu := rodaKayu{}
	besi := rodaBesi{}

	mobil := kendaraan{
		roda: []roda{karet, kayu, besi, karet},
	}

	pintuKanan := pintuKananType{}
	pintuKiri := pintuKiriType{}

	mobil.pintu = pintuKanan
	fmt.Println("Jenis Roda:", mobil.getRoda())
	fmt.Println("Bunyi Pintu:", mobil.getPintu())
	fmt.Println("Bunyi Pintu Buka:", mobil.getPintuBuka())

	mobil.pintu = pintuKiri
	fmt.Println("Bunyi Pintu:", mobil.getPintu())
	fmt.Println("Bunyi Pintu Buka:", mobil.getPintuBuka())
}
