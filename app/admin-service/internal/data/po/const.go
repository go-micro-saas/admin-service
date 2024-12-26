package po

const (
	KeyPrefix = "bas_"
)

func Key(k string) string {
	return KeyPrefix + k
}
