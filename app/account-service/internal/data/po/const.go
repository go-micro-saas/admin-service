package po

const (
	KeyPrefix = "as_"
)

func Key(k string) string {
	return KeyPrefix + k
}
