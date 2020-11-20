package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(202, gin.H{
			"message": "世界你好呀",
		})
	})
	r.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"methods": "GET BOOK",
		})
	})
	r.POST("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"methods": "GET POST",
		})
	})
	r.PUT("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"methods": "GET PUT",
		})
	})
	r.DELETE("/book", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"methods": "GET DELETE",
		})
	})
	// json格式数据渲染
	r.GET("/jsonTest", func(ctx *gin.Context) {
		var person struct {
			Name   string `json:"name"` // 首字母小写无法渲染
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}
		person.Name = "李四"
		person.Age = 16
		person.Gender = "男"
		ctx.JSON(http.StatusOK, person)
	})

	r.GET("/json2", func(ctx *gin.Context) {
		/*
			// H is a shortcut for map[string]interface{}
			type H map[string]interface{}
		*/
		ctx.JSON(http.StatusOK, gin.H{"hello": "json格式数据"})
	})

	r.GET("/json", func(ctx *gin.Context) {
		var person = make(map[string]string)
		person["name"] = "李白"
		person["age"] = "26岁"
		person["gender"] = "女"
		ctx.JSON(http.StatusOK, person)
	})

	// 获取GET参数
	r.GET("/user", func(ctx *gin.Context) {

		name := ctx.Query("name")
		age := ctx.Query("age")
		// 设置默认值
		gender := ctx.DefaultQuery("gender", "男") // 取不到就是要默认的值
		gender, ok := ctx.GetQuery("gender")
		if !ok {
			gender = gender
		}
		fmt.Println(name)
		fmt.Println(ctx)

		ctx.JSON(http.StatusOK, gin.H{"message": "query参数获取", name: name, age: age, gender: gender})
	})
	r.POST("/login", func(ctx *gin.Context) {
		userName := ctx.PostForm("userName")
		userAge := ctx.PostForm("userAge")
		fmt.Println("--------", userName, userAge)
		ctx.JSON(http.StatusOK, gin.H{"name": userName, "age": userAge})
	})
	r.GET("/user/search/:userName/:userAge", func(ctx *gin.Context) {
		userName := ctx.Param("userName")
		userAge := ctx.Param("userAge")
		fmt.Println("--------", userName, userAge)
		ctx.JSON(http.StatusOK, gin.H{"name": userName, "age": userAge})
	})

	type Login struct {
		UseName  string `form:"userName" json:"userName" binding:"required"`
		PassWord int    `form:"passWord" json:"passWord" binding:"required"`
	}

	// 参数绑定
	r.POST("/loginJson", func(ctx *gin.Context) {
		var login Login
		if error := ctx.ShouldBind(&login); error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": error.Error()})
		} else {
			fmt.Printf("login info:%#v\n", login)
			// main.Login{UseName:"李四", PassWord:123456789}
			ctx.JSON(http.StatusOK, gin.H{
				"userName": login.UseName,
				"passWord": login.PassWord,
			})
		}
	})

	r.POST("/loginForm", func(ctx *gin.Context) {
		var login Login
		if error := ctx.ShouldBind(&login); error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": error.Error()})
		} else {
			fmt.Printf("login info:%#v\n", login)
			// main.Login{UseName:"李四", PassWord:123456789}
			ctx.JSON(http.StatusOK, gin.H{
				"userName": login.UseName,
				"passWord": login.PassWord,
			})
		}
	})

	// 单文件上传
	r.POST("/upload", func(ctx *gin.Context) {
		// 单个文件
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		/*
		Header:map[
			Content-Disposition:[form-data; name="file"; filename="850184275cef3311aaca0b8658a88db0.png"]
			Content-Type:[image/png]]
			Size:374486
			content:[],
		    tmpfile:
		]
		 */
		fmt.Printf("%+v\n",file)

		log.Println(file.Filename)
		dst := fmt.Sprintf("D:\\code\\goproject\\src\\gin_demo\\tmp/%s", file.Filename)
		// 上传文件到指定的目录
		ctx.SaveUploadedFile(file, dst)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	r.POST("/uploads", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("D:\\code\\goproject\\src\\gin_demo\\tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run(":3000")
}
