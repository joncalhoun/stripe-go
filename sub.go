package stripe

import "encoding/json"

// SubscriptionStatus is the list of allowed values for the subscription's status.
// Allowed values are "trialing", "active", "past_due", "canceled", "unpaid".
type SubscriptionStatus string

const (
	Trialing SubscriptionStatus = "trialing"
	Active   SubscriptionStatus = "active"
	PastDue  SubscriptionStatus = "past_due"
	Canceled SubscriptionStatus = "canceled"
	Unpaid   SubscriptionStatus = "unpaid"
)

// SubParams is the set of parameters that can be used when creating or updating a subscription.
// For more details see https://stripe.com/docs/api#create_subscription and https://stripe.com/docs/api#update_subscription.
type SubParams struct {
	Params
	Customer, Plan       string
	Coupon, Token        string
	TrialEnd             int64
	Card                 *CardParams
	Quantity             uint64
	FeePercent           float64
	NoProrate, EndCancel bool
}

// SubListParams is the set of parameters that can be used when listing active subscriptions.
// For more details see https://stripe.com/docs/api#list_subscriptions.
type SubListParams struct {
	ListParams
	Customer string
}

// Subscription is the resource representing a Stripe subscription.
// For more details see https://stripe.com/docs/api#subscriptions.
type Subscription struct {
	Id          string             `json:"id"`
	EndCancel   bool               `json:"cancel_at_period_end"`
	Customer    *Customer          `json:"customer"`
	Plan        *Plan              `json:"plan"`
	Quantity    uint64             `json:"quantity"`
	Status      SubscriptionStatus `json:"status"`
	FeePercent  float64            `json:"application_fee_percent"`
	Canceled    int64              `json:"canceled_at"`
	PeriodEnd   int64              `json:"current_period_end"`
	PeriodStart int64              `json:"current_period_start"`
	Discount    *Discount          `json:"discount"`
	Ended       int64              `json:"ended_at"`
	Meta        map[string]string  `json:"metadata"`
	TrialEnd    int64              `json:"trial_end"`
	TrialStart  int64              `json:"trial_start"`
}

// SubscriptionList is a list object for subscriptions.
type SubscriptionList struct {
	ListResponse
	Values []*Subscription `json:"data"`
}

func (s *Subscription) UnmarshalJSON(data []byte) error {
	type sub Subscription
	var ss sub
	err := json.Unmarshal(data, &ss)
	if err == nil {
		*s = Subscription(ss)
	} else {
		// the id is surrounded by escaped \, so ignore those
		s.Id = string(data[1 : len(data)-1])
	}

	return nil
}
