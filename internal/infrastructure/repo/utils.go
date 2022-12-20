package repo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"
)

const _defaultEntityCap = 50

func decodeCursor(cursor string) (t time.Time, id, suffix string, err error) {
	decoded, err := base64.URLEncoding.DecodeString(cursor)
	if err != nil {
		return
	}

	splited := strings.Split(string(decoded), ",")
	if len(splited) != 3 {
		err = errors.New("cursor is invalid")
		return
	}

	t, err = time.Parse(time.RFC3339, splited[0])
	if err != nil {
		return
	}

	id = splited[1]
	suffix = splited[2]
	return
}

func encodeCursor(t time.Time, id, suffix string) string {
	cursor := fmt.Sprintf("%s,%s,%s", t.Format(time.RFC3339), id, suffix)
	return base64.URLEncoding.EncodeToString([]byte(cursor))
}
