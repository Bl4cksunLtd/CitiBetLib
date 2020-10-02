package citibetlib

import (
	"time"
	"net"
	"net/http"
	"errors"
	"regexp"
	"io/ioutil"
	"encoding/json"
//	"github.com/pquerna/ffjson/ffjson"
	"math/rand"
	"log"
)

var		(
)

const	(
	version	=	"1.0a"
)



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

	rand.Seed(time.Now().UnixNano())

	c:=new(Client)
	
	c.config = config
	
	netTransport:=&http.Transport{
		Dial:	(&net.Dialer{
					Timeout: 2 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 2*time.Second,
	}
	c.HttpClient=&http.Client{
		Timeout: 	2*time.Second,
		Transport:	netTransport,
	}
	c.rx=regexp.MustCompile(`\b0*(\d+)`)
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
	c.rx.ReplaceAll(data, []byte("$1"))
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