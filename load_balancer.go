package nataelb

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

type LoadBalancer struct {
	candidates []*Candidate
	iterator   int
}

func NewLB() *LoadBalancer {
	return &LoadBalancer{}
}

func (lb *LoadBalancer) Start(config *Config) error {
	// Initialize candidates
	lb.candidates = make([]*Candidate, len(config.URLInfos))
	for i, urlInfo := range config.URLInfos {
		lb.candidates[i] = &Candidate{
			url:             urlInfo.URL,
			healthcheckPath: urlInfo.HealthcheckPath,
			isAlive:         false,
		}
	}

	// Start healthcheck
	go lb.healthcheck_goroutine(config.HealthcheckInterval)

	log.Println(fmt.Sprintf("Load balancer started!, port: %d", config.Port))
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", config.Port),
		http.HandlerFunc(lb.handle))

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (lb *LoadBalancer) healthcheck_goroutine(interval int) {
	for {
		for _, candidate := range lb.candidates {
			resopnse, err := http.Get(candidate.GetHealthcheckURL())
			if err != nil {
				log.Println("Healthcheck error:", err)
				candidate.SetDead()
				continue
			}

			if resopnse.StatusCode == 200 {
				candidate.SetAlive()
			} else {
				candidate.SetDead()
			}
			resopnse.Body.Close()
		}

		time.Sleep(time.Second * time.Duration(interval))
	}
}

func (lb *LoadBalancer) handle(w http.ResponseWriter, req *http.Request) {
	// Find target with round robin
	var targetCandidate *Candidate
	for i := 0; i < len(lb.candidates); i++ {
		index := lb.iterator % len(lb.candidates)
		if lb.candidates[index].IsAlive() {
			targetCandidate = lb.candidates[index]
		}
	}
	lb.iterator++

	if targetCandidate == nil {
		log.Println("Not available for all backends")
		w.WriteHeader(503)
		return
	}

	// Forward to targetURL
	reqBodyBuf := make([]byte, req.ContentLength)
	req.Body.Read(reqBodyBuf)
	defer req.Body.Close()

	targetURL := targetCandidate.GetTargetURL(req.URL.Path)
	contentType := req.Header["Content-Type"][0]
	res, err := http.Post(targetURL, contentType, bytes.NewBuffer(reqBodyBuf))
	if err != nil {
		targetCandidate.SetDead()
		w.WriteHeader(500)
		return
	}

	resBodyBuf := make([]byte, res.ContentLength)
	res.Body.Read(resBodyBuf)
	defer res.Body.Close()
	w.WriteHeader(res.StatusCode)
	w.Write(resBodyBuf)
}
