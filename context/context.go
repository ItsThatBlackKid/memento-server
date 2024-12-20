package context

import (
	"gorm.io/gorm"
)

type RequestContext struct {
	Userid int16
	DB     *gorm.DB
}

var Context RequestContext
