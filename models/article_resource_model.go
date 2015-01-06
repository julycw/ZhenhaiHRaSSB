package models

type ArticleResource struct {
	GID      string `col:"gid" PK:"true"`
	IID      string `col:"iid"`
	Title    string `col:"title"`
	Resource string `col:"resource"`
	RType    string `col:"rtype"`
}
