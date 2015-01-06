package models

type Catalog struct {
	GID     string `col:"gid" PK:"true"`
	Subject string `col:"subject"`
}
