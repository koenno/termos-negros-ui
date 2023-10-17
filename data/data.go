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

func (m Menu) GetData() (domain.Menu, error) {
	rateLimiter := rate.NewLimiter(rate.Every(time.Second), 1)
	reqSender := client.New(rateLimiter)
	req, err := portal.NewRequestFactory().NewRequest(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create a request: %v", err)
	}

	bb, headers, err := reqSender.Send(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send a request: %v", err)
	}

	menuParser := parser.NewMenuParser()
	respParser := portal.NewBodyParser(menuParser)
	menu, err := respParser.Parse(bb, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to parse menu: %v", err)
	}

	return menu, nil
}
