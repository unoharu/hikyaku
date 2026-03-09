package edo

import "math/rand"

type Weather struct {
	Label string
	Line  string
}

var weathers = []Weather{
	{"☀️ 晴れ", "追い風だ！気分が上がるぜ！"},
	{"🌧️ 雨", "足元がぬかるんでいやがる…"},
	{"⛈️ 嵐", "台風だー！荷が濡れちまう！"},
}

// RandomWeather returns a random weather message for the current run.
func RandomWeather() Weather {
	return weathers[rand.Intn(len(weathers))]
}
