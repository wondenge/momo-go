/*
 * Remittance
 *
 * Partner Gateway API document
 *
 * API version: 1.0
 *
 */

package remittance

// PreApprovalResult struct for PreApprovalResult
type PreApprovalResult struct {
	Payer Party `json:"payer,omitempty"`
	// ISO4217 Currency
	PayerCurrency string `json:"payerCurrency,omitempty"`
	// The mesage that is shown to the approver.
	PayerMessage string `json:"payerMessage,omitempty"`
	// The request validity time of the pre-approval
	ValidityTime int32       `json:"validityTime,omitempty"`
	Status       string      `json:"status,omitempty"`
	Reason       ErrorReason `json:"reason,omitempty"`
}
