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