package main

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//_"lovehome/models"
	"net/http"
	"strings"
	"github.com/astaxie/beego/context"
)
//
//func insertuser(){
//
//	o := orm.NewOrm()
//
//	user := models.User{}
//
//	user.Name = "mhq5"
//
//	if id, err := o.Insert( &user ); err == nil {
//
//		beego.Info("insert id is",id)
//
//	}else{
//
//		beego.Info("insert fail")
//
//	}
//
//}
//
//func insertusers(){
//
//	o := orm.NewOrm()
//
//	users := []models.User{
//		{Name:"mhq8"},
//		{Name:"mhq9"},
//	}
//
//	successnum,err := o.InsertMulti(1,users)
//
//	if err != nil{
//		beego.Info("有错误")
//	}
//
//	beego.Info("插入的数量",successnum)
//}
//
//func queryuser() {
//
//	o := orm.NewOrm()
//
//	user := models.User{
//		Name: "mhq6",
//	}
//
//	err := o.Read(&user,"Name")
//
//	if err != nil {
//
//		beego.Info("read fail")
//
//	}
//
//	beego.Info("this user  is", user)
//}
//
//func readorcreate(){
//
//	o := orm.NewOrm()
//
//	user := models.User{Name:"mhq7"}
//
//	created , id , err := o.ReadOrCreate(&user , "Name")
//
//	if created {
//		beego.Info("新创建的数据id是",id)
//		return
//	}
//
//	if err != nil {
//		beego.Info("有错误",err)
//		return
//	}
//
//
//	beego.Info("不是新创建的")
//	beego.Info(created , err , id)
//
//}
//func updateduser(){
//
//	o := orm.NewOrm()
//
//	user := models.User{Id:7}
//
//	if o.Read(&user) == nil {
//
//		user.Name = "mhq6"
//
//		num , err := o.Update( &user )
//
//		if err == nil {
//
//			beego.Info("更新了",num,"条")
//
//		}
//
//	}
//}
//func deleteduser(){
//
//	o := orm.NewOrm()
//
//	user := models.User{Id:7}
//
//	if num , err := o.Delete(&user);err == nil{
//
//		beego.Info("删除的条数", num )
//
//	}else {
//
//		beego.Info("删除失败！")
//
//	}
//
//}
//
//func insertorder(){
//
//	o := orm.NewOrm()
//
//	order := models.User_order{}
//
//	order.OrderDate = "this is data1"
//
//	user := models.User{Id:1}
//
//	order.User = &user
//
//	if _,err := o.Insert(&order);err == nil{
//
//		beego.Info("插入了")
//	}
//}
//
//func queryorder(){
//
//	o := orm.NewOrm()
//
//	var orders []*models.User_order
//
//	qs := o.QueryTable("User_order")
//
//	_,err := qs.Filter("user__id",1).All(&orders)
//
//	if err == nil{
//
//		for _,order := range orders{
//
//			beego.Info(order)
//
//		}
//
//	}
//
//
//
//}

func main() {
	ignoreStaticPath()
	beego.Run()
}



//chuangdingxiang
func ignoreStaticPath() {

	//透明static
	//beego.SetStaticPath("group1/M00/","fdfs/storage_data/data/")

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
