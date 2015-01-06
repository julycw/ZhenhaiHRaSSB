package models

import (
	"time"
)

type Article struct {
	GID     string    `col:"gid" PK:"true"`
	CID     string    `col:"cid"`
	Title   string    `col:"title"`
	Content string    `col:"content"`
	Author  string    `col:"author"`
	Status  int       `col:"status"`
	InTime  time.Time `col:"intime"`
}
