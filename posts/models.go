package posts 

import (
	"github.com/astaxie/beego/orm"
)

//Define models in here

type Post struct {
	Id    int
	Title string
	Body  string
}

//register models here

func init() {
	orm.RegisterModel(new(Post))
}
