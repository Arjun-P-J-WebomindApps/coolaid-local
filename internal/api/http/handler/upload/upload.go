package upload

// import "github.com/gin-gonic/gin"

// func UploadHandler(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	query := c.PostForm("data")
// 	category := c.PostForm("category")

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "file required"})
// 		return
// 	}

// 	f, err := file.Open()
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "cannot open file"})
// 		return
// 	}
// 	defer f.Close()

// 	if err := uploadService.ImportCSV(
// 		ctx,
// 		query,
// 		category,
// 		f,
// 	); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, gin.H{"status": "uploaded"})
// }
