package http

import (
	"io"
	"io/ioutil"
	_http "net/http"
	"time"

	"github.com/crazy-genius/imgresizer/internal/imgresizer/resizer"
	"github.com/gin-gonic/gin"
)

type resizeModel struct {
	Path string `form:"path"`

	Width  uint `form:"width"`
	Height uint `form:"height"`

	err string
}

func (rm *resizeModel) isValid() bool {

	if len(rm.Path) == 0 {
		rm.err = "No path provided"
		return false
	}

	if rm.Width == 0 {
		rm.err = "Width shoupd be grether than 0"
		return false
	}

	if rm.Height == 0 {
		rm.err = "Height shoupd be grether than 0"
		return false
	}

	rm.err = "No errors."

	return true
}

func (rm resizeModel) getError() string {
	return rm.err
}

func hello(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.String(_http.StatusOK, "Welcome Gin Server")
}

func resize(c *gin.Context) {
	var queryModel resizeModel

	if err := c.ShouldBindQuery(&queryModel); err != nil {
		reportError(c, err)
		return
	}

	if queryModel.isValid() {
		rs := resizer.NewResizer()

		resp, err := _http.Get(queryModel.Path)
		if err != nil {
			reportError(c, err)
			return
		}

		c.Request.Header.Set("Content-Type", "image/jpg")

		c.Stream(func(w io.Writer) bool {
			err := rs.Resize(ioutil.NopCloser(resp.Body), c.Writer, resizer.ResizeConfig{
				Dimenstions: resizer.Dimenstions{
					Height: queryModel.Height,
					Width:  queryModel.Width,
				},
				Quality: 90,
			})

			if err != nil {
				reportError(c, err)
				return true
			}

			if err := resp.Body.Close(); err != nil {
				reportError(c, err)
				return true
			}

			return false
		})
	} else {
		c.String(_http.StatusBadRequest, queryModel.getError())
	}
}

func reportError(ctx *gin.Context, err error) {
	ctx.String(_http.StatusInternalServerError, err.Error())
}
