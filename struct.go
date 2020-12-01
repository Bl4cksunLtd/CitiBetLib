package citibetlib

import (
//	"time"
	"net/http"
	"sync"
)

const	(
	
)

type	Config	struct	{
	UserName	string
	ApiKey		string
	Locale		string
	Url			string			//	http://citi-bet-ip/
	Debug		bool			// 	don't send http requests
	Info		bool			//	report Urls used via log
}


type	Client	struct	{
	config		*Config
	once		sync.Once
	HttpClient	*http.Client
}

type	Card	struct	{
	CardId		int				`json:",string"`
	RaceDate	string
	Country		string
	Venue		string
	RaceType	string
	Dividends	string
}

type	CardListResponse	struct	{
	CardList	[]Card
}

type	Event	struct	{
	RaceStatus		int			`json:",string"`
	RaceTime		string
	Race			int			`json:",string"`
}

type	EventListResponse	struct	{
	EventList	[]Event
}

type	Runner	struct	{
	No				string
	JockeyName		string
	TrainerName		string
	HsName			string
	Draw			string		//	`json:",string"`
	Wgt				string		// `json:",string"`
}

type	RunnerListResponse	struct	{
	RunnerList		[]Runner
}

type 	LoginResponse	struct	{
	Status 			int
	Code			int
}

type 	Pending		struct	{
	Race			int
	Horse			string
	WinTickets		int
	PlaceTickets	int
	TicketPrice		float64
	Limits			string
}

type	ResponseStatus	struct	{
	Status			int
	Code			int
}

type	News		struct	{
	News1			string
	News2			string
	News3			string
}

type	MainInfo	struct	{
	Balance			float64		`json:",string"`
	Pls				float64		`json:",string"`
	IsLocked		bool		`json:",string"`
	IsSuspended		bool		`json:",string"`
	BetTixLmt		float64		`json:",string"`
	EatTixLmt		float64		`json:",string"`
	BetTixLmtLive	float64		`json:",string"`
	EatTixLmtLive	float64		`json:",string"`
	QBetTixLmt		float64		`json:",string"`
	QEatTixLmt		float64		`json:",string"`
	BetTixLmtMinor	float64		`json:",string"`
	EatTixLmtMinor	float64		`json:",string"`
	QBetTixLmtMinor	float64		`json:",string"`
	QEatTixLmtMinor	float64		`json:",string"`
	BetTax			float64		`json:",string"`
	EatTax			float64		`json:",string"`
	FcBetTax		float64		`json:",string"`
	FcEatTax		float64		`json:",string"`
	QBetTax			float64		`json:",string"`
	QEatTax			float64		`json:",string"`
	BetTaxLive		float64		`json:",string"`
	EatTaxLive		float64		`json:",string"`
	BetTaxMinor		float64		`json:",string"`
	EatTaxMinor		float64		`json:",string"`
	FcBetTaxMinor	float64		`json:",string"`
	FcEatTaxMinor	float64		`json:",string"`
	QBetTaxMinor	float64		`json:",string"`
	QEatTaxMinor	float64		`json:",string"`
}