package APIs

import (
	. "github.com/julycw/ZhenhaiHRaSSB/controllers"
	"github.com/julycw/ZhenhaiHRaSSB/models"
	"github.com/julycw/orm"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ArticleAPIController struct {
	BaseController
}

func (this *ArticleAPIController) Get() {
}

func (this *ArticleAPIController) GetArticleList() {
	conditions := []orm.Condition{}
	orders := []orm.Order{}
	pageIndex, _ := this.GetInt("page")
	if pageIndex <= 0 {
		pageIndex = 1
	}
	pageSize, _ := this.GetInt("size")
	if pageSize <= 0 {
		pageSize = PageSize
	}

	cid := this.GetString("cid")
	if cid != "" {
		conditions = append(conditions, orm.Condition{Name: "cid", Compare: "=", Value: "'" + cid + "'"})
	}

	orderby := this.GetString("orderby")
	if orderby != "" {
		if strings.HasPrefix(orderby, "-") {
			orders = append(orders, orm.Order{Name: orderby[1:], By: orm.DESC})
		} else {
			orders = append(orders, orm.Order{Name: orderby, By: orm.ASC})
		}
	}

	articleList := models.ArticleStore.GetByPageAndConditionAndOrder(int(pageIndex), int(pageSize), conditions, orders)
	dataBag := DataBag{
		Message:  "",
		Datetime: time.Now(),
	}
	if len(articleList) > 0 {
		dataBag.Status = DataBagStatusSuccess
		dataBag.Data = articleList
	} else {
		dataBag.Status = DataBagStatusFailed
	}
	this.Data["json"] = &dataBag
	this.ServeJson()
}

func (this *ArticleAPIController) GetArticleOne() {
	gid := this.GetString(":gid")
	articleList := models.ArticleStore.GetByPageAndConditionAndOrder(1, 1, []orm.Condition{
		orm.Condition{Name: "gid", Compare: "=", Value: gid},
	}, []orm.Order{})

	dataBag := DataBag{
		Data:     nil,
		Message:  "",
		Datetime: time.Now(),
	}
	if len(articleList) > 0 {
		dataBag.Status = DataBagStatusSuccess
		dataBag.Data = articleList[0]
	} else {
		dataBag.Status = DataBagStatusFailed
	}
	this.Data["json"] = &dataBag
	this.ServeJson()
}

func (this *ArticleAPIController) GetArticleIcon() {
	gid := this.GetString(":gid")
	size := this.GetString(":size")
	sizeParams := strings.Split(strings.ToLower(size), "x")
	var width, height int
	if len(sizeParams) == 2 {
		width, _ = strconv.Atoi(sizeParams[0])
		height, _ = strconv.Atoi(sizeParams[1])
	}

	if width <= 0 {
		width = 120
	}

	if height <= 0 {
		height = 80
	}

	//local C:\\Rso\\cms\\ido\\resource\\2\\
	thumbIconURL := ".\\static\\img\\thumb\\" + gid + "-" + size + ".jpg"
	if _, err := os.Stat(thumbIconURL); err == nil {
		log.Printf("return img from file system:%v", thumbIconURL)
		this.ResponseFile("image/jpeg", thumbIconURL)
		return
	} else {
		log.Printf("can't find img in file system!")
	}
	iid := "'IID" + gid + "'"
	log.Println(iid)

	conditions := []orm.Condition{}
	conditions = append(conditions, orm.Condition{Name: "iid", Compare: "=", Value: iid})
	resourceList := models.ArticleResourceStore.GetByPageAndConditionAndOrder(1, 1, conditions, []orm.Order{})

	if len(resourceList) > 0 {
		resource := resourceList[0].(models.ArticleResource)
		if strings.HasSuffix(resource.Resource, ".jpg") {
			if err := createThumbPic(
				"C:\\Rso\\cms\\ido\\resource\\2\\"+resource.Resource,
				thumbIconURL,
				uint(width),
				uint(height)); err != nil {
				log.Printf(err.Error())
			} else {
				log.Printf("create thumb and return img from file system:%v", thumbIconURL)
				this.ResponseFile("image/jpeg", thumbIconURL)
				return
			}
		}
	} else {
		log.Printf("no compeated iid!\n")
	}

	// thumbIconURL = ".\\static\\img\\thumb\\default-" + size + ".jpg"
	log.Printf("return 404\n")
	this.Abort("404")

	// this.ResponseFile("image/jpeg", thumbIconURL)
}

func createThumbPic(source, to string, width, height uint) error {
	source_file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer source_file.Close()

	thumb_pic, err := os.Create(to)
	if err != nil {
		return err
	}
	defer thumb_pic.Close()

	// decode jpeg into image.Image
	img, err := jpeg.Decode(source_file)
	if err != nil {
		return err
	}

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Thumbnail(width, height, img, resize.Lanczos3)

	// write new image to file
	if err := jpeg.Encode(thumb_pic, m, nil); err != nil {
		return err
	}
	return nil
}
