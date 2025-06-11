package global

const (
	KeyPrefix = "acc_s_"
)

func Key(k string) string {
	return KeyPrefix + k
}
