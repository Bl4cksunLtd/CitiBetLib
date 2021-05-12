package citibetlib

import (
//	"time"
//	"net"
	"errors"
	"encoding/json"
	"log"
	"fmt"
	"math/rand"
	"io/ioutil"
	"strings"
	"strconv"
)

func (cf *CBfloat64) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		  return json.Unmarshal(b, (*float64)(cf))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		  return err
	}
	fl, err := strconv.ParseFloat(s,64)
	if err != nil {
		  return err
	}
	*cf = CBfloat64(fl)
	return nil
}


// client.Login() returns a responsestatus as returned by citibetlib
// if the request fails then the response has both fields set to 0
// debug mode will cause a fatal error on http errors
func	(c 	*Client)Login()	(ResponseStatus,error)	{
	url:=fmt.Sprintf("%sapi/service/login?api=%s&uid=%s&x=%.16f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rand.Float64())
	if c.config.Info	{
		log.Printf("(Login) Url:%s\n",url)
	}
	rs:=ResponseStatus{}
	err:=c.Request(url,&rs)
	if err!=nil	{
		if c.config.Info	{
			log.Println("(Login) Request failed: ",err)
		}
		if c.config.Debug	{
			log.Fatal("(Login) Request failed: ",err)
		}
		return rs,errors.New(fmt.Sprintf("(Login) http.Request failed"))
	}
	if	rs.Status==0	{
		if c.config.Info	{
			log.Println("(Login) Request failed StatusCode: ",rs.Code)
		}
		return rs,errors.New(fmt.Sprintf("(Login) Request failed : %04d",rs.Code))
	}
	return rs,nil
}

func	(c	*Client)CardList()	(clr	CardListResponse,err	error)	{
	url:=c.config.Url+"api/service/raceinfoservlet?method=cardlist"
	if c.config.Info	{
		log.Printf("(CardList) Url:%s\n",url)
	}
	err=c.Request(url,&clr)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(Cardlist) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(Cardlist) Request failed: ",err)
		}
		return clr,errors.New(fmt.Sprintf("(CardList) http.Request failed"))
	}
	if c.config.Info	{
		log.Println("(CardList) Request returned : ",clr)
	}
	return
}

func	(c	*Client)EventList(rd string,cId	int)	(elr	EventListResponse,err error)	{
	url:=fmt.Sprintf("%sapi/service/raceinfoservlet?method=eventlist&raceDate=%s&cardId=%d",
					c.config.Url,
					rd,
					cId)
					
	if c.config.Info	{
		log.Printf("(EventList) Url:%s\n",url)
	}
	err=c.Request(url,&elr)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(EventList) Request failed: ",err)
		}	
		if c.config.Info	{
			log.Println("(EventList) Request failed: ",err)
		}
		return elr,errors.New("(CardList) http.Request failed")
	}
	if c.config.Info	{
		log.Println("(EventList) Request returned : ",elr)
	}
	return
}

func	(c	*Client)RunnerList(rd string,cId	int,r	int)	(rlr	RunnerListResponse,err error)	{

	url:=fmt.Sprintf("%sapi/service/raceinfoservlet?method=runnerlist&raceDate=%s&cardId=%d&race=%d",
					c.config.Url,
					rd,
					cId,
					r)
					
	if c.config.Info	{
		log.Printf("(RunnerList) Url:%s\n",url)
	}
	err=c.Request(url,&rlr)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(RunnerList) Request failed: ",err)
		}	
		if c.config.Info	{
			log.Println("(RunnerList) Request failed: ",err)
		}
		return rlr,errors.New("(RunnerList) http.Request failed")
	}
	if c.config.Info	{
		log.Println("(RunnerList) Request returned : ",rlr)
	}
	return
}

func	(c	*Client)BetPendingList(rd string,rt	string,r	int,cur 	int,inplay	bool)	(bpl	[]Pending,err error)	{
	inplaystr:=""
	if inplay	{
		inplaystr="&lu=1"
	}

	url:=fmt.Sprintf("%sapi/service/betdata?api=%s&uid=%s&race_date=%s&race_type=%s&rc=%d&c=%d&m=SG%s",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rd,
					rt,
					r,
					cur,
					inplaystr)
					
	if c.config.Info	{
		log.Printf("(BetPendingList) Url:%s\n",url)
	}
	resp, err := c.HttpClient.Get(url)
	if err != nil {
		if c.config.Debug	{
			log.Fatal("(BetPendingList) Get failed: ",err)
		}
		if c.config.Info	{
			log.Println("(BetPendingList) Get failed: ",err)
		}
		return bpl,errors.New("(BetPendingList) http.Request failed")
	}


	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil 	{
		if c.config.Debug	{
			log.Fatal("(BetPendingList) ReadAll failed: ",err)
		}	
		if c.config.Info	{
			log.Println("(BetPendingList) ReadAll failed: ",err)
		}
		return bpl,errors.New("(BetPendingList) http.ReadAll failed")
	}
	if len(data)==0		{
		if c.config.Info	{
			log.Println("(BetPendingList) Get Returned Null: URL: ",url)
		}
		return			// returns no error -> empty structure
	}
	
	if resp.StatusCode != 200 {
		return bpl,errors.New(resp.Status)
	}
	
	if c.config.Info	{
		log.Println("(BetPendingList) Request returned : ",data)
	}
	
	lines:=strings.Split(string(data),"\n")
	if len(lines)==0	{
		if c.config.Info	{
			log.Println("(BetPendingList) zero lines: URL: ",url)
		}
		return			// returns no error -> empty structure
	}
	var		ri,wt,pt				int
	var		e1,e2,e3,e4,e5			error
	var		tp						float64
	for i:=0;i<len(lines);i++	{
		if len(lines[i])==0	{
			continue
		}
		field:=strings.Split(lines[i],"\t")
		if len(field)!=6	{
			if c.config.Info	{
				log.Println("(BetPendingList) Invalid field length: ",lines[i]," from :", url)
			}
			continue
		}
		ri,e1=strconv.Atoi(field[0])
//		hi,e2=strconv.Atoi(field[1])	// horse number is a string as it could be values like 1a
		wt,e3=strconv.Atoi(field[2])
		pt,e4=strconv.Atoi(field[3])
		tp,e5=strconv.ParseFloat(field[4],64)
		if e1!=nil || e2!=nil || e3!=nil || e4!=nil || e5!=nil {
			if c.config.Debug	{
				log.Fatal("(BetPendingList) Error occured converting ",lines[i]," into numbers.")
			}
			if c.config.Info	{
				log.Println("(BetPendingList) Error occured converting ",lines[i]," into numbers.")
			}
			return bpl,errors.New("(BetPendingList) Data is corrupt")
		}
		bpl=append(bpl,Pending{
			Race:			ri,
			Horse:			field[1],
			WinTickets: 	wt,
			PlaceTickets:	pt,
			TicketPrice:	tp,
			Limits:			field[5]})
	}
	
	return
}

func	(c	*Client)EatPendingList(rd string,rt	string,r	int,cur 	int,inplay	bool)	(epl	[]Pending,err error)	{
	inplaystr:=""
	if inplay	{
		inplaystr="&lu=1"
	}

	url:=fmt.Sprintf("%sapi/service/eatdata?api=%s&uid=%s&race_date=%s&race_type=%s&rc=%d&c=%d&m=SG%s",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rd,
					rt,
					r,
					cur,
					inplaystr)
					
	if c.config.Info	{
		log.Printf("(EatPendingList) Url:%s\n",url)
	}
	resp, err := c.HttpClient.Get(url)

	if err != nil {
		if c.config.Debug	{
			log.Fatal("(EatPendingList) Get failed: ",err)
		}
		if c.config.Info	{
			log.Println("(EatPendingList) Get failed: ",err)
		}
		return epl,errors.New("(EatPendingList) http.Request failed")
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.config.Debug	{
			log.Fatal("(EatPendingList) ReadAll failed: ",err)
		}	
		if c.config.Info	{
			log.Println("(EatPendingList) ReadAll failed: ",err)
		}
		return epl,errors.New("(EatPendingList) http.ReadAll failed")
	}
	if len(data)==0		{
		if c.config.Info	{
			log.Println("(EatPendingList) Get Returned Null: URL: ",url)
		}
		return			// returns no error -> empty structure
	}
	
	if resp.StatusCode != 200 {
		return epl,errors.New(resp.Status)
	}
	
	if c.config.Info	{
		log.Println("(EatPendingList) Request returned : ",data)
	}
	
	lines:=strings.Split(string(data),"\n")
	if len(lines)==0	{
		if c.config.Info	{
			log.Println("(EatPendingList) zero lines: URL: ",url)
		}
		return			// returns no error -> empty structure
	}
	
	var		ri,wt,pt				int
	var		e1,e2,e3,e4,e5			error
	var		tp						float64
	for i:=0;i<len(lines);i++	{
		if len(lines[i])==0	{
			continue
		}
		field:=strings.Split(lines[i],"\t")
		if len(field)!=6	{
			if c.config.Info	{
				log.Println("(EatPendingList) Invalid field length: ",lines[i]," from :", url)
			}
			continue
		}
		ri,e1=strconv.Atoi(field[0])
//		hi,e2=strconv.Atoi(field[1])
		wt,e3=strconv.Atoi(field[2])
		pt,e4=strconv.Atoi(field[3])
		tp,e5=strconv.ParseFloat(field[4],64)
		if e1!=nil || e2!=nil || e3!=nil || e4!=nil || e5!=nil {
			if c.config.Debug	{
				log.Fatal("(EatPendingList) Error occured converting ",lines[i]," into numbers.")
			}
			if c.config.Info	{
				log.Println("(EatPendingList) Error occured converting ",lines[i]," into numbers.")
			}
			return epl,errors.New("(EatPendingList) Data is corrupt")
		}
		epl=append(epl,Pending{
			Race:			ri,
			Horse:			field[1],
			WinTickets: 	wt,
			PlaceTickets:	pt,
			TicketPrice:	tp,
			Limits:			field[5]})
	}
	
	return
}


func	(c	*Client)MainInfo()	(maininfo MainInfo,err error)	{
	url:=fmt.Sprintf("%sapi/service/datastore?api=%s&uid=%s&x=%0.9f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rand.Float64())
					
	if c.config.Info	{
		log.Printf("(MainInfo) Url:%s\n",url)
	}
	err=c.Request(url,&maininfo)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(MainInfo) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(MainInfo) Request failed: ",err)
		}
		return maininfo,errors.New("(MainInfo) http.Request failed")
	}
	if c.config.Info	{
		log.Println("(MainInfo) Request returned : ",maininfo)
	}
	return
}

func	(c	*Client)TransActionDetails(racedate string,racetype string,race int)	(tad TransactionDetails,err error)	{
	url:=fmt.Sprintf("%sapi/service/transactionsdetails?api=%s&uid=%s&race_date=%s&race_type=%s&race=%d&rd=%0.9f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					racedate,
					racetype,
					race,
					rand.Float64())
					
	if c.config.Info	{
		log.Printf("(TransActionDetails) Url:%s\n",url)
	}
	err=c.Request(url,&tad)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(TransActionDetails) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(TransActionDetails) Request failed: ",err)
		}
		return tad,errors.New("(TransActionDetails) http.Request failed")
	}
	if c.config.Info	{
		log.Println("(TransActionDetails) Request returned : ",tad)
	}
	return
}


func	(c	*Client)Transactions(racedate string,racetype string,race int)	(tad []Transaction,err error)	{
	url:=fmt.Sprintf("%sapi/service/transactions?api=%s&uid=%s&type=query&race_date=%s&race_type=%s&race=%d&rd=%0.9f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					racedate,
					racetype,
					race,
					rand.Float64())
					
	if c.config.Info	{
		log.Printf("(TransActions) Url:%s\n",url)
	}
	err=c.Request(url,&tad)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(TransActions) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(TransActions) Request failed: ",err)
		}
		return tad,errors.New("(TransActions) http.Request failed")
	}
	if c.config.Info	{
		log.Println("(TransActions) Request returned : ",tad)
	}
	return
}

func	(c	*Client)SubmitBet(racedate string,racetype string,race int,horse string,win int,place int,
									amount float64,wl float64,pl float64,
									wtck	int,ptck	int,live bool)	(bs BetResponse,err error)	{
	var	livestr		string
	if live	{
		livestr=fmt.Sprintf("&show=%d&lu=1",race)
	}
	url:=fmt.Sprintf("%sapi/service/bets?api=%s&uid=%s&race_date=%s&race_type=%s&race=%d&horse=%s&win=%d&place=%d&amount=%.2f"+
						"&l_win=%.1f&l_place=%.1f&wtck=%d&ptck=%d%s&rd=%0.9f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					racedate,
					racetype,
					race,
					horse,
					win,
					place,
					amount,
					wl,
					pl,
					wtck,
					ptck,
					livestr,
					rand.Float64())
					
	if c.config.Info	{
		log.Printf("(SubmitBet) Url:%s\n",url)
	}
	if !c.config.Bet	{
		log.Printf("(SubmitBet) Url:%s\n",url)
		return
	}
	err=c.RequestDebug(url,&bs)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(SubmitBet) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(SubmitBet) Request failed: ",err)
		}
		return bs,errors.New("(SubmitBet) http.Request failed")
	} 
	if c.config.Info	{
		log.Println("(SubmitBet) Request returned : ",bs)
	}
	if len(bs.Bid)>0	{
		b,h,a,lw,lp,w,p:=ParseBid(bs.Bid[0])
		if c.config.Info	{
			log.Println("(SubmitBet.ParseBid) : ",b,h,a,lw,lp,w,p)
		}
		bs.Win=w
		bs.Place=p
		bs.BetId=b
		bs.Pending=1
	}
	if len(bs.Transacted)>0	{
		ht,wt,pt:=ParseTransacted(bs.Transacted[0])
		if c.config.Info	{
			log.Println("(SubmitBet.ParseTransacted) : ",ht,wt,pt)
		}
	}
	return
}

func	(c	*Client)SubmitBetRequest(racedate string,racetype string,race int,horse string,win int,place int,
									amount float64,wl float64,pl float64,
									wtck	int,ptck	int)	(bs BetResponse,err error)	{
	return c.SubmitBet(racedate,racetype,race,horse,win,place,amount,wl,pl,wtck,ptck,false)
}

func	(c	*Client)SubmitLiveBetRequest(racedate string,racetype string,race int,horse string,win int,place int,
									amount float64,wl float64,pl float64,
									wtck	int,ptck	int)	(bs BetResponse,err error)	{
	return c.SubmitBet(racedate,racetype,race,horse,win,place,amount,wl,pl,wtck,ptck,true)
}

func	(c	*Client)SubmitEat(racedate string,racetype string,race int,horse string,win int,place int,
									amount float64,wl float64,pl float64,
									wtck	int,ptck	int,live bool)	(bs BetResponse,err error)	{
	var	livestr		string
	if live	{
		livestr=fmt.Sprintf("&show=%d&lu=1",race)
	}
	url:=fmt.Sprintf("%sapi/service/bookings?api=%s&uid=%s&race_date=%s&race_type=%s&race=%d&horse=%s&win=%d&place=%d&amount=%.2f"+
						"&l_win=%.1f&l_place=%.1f&wtck=%d&ptck=%d%s&rd=%0.9f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					racedate,
					racetype,
					race,
					horse,
					win,
					place,
					amount,
					wl,
					pl,
					wtck,
					ptck,
					livestr,
					rand.Float64())
					
	if c.config.Info	{
		log.Printf("(SubmitEat) Url:%s\n",url)
	}
	if !c.config.Bet	{
		log.Printf("(SubmitEat) Url:%s\n",url)
		return
	}
	err=c.RequestDebug(url,&bs)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(SubmitEat) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(SubmitEat) Request failed: ",err)
		}
		return bs,errors.New("(SubmitEat) http.Request failed")
	}  
	if c.config.Info	{
		log.Println("(SubmitEat) Request returned : ",bs)
	}
	if len(bs.Bid)>0	{
		b,h,a,lw,lp,w,p:=ParseBid(bs.Bid[0])
		if c.config.Info	{
			log.Println("(SubmitEat.ParseBid) : ",b,h,a,lw,lp,w,p)
		}
		bs.Win=w
		bs.Place=p
		bs.BetId=b
		bs.Pending=1
	}
	if len(bs.Transacted)>0	{
		ht,wt,pt:=ParseTransacted(bs.Transacted[0])
		if c.config.Info	{
			log.Println("(SubmitEat.ParseTransacted) : ",ht,wt,pt)
		}
	}
	
	return
}

func	ParseBid(bidstr string)	(bid int64,horse string,amount float64,lwin,lplace float64,win,place float64)	{
	if len(bidstr)==0	{
		return
	}
	str:=strings.ReplaceAll(bidstr,"[","")
	str=strings.ReplaceAll(str,"]","")
	
	bids := strings.Split(str, "_")
	if len(bids)<7	{
		log.Println("(ParseBid) Invalid string len : ",bids)
		return
	}
//	log.Println("(ParseBid) Bids : ",bids)
	
	bid,_=strconv.ParseInt(bids[0],10,64)
//	log.Printf("(ParseBid) {%s} = %d : %v\n",bids[0],bid,err)
	horse=bids[1]
	amount,_=strconv.ParseFloat(bids[2],64)
	lwin,_=strconv.ParseFloat(bids[3],64)
	lplace,_=strconv.ParseFloat(bids[4],64)
	win,_=strconv.ParseFloat(bids[5],64)
	place,_=strconv.ParseFloat(bids[6],64)
//	log.Println("(ParseBid) ",str," = ",bid,horse,amount,lwin,lplace,win,place)
	return
}

func ParseTransacted(str string)	(horse string,win,place float64)	{
	if len(str)==0	{
		return
	}
	trans:= strings.Split(str, "_")
	if len(trans)<3	{
		log.Println("(ParseTransacted) Invalid string len : ",trans)
		return
	}
//	log.Println("(ParseTransacted) Trans: ",trans)
	horse=trans[0]
	win,_=strconv.ParseFloat(trans[1],64)
	place,_=strconv.ParseFloat(trans[2],64)
//	log.Printf("(ParseTransacted) %s = %s, %.2f, %.2f\n",str,horse,win,place)
	return
}
	
	


func	(c	*Client)SubmitEatRequest(racedate string,racetype string,race int,horse string,win int,place int,
									amount float64,lwin float64,pwin float64,
									wtck	int,ptck	int)	(bs BetResponse,err error)	{
	return c.SubmitEat(racedate,racetype,race,horse,win,place,amount,lwin,pwin,wtck,ptck,false)
}

func	(c	*Client)SubmitLiveEatRequest(racedate string,racetype string,race int,horse string,win int,place int,
									amount float64,lwin float64,pwin float64,
									wtck	int,ptck	int)	(bs BetResponse,err error)	{
	return c.SubmitEat(racedate,racetype,race,horse,win,place,amount,lwin,pwin,wtck,ptck,true)
}


func	(c	*Client)DeleteBet(racedate string,racetype string,race int,bid int64,bettype string,x int)	(rs ResponseStatus,err error)	{
	url:=fmt.Sprintf("%sapi/service/transactions?api=%s&uid=%s&race_date=%s&race_type=%s&type=del&race=%d&bid=%d"+
					"&betType=%s&x=%d&show=%d&rd=%0.9f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					racedate,
					racetype,
					race,
					bid,
					bettype,
					x,
					race,
					rand.Float64())
					
	if c.config.Info	{
		log.Printf("(DeleteBet) Url:%s\n",url)
	}
	err=c.Request(url,&rs)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(DeleteBet) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(DeleteBet) Request failed: ",err)
		}
		return rs,errors.New("(DeleteBet) http.Request failed")
	} 
	if c.config.Info	{
		log.Println("(DeleteBet) Request returned : ",rs)
	}
	return
}


//DeleteBet(racedate string,racetype string,race int,bid int64,bettype string,x int)	(rs ResponseStatus,err error)	{
func	(c	*Client)DeletePendingBet(racedate string,racetype string,race int,bid int64)	(rs ResponseStatus,err error)	{
	return c.DeleteBet(racedate,racetype,race,bid,"",1)
}

func	(c	*Client)DeletePendingEat(racedate string,racetype string,race int,bid int64)	(rs ResponseStatus,err error)	{
	return c.DeleteBet(racedate,racetype,race,bid,"",2)
}

func	(c	*Client)DeleteAllPendingBet(racedate string,racetype string,race int,live bool)	(rs ResponseStatus,err error)	{
	if !live {
		return c.DeleteBet(racedate,racetype,race,int64(race),racetype,10)
	}	else	{
		return c.DeleteBet(racedate,racetype,race,int64(race),racetype,14)
	}
	
}

func	(c	*Client)DeleteAllPendingEat(racedate string,racetype string,race int,live bool)	(rs ResponseStatus,err error)	{
	if !live	{
		return c.DeleteBet(racedate,racetype,race,int64(race),racetype,5)
	}	else	{
		return c.DeleteBet(racedate,racetype,race,int64(race),racetype,15)
	}
}

func	(c	*Client)News(cardId	int)	(news	[]string,err error)	{
	var	newsJSON		News
	url:=fmt.Sprintf("%sapi/service/news?api=%s&uid=%s&location=%d",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					cardId)
					
	if c.config.Info	{
		log.Printf("(News) Url:%s\n",url)
	}
	err=c.Request(url,&newsJSON)
	if err!=nil	{
		if c.config.Debug	{
			log.Fatal("(News) Request failed: ",err)
		}
		if c.config.Info	{
			log.Println("(News) Request failed: ",err)
		}
		return news,errors.New("(News) http.Request failed")
	}
	if c.config.Info	{
		log.Println("(News) Request returned : ",news)
	}
	news=append(news,newsJSON.News1,newsJSON.News2,newsJSON.News3)
	return
}