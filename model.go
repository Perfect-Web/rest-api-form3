package main

// swagger:model
type RestApi struct {
	Data []struct {
		Type           string `json:"type"`
		ID             string `json:"id"`
		Version        int    `json:"version"`
		OrganisationID string `json:"organisation_id"`
		Attributes     struct {
			Amount           string `json:"amount"`
			BeneficiaryParty struct {
				AccountName       string `json:"account_name"`
				AccountNumber     string `json:"account_number"`
				AccountNumberCode string `json:"account_number_code"`
				AccountType       int    `json:"account_type"`
				Address           string `json:"address"`
				BankID            string `json:"bank_id"`
				BankIDCode        string `json:"bank_id_code"`
				Name              string `json:"name"`
			} `json:"beneficiary_party"`
			ChargesInformation struct {
				BearerCode    string `json:"bearer_code"`
				SenderCharges []struct {
					Amount   string `json:"amount"`
					Currency string `json:"currency"`
				} `json:"sender_charges"`
				ReceiverChargesAmount   string `json:"receiver_charges_amount"`
				ReceiverChargesCurrency string `json:"receiver_charges_currency"`
			} `json:"charges_information"`
			Currency    string `json:"currency"`
			DebtorParty struct {
				AccountName       string `json:"account_name"`
				AccountNumber     string `json:"account_number"`
				AccountNumberCode string `json:"account_number_code"`
				Address           string `json:"address"`
				BankID            string `json:"bank_id"`
				BankIDCode        string `json:"bank_id_code"`
				Name              string `json:"name"`
			} `json:"debtor_party"`
			EndToEndReference string `json:"end_to_end_reference"`
			Fx                struct {
				ContractReference string `json:"contract_reference"`
				ExchangeRate      string `json:"exchange_rate"`
				OriginalAmount    string `json:"original_amount"`
				OriginalCurrency  string `json:"original_currency"`
			} `json:"fx"`
			NumericReference     string `json:"numeric_reference"`
			PaymentID            string `json:"payment_id"`
			PaymentPurpose       string `json:"payment_purpose"`
			PaymentScheme        string `json:"payment_scheme"`
			PaymentType          string `json:"payment_type"`
			ProcessingDate       string `json:"processing_date"`
			Reference            string `json:"reference"`
			SchemePaymentSubType string `json:"scheme_payment_sub_type"`
			SchemePaymentType    string `json:"scheme_payment_type"`
			SponsorParty         struct {
				AccountNumber string `json:"account_number"`
				BankID        string `json:"bank_id"`
				BankIDCode    string `json:"bank_id_code"`
			} `json:"sponsor_party"`
		} `json:"attributes"`
	} `json:"data"`
}