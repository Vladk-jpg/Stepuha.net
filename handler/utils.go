package handler

import (
	"fmt"
	"github.com/oklog/ulid"
	"math/rand"
	"strings"
	"time"
)

func randomFileName() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return strings.ToLower(fmt.Sprintf("%v", ulid.MustNew(ulid.Timestamp(t), entropy)))
}
