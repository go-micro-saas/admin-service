package global

const (
	KeyPrefix = "account_service_"
)

func Key(k string) string {
	return KeyPrefix + k
}
