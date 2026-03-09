package ui

type Station struct {
	Percent float64
	Name    string
	Quote   string
}

var stations = []Station{
	{Percent: 0.00, Name: "日本橋", Quote: "さあ出発だ！江戸の都を背に走るぜ！"},
	{Percent: 0.15, Name: "箱根", Quote: "箱根の山越えってところだ。きついぜ！"},
	{Percent: 0.30, Name: "由比", Quote: "駿河の海が見えてきたぜ。潮風が気持ちいいな！"},
	{Percent: 0.50, Name: "浜松", Quote: "ちょうど半分だ。まだまだ走れるぜ！"},
	{Percent: 0.65, Name: "宮", Quote: "尾張まで来たか。七里の渡しを越えてきたぜ！"},
	{Percent: 0.85, Name: "草津", Quote: "もうすぐだ！京の都が待ってるぜ！"},
	{Percent: 1.00, Name: "三条大橋", Quote: "ガッテンだ！無事に荷を届けたぜ。受け取りの判をもらってきな！"},
}

func CurrentStation(percent float64) Station {
	for i := len(stations) - 1; i >= 0; i-- {
		if percent >= stations[i].Percent {
			return stations[i]
		}
	}
	return stations[0]
}

func StartStation() Station {
	return stations[0]
}

func EndStation() Station {
	return stations[len(stations)-1]
}
