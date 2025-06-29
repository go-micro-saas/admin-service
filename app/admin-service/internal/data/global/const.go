package global

const (
	KeyPrefix = "admin_service_"
)

func Key(k string) string {
	return KeyPrefix + k
}
