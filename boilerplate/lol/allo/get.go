// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package allo

import (
  "github.com/zhuharev/allo/pkg/context"
)


func Get(ctx *context.Context) {

    ctx.HTML(200, "allo/get")
}
