package boilerplate

import (
	"strings"
	"testing"

	"github.com/zhuharev/guse/routes/parser"
)

func TestRoutes(t *testing.T) {
	rw := NewRoutesWriter()
	rw.prefix = "github.com/zhuharev"
	rw.projectName = "allo"

	roteus, _ := parser.Parse(strings.NewReader(testRoutes))

	rw.WriteRoutes("./lol", roteus)

	//os.RemoveAll("./lol")
}

var testRoutes = `# Администратор
Text      shops.add_director_ 		      	shops.AddDirector

# Директор
Cb          shops.create                      shops.Create
Cb          shops.list 						  shops.List

# Магазин
Prefix 			shops.set_geo_ 	            shops.SetGeo
Prefix 			shops.withdraw_ 				shops.Withdraw
Prefix          shops.stats                     shops.Stats
Prefix          shops.add_manager               shops.AddManager
Prefix          shops.toggle_state              shops.ToggleState

Prefix 			products.upload_photo_         products.Upload
Prefix 	        products.set_price_            products.SetPrice
Prefix          products.set_title_            products.SetTitle
Cb              products.search                   products.Search
Prefix          products.buy_                  products.Buy
Prefix          products.buy_qiwi_             products.BuyQiwi`
