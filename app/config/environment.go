package config

// Environment variables
var Environment = map[string]interface{}{
	"app_id":                 "6A6EF734-F557-4F7E-9C2A-32CD28E43420",
	"app_version":            "v1.0.0",
	"app_name":               "unknown",
	"app_description":        "",
	"port":                   8000,
	"timezone":               "UTC",
	"enable_migration":       true, // always set true this value for production usage
	"endpoint":               "/api/v1",
	"environment":            "development",
	"db_host":                "postgres",
	"db_port":                5432,
	"db_user":                "postgres",
	"db_pass":                "postgres",
	"db_name":                "postgres",
	"db_table_prefix":        "",
	"redis_host":             "redis",
	"redis_port":             6379,
	"redis_pass":             "",
	"redis_index":            0,
	"prefork":                false,
	"json_engine":            "sonic",                            // available options: sonic, go-json, encoding/json
	"aes":                    "mB53IvZupVcalBnlEPmLyl4xJD4YN6g4", // AES 256-bit must have key at least 32 byte
	"host_multimedia_server": "http://ms-multimedia:8003/api/v1",
	"host_oauth2_server":     "http://oauth2:9096/api/v1",
	"oauth2_client_id":       "my-client-id",
	"oauth2_client_secret":   "password",
	"oauth2_domain":          "http://localhost",
	"header_token_key":       "x-Token",
	"value_token_key":        "v0x37KYbJqLodL0363Xa6jxaRTTN2eD1",
	"access_token":           "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjFjNjdkYWFkLTVhODYtNDBlNi1hNzlhLTQ5ZGVhYWE2ZDg0YSIsImJ1c2luZXNzX2lkIjoiNmExYzEyZmQtNzEzNC00YTVhLTkwYmQtZGM1YzEyNmZkNThhIiwiYXVkIjoibXktY2xpZW50LWlkIiwiZXhwIjoxNjg2Njc3NzkyLCJzdWIiOiIxYzY3ZGFhZC01YTg2LTQwZTYtYTc5YS00OWRlYWFhNmQ4NGEifQ.-qbG3YQn6WsOvn5dEzIVCgUfA_wXmXjQcgWeBoK62KUyFtvyiOw5dAN9zWmZcnhBr0jMpME4nVVKlrP3gDvi0A",
	"user_id":                "1c67daad-5a86-40e6-a79a-49deaaa6d84a",
	"business_id":            "6a1c12fd-7134-4a5a-90bd-dc5c126fd58a",
	"enable_subscription":    true,
	"default_length_count":   []int{4, 4, 4, 3, 3, 5, 5, 4, 4}, // employee, customer, product, supplier, business_location, purchase_order, transaction_selling, stock_opname, payment_transaction
}
