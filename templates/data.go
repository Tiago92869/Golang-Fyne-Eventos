package principal

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/lxn/win"
)

/**STRUCT datastruct**/
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

/**INIT datastruct**/
var datastruct data

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
	datastruct.vbox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//hbox one
	datastruct.hbox_one, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.090))
	//hbox two
	datastruct.hbox_two, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.100))
	//hbox three
	datastruct.hbox_three, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.110))
	//hbox four
	datastruct.hbox_four, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.120))
	//pack on vbox
	datastruct.vbox.PackStart(datastruct.hbox_one, true, true, 0)
	datastruct.vbox.PackStart(datastruct.hbox_two, true, true, 0)
	datastruct.vbox.PackStart(datastruct.hbox_three, true, true, 0)
	datastruct.vbox.PackStart(datastruct.hbox_four, true, true, 0)

	/**LAYOUT**/
	//label esquerda
	datastruct.label_one, _ = gtk.LabelNew("Dinheiro Ganho no Mês Atual:")
	datastruct.label_one.SetName("dinheiromensal")
	//packing
	datastruct.hbox_one.PackStart(datastruct.label_one, true, true, 2)
	//label resposta
	datastruct.label_two, _ = gtk.LabelNew("GANHO")
	datastruct.label_two.SetName("quantidademensal")
	//packing
	datastruct.hbox_two.PackStart(datastruct.label_two, true, true, 2)

	//label esquerda
	datastruct.label_three, _ = gtk.LabelNew("Evento com mais pessoas:")
	datastruct.label_one.SetName("eventomaior")
	//packing
	datastruct.hbox_three.PackStart(datastruct.label_three, true, true, 2)
	//label esquerda
	datastruct.label_four, _ = gtk.LabelNew("TOTAL PESSOAS")
	datastruct.label_four.SetName("dinheiromensal")
	//packing
	datastruct.hbox_two.PackStart(datastruct.label_four, true, true, 2)

	//label esquerda
	datastruct.label_five, _ = gtk.LabelNew("Total de Dinheiro Ganho:")
	datastruct.label_five.SetName("dinheirototal")
	//packing
	datastruct.hbox_three.PackStart(datastruct.label_five, true, true, 2)
	//label esquerda
	datastruct.label_six, _ = gtk.LabelNew("DINHEIRO TOTAL")
	datastruct.label_six.SetName("valortotal")
	//packing
	datastruct.hbox_four.PackStart(datastruct.label_six, true, true, 2)

	//label esquerda
	datastruct.label_seven, _ = gtk.LabelNew("Nº Total de Bilhetes Vendidos:")
	datastruct.label_seven.SetName("numerobilhetes")
	//packing
	datastruct.hbox_three.PackStart(datastruct.label_seven, true, true, 2)
	//label esquerda
	datastruct.label_eight, _ = gtk.LabelNew("BILHETES VENDIDOS")
	datastruct.label_eight.SetName("valorbilhetes")
	//packing
	datastruct.hbox_four.PackStart(datastruct.label_eight, true, true, 2)

}

//reutrn da cabeca
func Getdatas() *gtk.Box {
	return datastruct.vbox
}
