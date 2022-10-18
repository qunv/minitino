package models

type Config struct {
	RootName string
	Intro    string
	FindOn   FindOn
}

type FindOn struct {
	Github  string
	Twitter string
}
