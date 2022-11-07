package ginupload

import (
	"awesomeProject1/common"
	"awesomeProject1/component/appctx"
	uploadbiz "awesomeProject1/module/upload/uploadbbusiness"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//folder := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename))
		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
