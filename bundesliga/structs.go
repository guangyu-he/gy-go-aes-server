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

type NextMatch struct {
	ID       int `json:"id"`
	HomeTeam struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Power int    // value used for evaluation
	} `json:"homeTeam"`
	AwayTeam struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Power int
	} `json:"awayTeam"`
	UtcDate    string `json:"utcDate"`
	MatchDay   int    `json:"matchday"`
	Prediction struct {
		HomeTeam int
		AwayTeam int
	}
}

type Response struct {
	Matches []Match `json:"matches"`
}

type ResponseNextMatch struct {
	Matches []NextMatch `json:"matches"`
}

type Competition struct {
	CurrentSeason struct {
		CurrentMatchday int `json:"currentMatchday"`
	} `json:"currentSeason"`
}
