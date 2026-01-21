package erogs

import "fmt"

func MakeDMMImageURL(dmm string) string {
	return fmt.Sprintf("https://pics.dmm.co.jp/digital/pcgame/%[1]s/%[1]spl.jpg", dmm)
}
