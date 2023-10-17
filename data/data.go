package data

import (
	"context"
	"fmt"
	"time"

	"github.com/koenno/termos-negros/client"
	"github.com/koenno/termos-negros/client/portal"
	"github.com/koenno/termos-negros/domain"
	"github.com/koenno/termos-negros/parser"
	"golang.org/x/time/rate"
)

type Menu struct {
}

func NewMenu() Menu {
	return Menu{}
}

func (m Menu) GetData() (<-chan domain.DayMenu, <-chan error) {
	dataPipe := make(chan domain.DayMenu)
	errPipe := make(chan error)
	go func() {
		defer close(dataPipe)
		defer close(errPipe)
		m.fetch(dataPipe, errPipe)
	}()
	return dataPipe, errPipe
}

func (m Menu) fetch(dataPipe chan<- domain.DayMenu, errPipe chan<- error) {
	rateLimiter := rate.NewLimiter(rate.Every(time.Second), 1)
	reqSender := client.New(rateLimiter)
	req, err := portal.NewRequestFactory().NewRequest(context.Background())
	if err != nil {
		errPipe <- fmt.Errorf("failed to create a request: %v", err)
		return
	}

	bb, headers, err := reqSender.Send(req)
	if err != nil {
		errPipe <- fmt.Errorf("failed to send a request: %v", err)
		return
	}

	menuParser := parser.NewMenuParser()
	respParser := portal.NewBodyParser(menuParser)
	menu, err := respParser.Parse(bb, headers)
	if err != nil {
		errPipe <- fmt.Errorf("failed to parse menu: %v", err)
		return
	}

	for _, entry := range menu {
		dataPipe <- entry
	}
}
