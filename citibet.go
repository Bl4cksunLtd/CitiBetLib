package citibetlib

import (
	"time"
	"net"
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
	"math/rand"
)

var		(
)

const	(
	
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
	
	return c,nil
}

func	(c *Client)Request(url string, v interface{}) error {

// params are included in the url


	
	resp, err := c.HttpClient.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	} else {
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
	}

	return nil
}