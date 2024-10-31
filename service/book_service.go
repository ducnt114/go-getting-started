package service

import (
	"context"
	"fmt"
	"github.com/samber/do"
	"net/http"
	"time"
)

type BookService interface {
	GetData(ctx context.Context) error
}

type bookServiceImpl struct {
}

func NewBookService(di *do.Injector) (BookService, error) {
	return &bookServiceImpl{}, nil
}

func (s *bookServiceImpl) GetData(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	//if err := fetData(ctx); err != nil {
	//	return err
	//}

	if err := wrapFunc(ctx); err != nil {
		return err
	}

	return nil
}

func wrapFunc(ctx context.Context) error {
	proc1 := make(chan struct{}, 1)
	go func() {
		// Would not be executed because timeout comes first
		_ = slowFunc(ctx)
		proc1 <- struct{}{}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-proc1:
			fmt.Println("slowFunc done")
			return nil
		}
	}
}

func slowFunc(ctx context.Context) error {
	time.Sleep(10 * time.Second)
	return nil
}

func fetData(ctx context.Context) error {
	// Create a custom HTTP client with a timeout
	client := &http.Client{}

	// Create a new request with the context
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/hello", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Printf("Response status: %s\n", resp.Status)
	return nil
}
