package date

import "time"

func CurrentTimestamp() int64{
	cur := time.Now()
	timestamp := cur.UnixNano() / 1000000
	return timestamp
}

