package citibetlib

import (
	"time"
	"net"
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"math/rand"
	"log"
)

var		(
)

const	(
	version	=	"1.1c-beta"
)

func	Version()		string	{
	return	string(version)
}

func	NewClient(config	*Config)	(*Client,error)	{
	if config.UserName==""	{
		return nil,errors.New("NewClient: Config missing username")
	}
	if config.ApiKey==""	{
		return nil,errors.New("NewClient: Config missing ApiKey")
	}
	if config.Url==""	{
		return nil,errors.New("NewClient: Config missing Url")
	}

	if config.Timeout==0	{
		config.Timeout=2
	}

	rand.Seed(time.Now().UnixNano())

	c:=new(Client)
	
	c.config = config
	
	netTransport:=&http.Transport{
		Dial:	(&net.Dialer{
					Timeout: time.Duration(config.Timeout) * time.Second,
				}).Dial,
				TLSHandshakeTimeout: time.Duration(config.Timeout)*time.Second,
	}
	c.HttpClient=&http.Client{
		Timeout: 	time.Duration(config.Timeout)*time.Second,
		Transport:	netTransport,
	}
	return c,nil
}

func	(c *Client)Request(url string, v interface{}) error {

// params are included in the url

	if c.config.Info	{
		log.Println("(Request) Url: ",url)
	}

	
	resp, err := c.HttpClient.Get(url)

	if err != nil {
		if c.config.Info	{
			log.Println("(Request) Get failed: ",err)
		}
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.config.Info	{
			log.Println("(Request) ReadAll failed: ",err," URL: ",url)
		}
		return err
	}
	if len(data)==0	|| data[0]==10	{
		if c.config.Info	{
			log.Println("(Request) Returned Null: URL: ",url,data)
		}
		return	nil		// returns no error -> empty structure
	}
	FixJSON(data,len(data)) // deal with leading zeros
	if resp.StatusCode != 200 {
		if c.config.Info	{
			log.Println("(Request) StatusCode not 200: ",resp.StatusCode," Status: ",resp.Status)
		}
		return errors.New(resp.Status)
	} else {
		if err := json.Unmarshal(data, v); err != nil {
			if c.config.Info	{
				log.Println("(Request) Unmarshal failed: ",err," raw data: ",data)
			}
			return err
		}
	}

	return nil
}


func	(c *Client)RequestDebug(url string, v interface{}) error {

// params are included in the url

	if c.config.Info	{
		log.Println("(Request) Url: ",url)
	}

	
	resp, err := c.HttpClient.Get(url)

	if err != nil {
		if c.config.Info	{
			log.Println("(Request) Get failed: ",err," URL: ",url)
		}
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.config.Info	{
			log.Println("(Request) ReadAll failed: ",err," URL: ",url)
		}
		return err
	}
	if len(data)==0	|| data[0]==10	{
		if c.config.Info	{
			log.Println("(Request) Returned Null: URL: ",url,data)
		}
		log.Println("(Request) Returned Null: URL: ",url,data)
		return	nil		// returns no error -> empty structure
	}
	FixJSON(data,len(data)) // deal with leading zeros
	if resp.StatusCode != 200 {
		if c.config.Info	{
			log.Println("(Request) StatusCode not 200: ",resp.StatusCode," Status: ",resp.Status)
		}
		return errors.New(resp.Status)
	} else {
		log.Println("(Request) Data: ",data)
		if err := json.Unmarshal(data, v); err != nil {
//			if c.config.Info	{
				log.Println("(Request) Unmarshal failed: ",err," raw data: ",data)
//			}
			return err
		}
	}

	return nil
}


func	FixJSON(bad	[]byte,ln int)	{
//	log.Println("(FixJSON) Before: ",bad)
	
	var	quote		bool
	var	m				int
	for n:=0;n<ln;n++	{
		if bad[n]=='\t'	{
			bad[n]=' '
		}
		if bad[n]=='"'	{
			quote=!quote
			bad[m]=bad[n]
			m++
			continue
		}
		if	quote	{
			bad[m]=bad[n]
			m++
			continue
		}

		if m>=1 && n<ln-1	&& bytes.Contains([]byte(" :,}-"),[]byte{bad[m-1]}) && bad[n]=='0' && bytes.Contains([]byte("0123456789"),[]byte{bad[n+1]})	{
			continue
		}
		bad[m]=bad[n]
		m++
	}
	for n:=m;n<ln;n++	{
		bad[n]=' '
	}
//	log.Println("(FixJSON) After: ",bad)
}