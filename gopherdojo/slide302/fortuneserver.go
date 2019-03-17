package slide302

import (
	"fmt"
	"math/rand"
	"net/http"
)

const SEED = 100
const MAX = 100

type Fortune struct {
	partOne int32
	partTwo int32
}

func (f Fortune) String() string {
	return fmt.Sprintf("%d-%d", f.partOne, f.partTwo)
}

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	f := Fortune{
		rand.Int31n(MAX),
		rand.Int31n(MAX),
	}
	fmt.Fprint(w, f.String())
}

func StartServer() {
	rand.Seed(SEED)
	http.HandleFunc("/fortune", fortuneHandler)
	http.ListenAndServe(":8080", nil)
}
