// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package products

import (
  "github.com/zhuharev/allo/web/context"
"github.com/zhuharev/allo/models"

)


// Upload is products.upload controller
func Upload(ctx *context.Context) {

	ctx.HTML(200, "products/upload")
}
}
