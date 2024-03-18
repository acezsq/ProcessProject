package model

import (
	"github.com/acezsq/project-common/errs"
)

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
)
