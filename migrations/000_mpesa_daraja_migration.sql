-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS customer_to_business (
                                      id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                      msisdn VARCHAR(100) NOT NULL,
                                      transaction_type VARCHAR(100) NOT NULL,
                                      transaction_id VARCHAR(100) NOT NULL,
                                      transaction_amount DECIMAL(16,2) DEFAULT NULL CHECK ( transaction_amount >= 0 ),
                                      mpesa_transaction_id VARCHAR(100) NOT NULL,
                                      mpesa_transaction_time VARCHAR(100) NOT NULL,
                                      business_short_code VARCHAR(100) NOT NULL,
                                      bill_reference_number VARCHAR(100) NOT NULL,
                                      invoice_number VARCHAR(100),
                                      org_account_balance DECIMAL(16,2) DEFAULT NULL CHECK ( org_account_balance >= 0 ),
                                      third_party_trans_id VARCHAR(100),
                                      first_name VARCHAR(100),
                                      middle_name VARCHAR(100),
                                      last_name VARCHAR(100),
                                      result_code VARCHAR(100),
                                      result_description VARCHAR(100),
                                      created_at TIMESTAMP DEFAULT NOW(),
                                      updated_at TIMESTAMP
);

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS initiator (
                           id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                           initiator_name VARCHAR(100) NOT NULL,
                           initiator_credential TEXT NOT NULL,
                           created_at TIMESTAMP DEFAULT NOW(),
                           updated_at TIMESTAMP
);

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS business_to_customer (
                                      id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                      initiator_id BIGINT UNSIGNED DEFAULT NULL CHECK ( initiator_id > 0 ),
                                      originator_conversation_id VARCHAR(100) NOT NULL,
                                      command_id VARCHAR(100) NOT NULL,
                                      amount DECIMAL(16,2) DEFAULT NULL CHECK ( amount >= 0 ),
                                      party_a VARCHAR(100) NOT NULL,
                                      party_b VARCHAR(100) NOT NULL,
                                      remarks TEXT,
                                      occasion VARCHAR(100),
                                      conversation_id VARCHAR(100) NOT NULL,
                                      response_code VARCHAR(100),
                                      response_description VARCHAR(100),
                                      error_request_id VARCHAR(100),
                                      error_code VARCHAR(100),
                                      error_message VARCHAR(100),
                                      result_type VARCHAR(100),
                                      result_code VARCHAR(100),
                                      result_description VARCHAR(100),
                                      mpesa_transaction_id VARCHAR(100),
                                      transaction_amount DECIMAL(16,2) DEFAULT NULL CHECK ( transaction_amount >= 0 ),
                                      mpesa_transaction_receipt VARCHAR(100),
                                      is_recipient_registered_customer VARCHAR(10),
                                      charges_paid_account_available_funds DECIMAL(16,2) DEFAULT NULL CHECK ( charges_paid_account_available_funds >= 0 ),
                                      receiver_party_public_name VARCHAR(100),
                                      mpesa_transaction_completed_date DATETIME,
                                      utility_account_available_funds DECIMAL(16,2) DEFAULT NULL CHECK ( utility_account_available_funds >= 0 ),
                                      working_account_available_funds DECIMAL(16,2) DEFAULT NULL CHECK ( working_account_available_funds >= 0 ),
                                      created_at TIMESTAMP DEFAULT NOW(),
                                      updated_at TIMESTAMP,
                                      FOREIGN KEY (initiator_id) REFERENCES initiator(id)
);

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS business_buy_goods (
                                    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                    initiator_id BIGINT UNSIGNED DEFAULT NULL CHECK ( initiator_id > 0 ),
                                    command_id VARCHAR(100) NOT NULL,
                                    sender_identifier_type VARCHAR(10),
                                    receiver_identifier_type VARCHAR(10),
                                    transaction_amount DECIMAL(16,2) DEFAULT NULL CHECK ( transaction_amount >= 0 ),
                                    party_a VARCHAR(100) NOT NULL,
                                    party_b VARCHAR(100) NOT NULL,
                                    remarks TEXT,
                                    account_reference VARCHAR(100),
                                    requester VARCHAR(100),
                                    originator_conversation_id VARCHAR(100) NOT NULL,
                                    conversation_id VARCHAR(100) NOT NULL,
                                    response_code VARCHAR(100),
                                    response_description VARCHAR(100),
                                    result_type VARCHAR(100),
                                    result_code VARCHAR(100),
                                    result_description VARCHAR(100),
                                    mpesa_transaction_id VARCHAR(100),
                                    debit_account_balance VARCHAR(300),
                                    amount DECIMAL(16,2) DEFAULT NULL CHECK ( amount >= 0 ),
                                    debit_party_affected_account_balance VARCHAR(300),
                                    mpesa_transaction_completed_time VARCHAR(100),
                                    debit_party_charges DECIMAL(16,2) DEFAULT NULL CHECK ( debit_party_charges >= 0 ),
                                    receiver_party_public_name VARCHAR(100),
                                    currency VARCHAR(100),
                                    initiator_account_current_balance VARCHAR(100),
                                    bill_reference_number VARCHAR(100),
                                    error_completed_time VARCHAR(100),
                                    created_at TIMESTAMP DEFAULT NOW(),
                                    updated_at TIMESTAMP,
                                    FOREIGN KEY (initiator_id) REFERENCES initiator(id)
);

-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS mpesa_express (
                               id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                               business_short_code VARCHAR(100) NOT NULL,
                               timestamp VARCHAR(100),
                               transaction_type VARCHAR(100) NOT NULL,
                               amount DECIMAL(16,2) DEFAULT NULL CHECK ( amount >= 0 ),
                               party_a VARCHAR(100) NOT NULL,
                               party_b VARCHAR(100) NOT NULL,
                               phone_number VARCHAR(100) NOT NULL,
                               account_reference VARCHAR(200),
                               transaction_description VARCHAR(200),
                               merchant_request_id VARCHAR(100),
                               checkout_request_id VARCHAR(100),
                               response_code VARCHAR(10),
                               response_description VARCHAR(200),
                               customer_message VARCHAR(200),
                               result_code VARCHAR(100),
                               result_description VARCHAR(100),
                               mpesa_receipt_number VARCHAR(100),
                               transaction_date VARCHAR(100),
                               created_at TIMESTAMP DEFAULT NOW(),
                               updated_at TIMESTAMP
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS customer_to_business;
DROP TABLE IF EXISTS initiator;
DROP TABLE IF EXISTS business_to_customer;
DROP TABLE IF EXISTS business_buy_goods;
DROP TABLE IF EXISTS mpesa_express;