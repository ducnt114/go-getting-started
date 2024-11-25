package order_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Saga struct {
	steps []SagaStep
}

func (s *Saga) AddStep(step SagaStep) {
	s.steps = append(s.steps, step)
}

func (s *Saga) Execute(data interface{}) error {
	var completedSteps []SagaStep

	for _, step := range s.steps {
		fmt.Println("Executing step:", step.GetName())
		if err := step.Execute(data); err != nil {
			fmt.Println("Error executing step:", err)
			s.Rollback(completedSteps, data)
			return err
		}
		completedSteps = append(completedSteps, step)
	}

	fmt.Println("Saga execution completed successfully.")
	return nil
}

func (s *Saga) Rollback(completedSteps []SagaStep, data interface{}) {
	fmt.Println("Rollback for failed saga...")
	for i := len(completedSteps) - 1; i >= 0; i-- {
		step := completedSteps[i]
		fmt.Println("rollback step: ", step.GetName())
		if err := step.Rollback(data); err != nil {
			fmt.Println("Error Rollback step:", err)
		}
	}
}

type SagaStep interface {
	GetName() string
	Execute(data interface{}) error
	Rollback(data interface{}) error
}

type PaymentStep struct {
}

func (c *PaymentStep) GetName() string {
	return "PaymentStep"
}

func (c *PaymentStep) Execute(data interface{}) error {
	postURL := "http://localhost:8081/api/v1/payment"
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return err
	}

	resp, err := http.Post(postURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return err
	}
	fmt.Println("Response status:", resp.Status)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("error response from payment api: %s", resp.Status)
	}
	defer resp.Body.Close()
	return nil
}

func (c *PaymentStep) Rollback(data interface{}) error {
	postURL := "http://localhost:8081/api/v1/refund"
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return err
	}

	resp, err := http.Post(postURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return err
	}
	fmt.Println("Response status:", resp.Status)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("error response from payment refund-api: %s", resp.Status)
	}
	defer resp.Body.Close()
	return nil
}

type DriverStep struct {
}

func (c *DriverStep) GetName() string {
	return "DriverStep"
}

func (c *DriverStep) Execute(data interface{}) error {
	postURL := "http://localhost:8080/api/v1/driver"
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return err
	}

	resp, err := http.Post(postURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return err
	}
	fmt.Println("Response status:", resp.Status)
	if resp.StatusCode >= 400 {
		return fmt.Errorf("error response from driver api: %s", resp.Status)
	}
	defer resp.Body.Close()
	return nil
}

func (c *DriverStep) Rollback(data interface{}) error {
	return nil
}
