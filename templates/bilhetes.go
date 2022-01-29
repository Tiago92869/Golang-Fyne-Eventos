package principal

import (
	//"fmt"
	//"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/lxn/win"
)

/**STRUCT BILHETES**/
type ticket struct {

	/**CORPO**/
	//hbox
	hbox *gtk.Box

	/**LADO ESQUERDO**/
	//vbox lado esquerdo
	vbox_esq *gtk.Box
	//label one esq
	label_esq_one *gtk.Label
	//calendar
	calendar *gtk.Calendar
	//label two esq
	label_esq_two *gtk.Label
	//label three esq
	label_esq_three *gtk.Label
	//label four esq
	label_esq_four *gtk.Label
	//label five esq
	label_esq_five *gtk.Label
	//label six esq
	label_esq_six *gtk.Label
	//spinner button
	spinner_button *gtk.SpinButton
	//button esq
	button_esq *gtk.Button
	/**LADO DIREITO**/
	//vbox lado direito
	vbox_direito *gtk.Box
	//scrolable para o list
	scroll_for_list *gtk.ScrolledWindow
	//list direito
	list_dir *gtk.ListBox
	//hbox dir baixo
	hbox_dir_baixo *gtk.Box
	//button one dir
	button_dir_one *gtk.Button
	//button two dir
	button_dir_two *gtk.Button
}

/**TICKET INIT**/
var tick ticket

/**INIT DO TICK**/
func InitTick() {

	/**SCREEN SIZES**/
	//width
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	//height
	//iniciar css
	initcss()
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	/**BODY**/
	//hbox
	tick.hbox, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.090))
	//vbox esquerda
	tick.vbox_esq, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//vbox direita
	tick.vbox_direito, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//empty hbox
	empty_hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	tick.hbox.PackStart(empty_hbox, false, false, 0)
	//pack on hbox
	tick.hbox.PackStart(tick.vbox_esq, true, true, 0)
	tick.hbox.PackStart(tick.vbox_direito, true, true, 0)

	/**ESQUERDA**/
	//empty
	empty_box_one, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_one, false, false, uint(float64(height)*0.02))
	//label esq one
	tick.label_esq_one, _ = gtk.LabelNew("Selecione a data que deseja ir")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_one, true, true, 2)
	tick.label_esq_one.SetName("eventstext")
	//calendar
	tick.calendar, _ = gtk.CalendarNew()
	//packing
	tick.vbox_esq.PackStart(tick.calendar, true, true, 2)
	println(tick.calendar.GetName())
	//empty
	empty_box_four, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_four, true, true, uint(float64(height)*0.02))
	//label esq one
	tick.label_esq_two, _ = gtk.LabelNew("Nome do Evento")
	tick.label_esq_two.SetProperty("xalign", 0.5)
	//naming
	tick.label_esq_two.SetName("eventstext")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_two, true, true, 2)
	//label esq three
	tick.label_esq_three, _ = gtk.LabelNew("o-nome-fica-aqui")
	tick.label_esq_three.SetProperty("xalign", 0.5)
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_three, true, true, 2)

	//empty
	empty_box_two, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_two, true, true, uint(float64(height)*0.02))

	//label esq four
	tick.label_esq_four, _ = gtk.LabelNew("PREÇO")
	//naming
	tick.label_esq_four.SetName("eventstext")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_four, true, true, 2)

	//label esq five
	tick.label_esq_five, _ = gtk.LabelNew("o-preço-fica-aqui")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_five, true, true, 2)

	//empty
	empty_box_three, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_three, true, true, uint(float64(height)*0.02))

	//label esq six
	tick.label_esq_six, _ = gtk.LabelNew("Quantidade")
	//naming
	tick.label_esq_six.SetName("eventstext")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_six, true, true, uint(float64(height)*0.02))

	//spinner
	tick.spinner_button, _ = gtk.SpinButtonNewWithRange(0, 100, 1)
	tick.spinner_button.SetProperty("xalign", 0.57)
	//packing
	tick.vbox_esq.PackStart(tick.spinner_button, true, true, uint(float64(height)*0.007))
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_four, true, true, 2)
	//empty
	empty_box_five, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_five, true, true, uint(float64(height)*0.04))
	//button
	tick.button_esq, _ = gtk.ButtonNewWithLabel("CRIAR")
	tick.button_esq.Connect("clicked", func() {
		ano, mes, dia := tick.calendar.GetDate()
		println(ano)
		println(mes)
		println(dia)
	})
	tick.button_esq.SetName("top-level")
	//packing
	tick.vbox_esq.PackStart(tick.button_esq, true, true, 0)

	/**DIREITO**/
	//scroll window
	tick.scroll_for_list, _ = gtk.ScrolledWindowNew(nil, nil)
	//hbox
	hbox_no_user, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//another one
	hbox_no_usertwo, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	hbox_no_user.PackStart(hbox_no_usertwo, false, false, uint(float64(height)*0.001))
	//packing
	hbox_no_user.PackStart(tick.scroll_for_list, true, true, 0)
	//hbox
	hbox_no_userthree, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	hbox_no_user.PackStart(hbox_no_userthree, false, false, uint(float64(height)*0.04))
	//pack
	tick.vbox_direito.PackStart(hbox_no_user, true, true, uint(float64(height)*0.05))
	//lista
	tick.list_dir, _ = gtk.ListBoxNew()
	tick.list_dir.SetName("tuvias")
	//create an row
	label, _ := gtk.LabelNew("alo")
	tick.list_dir.Add(label)
	//pack the scroll
	tick.scroll_for_list.Add(tick.list_dir)
	//hbox
	tick.hbox_dir_baixo, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.20))
	//packing
	tick.vbox_direito.PackStart(tick.hbox_dir_baixo, false, false, uint(float64(height)*0.006))
	//button dir one
	tick.button_dir_one, _ = gtk.ButtonNewWithLabel("INFO")
	//naming
	tick.button_dir_one.SetName("buttonevents")
	//ACAO DO BOTAO
	tick.button_dir_one.Connect("clicked", func() {

		label, _ = gtk.LabelNew("lol")
		tick.list_dir.Add(label)
		tick.list_dir.ShowAll()
	})

	//pack
	tick.hbox_dir_baixo.PackStart(tick.button_dir_one, true, true, 0)
	//button dir two
	tick.button_dir_two, _ = gtk.ButtonNewWithLabel("DELETE")
	//naming
	tick.button_dir_two.SetName("buttonevents")
	//ACAO DO BOTAO
	tick.button_dir_two.Connect("clicked", func() {
		tick.list_dir.Remove(tick.list_dir.GetSelectedRow())
	})

	//pack
	tick.hbox_dir_baixo.PackStart(tick.button_dir_two, true, true, uint(float64(width)*0.043))
	//listener
	/**go func() {
		for range time.Tick(time.Second) {
			ano, mes, dia := tick.calendar.GetDate()
			fmt.Println("Escutei o ano->", ano, "o mes->", mes, "e o dia->", dia)
		}
	}()**/

}
func GetHeadTick() *gtk.Box {
	return tick.hbox
}
