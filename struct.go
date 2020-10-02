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
	CardId		int
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
	RaceStatus		int
	RaceTime		string
	Race			int
}

type	EventListResponse	struct	{
	EventList	[]Event
}

type	Runner	struct	{
	No				int
	JockeyName		string
	TrainerName		string
	HsName			string
	Draw			int
	Wgt				float64
}

type	RunnerListResponse	struct	{
	RunnerList		[]Runner
}

type 	LoginResponse	struct	{
	Status 			int
	Code			string
}

type 	Pending		struct	{
	Race			int
	Horse			int
	WinTickets		int
	PlaceTickets	int
	TicketPrice		float64
	WinLimit		int
	PlaceLimit		int
}

type	ResponseStatus	struct	{
	Status			int
	Code			string
}
