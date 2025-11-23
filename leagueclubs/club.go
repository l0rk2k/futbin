package leagueclubs

type Club struct {
	ID   int    `json:"club_id"`
	Name string `json:"club_name"`
}

type League struct {
	ID    int    `json:"league_id"`
	Name  string `json:"league_name"`
	Clubs []Club `json:"clubs"`
}

type leaguesAndClubs struct {
	Leagues []League `json:"data"`
}
