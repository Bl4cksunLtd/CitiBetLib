package citibetlib

import (
//	"time"
//	"net"
	"errors"
//	"encoding/json"
	"log"
	"fmt"
	"math/rand"
	"io/ioutil"
	"strings"
	"strconv"
)

func	(c 	*Client)Login()	(ResponseStatus,error)	{
	url:=fmt.Sprintf("%sapi/service/login?api=%s&uid=%s&x=%.16f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rand.Float64())
	if c.config.Info	{
		log.Printf("(Login) Url:%s\n",url)
	}
	if c.config.Debug	{
		return	ResponseStatus{1,0},nil
	}	
	
	rs:=ResponseStatus{}
	err:=c.Request(url,&rs)
	if err!=nil	{
		log.Fatal("(Login) Request failed: ",err)
	}
	if	rs.Status==0	{
		if c.config.Info	{
			log.Println("(Login) Request failed StatusCode: ",rs.Code)
		}
		return rs,errors.New(fmt.Sprintf("(Login) Request failed StatusCode: %04d",rs.Code))
	}
	return rs,nil
}

func	(c	*Client)CardList()	(clr	CardListResponse,err	error)	{
	url:=c.config.Url+"api/service/raceinfoservlet?method=cardlist"
	if c.config.Info	{
		log.Printf("(CardList) Url:%s\n",url)
	}
	if c.config.Debug	{
		return
	}	
	err=c.Request(url,&clr)
	if err!=nil	{
		log.Fatal("(Cardlist) Request failed: ",err)
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
	if c.config.Debug	{
		return	
	}	
	err=c.Request(url,&elr)
	if err!=nil	{
		log.Fatal("(EventList) Request failed: ",err)
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
	if c.config.Debug	{
		return	
	}	
	err=c.Request(url,&rlr)
	if err!=nil	{
		log.Fatal("(RunnerList) Request failed: ",err)
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
	if c.config.Debug	{
		return	
	}	
	resp, err := c.HttpClient.Get(url)

	if err != nil {
		log.Fatal("(BetPendingList) Get failed: ",err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("(BetPendingList) ReadAll failed: ",err)
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
			log.Fatal("(BetPendingList) Error occured converting ",lines[i]," into numbers.")
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
	if c.config.Debug	{
		return	
	}	
	resp, err := c.HttpClient.Get(url)

	if err != nil {
		log.Fatal("(EatPendingList) Get failed: ",err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("(EatPendingList) ReadAll failed: ",err)
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
			log.Fatal("(BetPendingList) Error occured converting ",lines[i]," into numbers.")
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
	if c.config.Debug	{
		return	
	}	
	err=c.Request(url,&newsJSON)
	if err!=nil	{
		log.Fatal("(News) Request failed: ",err)
	}
	if c.config.Info	{
		log.Println("(News) Request returned : ",news)
	}
	news=append(news,newsJSON.News1,newsJSON.News2,newsJSON.News3)
	return
}