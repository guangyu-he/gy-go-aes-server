package bundesliga

type Match struct {
	ID       int `json:"id"`
	HomeTeam struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"homeTeam"`
	AwayTeam struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"awayTeam"`
	UtcDate  string `json:"utcDate"`
	Status   string `json:"status"`
	MatchDay int    `json:"matchday"`
	Score    struct {
		FullTime struct {
			HomeTeam int `json:"home"`
			AwayTeam int `json:"away"`
		} `json:"fullTime"`
	} `json:"score"`
	Winner struct {
		Name string ""
		ID   int    ""
	}
	Draw bool
}

type Response struct {
	Matches []Match `json:"matches"`
}

type Competition struct {
	CurrentSeason struct {
		CurrentMatchday int `json:"currentMatchday"`
	} `json:"currentSeason"`
}
