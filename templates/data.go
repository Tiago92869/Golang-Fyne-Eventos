package principal

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/lxn/win"
)

/**STRUCT DATA**/
type data struct {
	/**APENAS UM CORPO**/

	//vbox
	vbox *gtk.Box
	//hbox one
	hbox_one *gtk.Box
	//hbox two
	hbox_two *gtk.Box
	//hbox three
	hbox_three *gtk.Box
	//hbox four
	hbox_four *gtk.Box

	//label one
	label_one *gtk.Label
	//label two
	label_two *gtk.Label
	//label three
	label_three *gtk.Label
	//label four
	label_four *gtk.Label
	//label five
	label_five *gtk.Label
	//label six
	label_six *gtk.Label
	//label seven
	label_seven *gtk.Label
	//label eight
	label_eight *gtk.Label
}

/**INIT DATA**/
var date data

/**INIT FUNC**/
func InitData() {
	/**SCREEN SIZES**/
	//width
	//width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	//iniciar css
	initcss()
	//heigth
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	/**BODY**/
	//vbox
	date.vbox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//hbox one
	date.hbox_one, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.090))
	//hbox two
	date.hbox_two, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.100))
	//hbox three
	date.hbox_three, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.110))
	//hbox four
	date.hbox_four, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.120))
	//pack on vbox
	date.vbox.PackStart(date.hbox_one, true, true, 0)
	date.vbox.PackStart(date.hbox_two, true, true, 0)
	date.vbox.PackStart(date.hbox_three, true, true, 0)
	date.vbox.PackStart(date.hbox_four, true, true, 0)

	/**LAYOUT**/
	//label esquerda
	date.label_one, _ = gtk.LabelNew("Dinheiro Ganho no Mês Atual:")
	date.label_one.SetName("dinheiromensal")
	//packing
	date.hbox_one.PackStart(date.label_one, true, true, 2)
	//label resposta
	date.label_two, _ = gtk.LabelNew("GANHO")
	date.label_two.SetName("quantidademensal")
	//packing
	date.hbox_two.PackStart(date.label_two, true, true, 2)

	//label esquerda
	date.label_three, _ = gtk.LabelNew("Evento com mais pessoas:")
	date.label_one.SetName("eventomaior")
	//packing
	date.hbox_three.PackStart(date.label_three, true, true, 2)
	//label esquerda
	date.label_four, _ = gtk.LabelNew("TOTAL PESSOAS")
	date.label_four.SetName("dinheiromensal")
	//packing
	date.hbox_two.PackStart(date.label_four, true, true, 2)

	//label esquerda
	date.label_five, _ = gtk.LabelNew("Total de Dinheiro Ganho:")
	date.label_five.SetName("dinheirototal")
	//packing
	date.hbox_three.PackStart(date.label_five, true, true, 2)
	//label esquerda
	date.label_six, _ = gtk.LabelNew("DINHEIRO TOTAL")
	date.label_six.SetName("valortotal")
	//packing
	date.hbox_four.PackStart(date.label_six, true, true, 2)

	//label esquerda
	date.label_seven, _ = gtk.LabelNew("Nº Total de Bilhetes Vendidos:")
	date.label_seven.SetName("numerobilhetes")
	//packing
	date.hbox_three.PackStart(date.label_seven, true, true, 2)
	//label esquerda
	date.label_eight, _ = gtk.LabelNew("BILHETES VENDIDOS")
	date.label_eight.SetName("valorbilhetes")
	//packing
	date.hbox_four.PackStart(date.label_eight, true, true, 2)

}

//reutrn da cabeca
func GetData() *gtk.Box {
	return date.hbox_one
}
