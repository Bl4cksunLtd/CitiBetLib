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
	url:=fmt.Sprintf("%sapi/service/login?api=%s&uid=%s&x=.16f",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rand.Float64)
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

func	(c	*Client)BetPendingList(rd string,rt	string,r	int,cur 	int)	(bpl	[]Pending,err error)	{

	url:=fmt.Sprintf("%sapi/service/betdata?api=%s&uid=%s&race_date=%s&race_type=%s&rc=%d&c=%d&m=SG",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rd,
					rt,
					r,
					cur)
					
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

	if resp.StatusCode != 200 {
		return bpl,errors.New(resp.Status)
	}
	
	if c.config.Info	{
		log.Println("(BetPendingList) Request returned : ",data)
	}
	
	lines:=strings.Split(string(data),"\n")
	var		ri,hi,wt,pt,wl,pl	int
	var		tp						float64
	for i:=0;i<len(lines);i++	{
		field:=strings.Split(lines[i],"\t")
		ri,_=strconv.Atoi(field[0])
		hi,_=strconv.Atoi(field[1])
		wt,_=strconv.Atoi(field[2])
		pt,_=strconv.Atoi(field[3])
		tp,_=strconv.ParseFloat(field[4],64)
		limits:=strings.Split(field[5],":")
		wl,_=strconv.Atoi(limits[0])
		pl,_=strconv.Atoi(limits[1])
		bpl=append(bpl,Pending{
			Race:			ri,
			Horse:			hi,
			WinTickets: 	wt,
			PlaceTickets:	pt,
			TicketPrice:	tp,
			WinLimit:		wl,
			PlaceLimit:		pl})
	}
	
	return
}

func	(c	*Client)EatPendingList(rd string,rt	string,r	int,cur 	int)	(epl	[]Pending,err error)	{

	url:=fmt.Sprintf("%sapi/service/eatdata?api=%s&uid=%s&race_date=%s&race_type=%s&rc=%d&c=%d&m=SG",
					c.config.Url,
					c.config.ApiKey,
					c.config.UserName,
					rd,
					rt,
					r,
					cur)
					
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

	if resp.StatusCode != 200 {
		return epl,errors.New(resp.Status)
	}
	
	if c.config.Info	{
		log.Println("(EatPendingList) Request returned : ",data)
	}
	
	lines:=strings.Split(string(data),"\n")
	var		ri,hi,wt,pt,wl,pl	int
	var		tp						float64
	for i:=0;i<len(lines);i++	{
		field:=strings.Split(lines[i],"\t")
		ri,_=strconv.Atoi(field[0])
		hi,_=strconv.Atoi(field[1])
		wt,_=strconv.Atoi(field[2])
		pt,_=strconv.Atoi(field[3])
		tp,_=strconv.ParseFloat(field[4],64)
		limits:=strings.Split(field[5],":")
		wl,_=strconv.Atoi(limits[0])
		pl,_=strconv.Atoi(limits[1])
		epl=append(epl,Pending{
			Race:			ri,
			Horse:			hi,
			WinTickets: 	wt,
			PlaceTickets:	pt,
			TicketPrice:	tp,
			WinLimit:		wl,
			PlaceLimit:		pl})
	}
	
	return
}
