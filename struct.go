package citibetlib

import (
//	"time"
	"net/http"
	"sync"
)

const	(
	
)

type	CBfloat64	float64

type	Config	struct	{
	UserName	string
	ApiKey		string
	Locale		string
	Url			string			//	http://citi-bet-ip/
	Debug		bool			// 	don't send http requests
	Info		bool			//	report Urls used via log
	Bet			bool			// only send bets if this is true
	Timeout		int				// 	http client timeout
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

type	TransactionDetail	struct	{
	Horse			int									// horse number
	Horse1			int									// Forecast/Q first horse number.
	Horse2			int									// Forecast/Q second horse number
	GameType		int									//	1:Win/Place 2:Forecast/Place Forecast/Q
	BetType			int									//	0:Bet 1:Booking(Eat)
	Amount			CBfloat64								//	ticket price
	Win				CBfloat64		//`json:",string"`		//	number of win tickets
	Place			CBfloat64		//`json:",string"`		//	number of place tickets
	Live			int			`json:",string"`		//	0:normal betting 1:in-running betting
	Tickets			CBfloat64		//`json:",string"`		//	Forecast/Q tickets
	Limit			string								//	matched payout
	TransType		int									//	0:Forecast/Quinella 1:Place Forecast/Qin-Place
	Tax				CBfloat64		//`json:",string"`		//	bets tax. if scratch, tax will not be displayed
	Pls				CBfloat64		//`json:",string"`		//	profit and loss. if scratch or unsettled, pls will not be displayed
	Status			int									//	0:unsettled 1:settled 2:scratch 4:void FC 5:void PFT 6:void FC&PFT 7:void race 8:void txn
	Tid				int64								// 	Bet Id?
}

type	WP			struct		{
	WpCount				int
	WpTransactions		[]TransactionDetail
}

type	FC			struct	{
	FcCount				int
	FcTransactions		[]TransactionDetail
}

type 	TransactionDetails			struct	{
	WP
	FC
}

type 	Transaction		struct	{
	RaceType		string								//	Race card's ID
	RaceDate		string								//	race date
	Race			int									//	race number
	Horse			int									//	horse number
	Horse1			int									//	Forecast/Q first horse number.
	Horse2			int									//	Forecast/Q second horse number
	GameType		int									//	1:Win/Place 2:Forecast/Place Forecast/Q
	BetType			int									//	0:Bet 1:Booking(Eat)
	Amount			CBfloat64		//`json:",string"`		//	ticket price
	Win				CBfloat64		//`json:",string"`		//	number of win tickets
	Place			CBfloat64		//`json:",string"`		//	number of place tickets
	Live			int				`json:",string"`		//	0:normal betting 1:in-running betting
	Tickets			CBfloat64									//	Forecast/Q tickets
	Limit			string								//	matched payout
	TransType		int									//	0:Forecast/Quinella 1:Place Forecast/Qin-Place
	Pending			int									//	0:transacted 1:pending
	Bid				int									//	pending's bet/booking ID
	Tid				int									// 	pending's bet/booking ID?
}

type	MyTransaction	struct	{
	Transactions 	[]Transaction
}

type 	BetResponse		struct	{
	ResponseStatus
	Race_no				int
	Horse_no			string
	Win					float64
	Place				float64
	BetId				int64
	Pending				int
	Bid				[]string
	Transacted			[]string
}
