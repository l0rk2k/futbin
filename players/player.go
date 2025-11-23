package players

// Role representa un rol específico con su nivel
type Role struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
}

// RolePosition representa roles disponibles para una posición
type RolePosition struct {
	Position string `json:"position"`
	Roles    []Role `json:"roles"`
}

type Player struct {
	ID         int `json:"ID"`
	PlayerID   int `json:"playerid"`
	ResourceID int `json:"resource_id"`
	Image      int `json:"playerimage"`

	Name       string `json:"playername"`
	CommonName string `json:"common_name"`

	// Posiciones
	MainPosition         string   `json:"position"`
	Position2            string   `json:"position2,omitempty"`
	Position3            string   `json:"position3,omitempty"`
	Position4            string   `json:"position4,omitempty"`
	AlternativePositions []string `json:"alternativePositions,omitempty"`
	Positions            string   `json:"pos_all"`
	Accelerate           string   `json:"accelerate"`

	// Imágenes
	PlayerImage        interface{} `json:"player_image,omitempty"` // Puede ser string o int
	SpecialImage       int         `json:"Special_Image,omitempty"`
	NormalPlayerImage  string      `json:"normalPlayerImage,omitempty"`
	SpecialPlayerImage string      `json:"specialPlayerImage,omitempty"`
	CardImage          string      `json:"cardImage,omitempty"`

	// Ratings y tipo
	Rating       int    `json:"rating"`
	RareType     int    `json:"raretype"`
	RareTypeName string `json:"rareTypeName,omitempty"`

	// Estadísticas principales
	Pace        int `json:"pac"`
	Shooting    int `json:"sho"`
	Passing     int `json:"pas"`
	Dribbling   int `json:"dri"`
	Defending   int `json:"def"`
	Physicality int `json:"phy"`
	TotalIGS    int `json:"TotalIGS,omitempty"`

	// Club, Nación, Liga
	ClubID   int    `json:"club"`
	Club     string `json:"club_name"`
	Nation   string `json:"nation_name"`
	NationID int    `json:"nation"`
	LeagueID int    `json:"league"`
	League   string `json:"league_name"`

	// Características del jugador
	Skills            int            `json:"skills,omitempty"`
	WeakFoot          int            `json:"Weak_Foot,omitempty"`
	PlayerFoot        string         `json:"Player_Foot,omitempty"`
	DefensiveWorkrate string         `json:"Defensive_Workrate,omitempty"`
	AttackWorkrate    string         `json:"Attack_Workrate,omitempty"`
	PlayerHeight      string         `json:"Player_Height,omitempty"`
	PlayerWeight      int            `json:"Player_Weight,omitempty"`
	BodyTypeCode      int            `json:"bodytypecode,omitempty"`
	Likes             int            `json:"likes,omitempty"`
	Playstyles        *string        `json:"playstyles,omitempty"`
	PlaystylesPlus    *string        `json:"playstylesPlus,omitempty"`
	RolesPerPosition  []RolePosition `json:"rolesPerPosition,omitempty"`

	// Precios PlayStation (pueden ser null)
	PsPrice      *int     `json:"ps_LCPrice"`
	PsPRP        *int     `json:"ps_PRP"`
	PsMinPrice   *int     `json:"ps_MinPrice"`
	PsMaxPrice   *int     `json:"ps_MaxPrice"`
	PsPriceTrend *float64 `json:"ps_PriceTrend,omitempty"`
	PsLCPClosing *int     `json:"ps_LCPClosing,omitempty"`

	// Precios PC (pueden ser null)
	PcPrice      *int     `json:"pc_LCPrice"`
	PcPRP        *int     `json:"pc_PRP"`
	PcMinPrice   *int     `json:"pc_MinPrice"`
	PcMaxPrice   *int     `json:"pc_MaxPrice"`
	PcPriceTrend *float64 `json:"pc_PriceTrend,omitempty"`
	PcLCPClosing *int     `json:"pc_LCPClosing,omitempty"`

	// Precios Xbox (pueden ser null)
	XboxPrice      *int     `json:"xbox_LCPrice"`
	XboxPRP        *int     `json:"xbox_PRP"`
	XboxMinPrice   *int     `json:"xbox_MinPrice"`
	XboxMaxPrice   *int     `json:"xbox_MaxPrice"`
	XboxPriceTrend *float64 `json:"xbox_PriceTrend,omitempty"`
}

func (p Player) Price(platform string) int {
	switch platform {
	case "PC":
		if p.PcPrice != nil {
			return *p.PcPrice
		}
	case "XB":
		if p.XboxPrice != nil {
			return *p.XboxPrice
		}
	default:
		if p.PsPrice != nil {
			return *p.PsPrice
		}
	}
	return 0
}

func (p Player) MinPrice(platform string) int {
	switch platform {
	case "PC":
		if p.PcMinPrice != nil {
			return *p.PcMinPrice
		}
	case "XB":
		if p.XboxMinPrice != nil {
			return *p.XboxMinPrice
		}
	default:
		if p.PsMinPrice != nil {
			return *p.PsMinPrice
		}
	}
	return 0
}

func (p Player) MaxPrice(platform string) int {
	switch platform {
	case "PC":
		if p.PcMaxPrice != nil {
			return *p.PcMaxPrice
		}
	case "XB":
		if p.XboxMaxPrice != nil {
			return *p.XboxMaxPrice
		}
	default:
		if p.PsMaxPrice != nil {
			return *p.PsMaxPrice
		}
	}
	return 0
}

func (p Player) PRP(platform string) int {
	switch platform {
	case "PC":
		if p.PcPRP != nil {
			return *p.PcPRP
		}
	case "XB":
		if p.XboxPRP != nil {
			return *p.XboxPRP
		}
	default:
		if p.PsPRP != nil {
			return *p.PsPRP
		}
	}
	return 0
}

type playersData struct {
	Players []Player `json:"data"`
}
