// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package shops

import (
  "github.com/zhuharev/allo/web/context"
"github.com/zhuharev/allo/models"

)


// Stats is shops.stats controller
func Stats(ctx *context.Context) {

	ctx.HTML(200, "shops/stats")
}
}
