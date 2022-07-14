package nataelb

import (
	"fmt"
	"log"
	"sync"
)

type Candidate struct {
	url             string
	healthcheckPath string
	isAlive         bool
	sync.RWMutex
}

func (c *Candidate) GetHealthcheckURL() string {

	protocol := ""

	// Only IP address can cause urls.parse() error
	if IsIPAddress(c.url) {
		protocol = "http://"
	}

	return fmt.Sprintf("%s%s/%s", protocol, c.url, c.healthcheckPath)
}

func (c *Candidate) GetTargetURL(apiPath string) string {
	protocol := ""

	// Only IP address can cause urls.parse() error
	if IsIPAddress(c.url) {
		protocol = "http://"
	}

	return fmt.Sprintf("%s%s/%s", protocol, c.url, apiPath)
}

func (c *Candidate) IsAlive() bool {
	c.RLock()
	defer c.RUnlock()
	return c.isAlive
}

func (c *Candidate) SetAlive() {
	c.Lock()
	defer c.Unlock()
	c.isAlive = true
	log.Println(fmt.Sprintf("Alive %s/%s", c.url, c.healthcheckPath))
}

func (c *Candidate) SetDead() {
	c.Lock()
	defer c.Unlock()
	c.isAlive = false
	log.Println(fmt.Sprintf("Dead %s/%s", c.url, c.healthcheckPath))
}
