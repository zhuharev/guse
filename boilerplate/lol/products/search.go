// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package products

import (
  "github.com/zhuharev/allo/web/context"
"github.com/zhuharev/allo/models"

)


// Search is products.search controller
func Search(ctx *context.Context) {

	ctx.HTML(200, "products/search")
}
}
