package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	a := emoji.Sprint("Hello :world_map:!")
	return a

}
