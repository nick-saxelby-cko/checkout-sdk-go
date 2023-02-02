package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments/nas"
)

func TestGetPaymentActions(t *testing.T) {
	paymentResponse := makeCardPayment(t, false, 10)

	cases := []struct {
		name         string
		instrumentId string
		checker      func(interface{}, error)
	}{
		{
			name:         "when payment id is valid then return payment actions",
			instrumentId: paymentResponse.Id,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.(*nas.GetPaymentActionsResponse).Actions)

				for _, action := range response.(*nas.GetPaymentActionsResponse).Actions {
					assert.NotEmpty(t, action.Amount)
					assert.True(t, action.Approved)
					assert.Nil(t, action.Links)
					assert.NotEmpty(t, action.ProcessedOn)
					assert.NotEmpty(t, action.Reference)
					assert.NotEmpty(t, action.ResponseCode)
					assert.NotEmpty(t, action.ResponseSummary)
					assert.NotEmpty(t, action.Type)
				}
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			process := func() (interface{}, error) { return client.GetPaymentActions(tc.instrumentId) }
			predicate := func(data interface{}) bool {
				response := data.(*nas.GetPaymentActionsResponse)
				return response.Actions != nil && len(response.Actions) >= 0
			}

			tc.checker(retriable(process, predicate, 2))
		})
	}
}
