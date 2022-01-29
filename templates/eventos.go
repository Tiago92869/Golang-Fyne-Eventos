package principal

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/lxn/win"
)

type eventos struct {

	/**CORPO**/
	//hbox
	hbox *gtk.Box

	/**LADO ESQUERDO**/
	//vbox lado esquerdo
	vbox_esquerdo *gtk.Box
	//label esquerdo one
	label_esq_one *gtk.Label
	//entry esquerdo one
	entry_esq_one *gtk.Entry
	//label esq two
	label_esq_two *gtk.Label
	//entry esq two
	entry_esq_two *gtk.Entry
	//label esqc three
	label_esq_three *gtk.Label
	//calendar
	calendar_esq *gtk.Calendar
	//label esqc four
	label_esq_four *gtk.Label
	//entry esq three
	entry_esq_three *gtk.Entry
	//hbox esq one
	hbox_esq_one *gtk.Box
	//label esq five
	label_esq_five *gtk.Label
	//label esqc six
	label_esq_six *gtk.Label
	//hbox esq two
	hbox_esq_two *gtk.Box
	//spinner
	spinner_one *gtk.SpinButton
	//spinner two
	spinner_two *gtk.SpinButton
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

/**VAR INIT**/
var event eventos

/**GET HEAD**/
func GetHeadEventos() *gtk.Box {
	return event.hbox
}

func InitEvents() {
	/**SCREEN SIZES**/
	//width
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	//height
	//iniciar css
	initcss()
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	/**BODY**/
	//hbox
	event.hbox, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.090))
	//vbox esquerda
	event.vbox_esquerdo, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//vbox direita
	event.vbox_direito, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//empty hbox
	empty_hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	event.hbox.PackStart(empty_hbox, false, false, 0)
	//pack on hbox
	event.hbox.PackStart(event.vbox_esquerdo, true, true, 0)
	event.hbox.PackStart(event.vbox_direito, true, true, 0)

	/**ESQUERDA**/
	//empty
	empty_box_one, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_one, false, false, uint(float64(height)*0.02))
	//label esq one
	event.label_esq_one, _ = gtk.LabelNew("NOME DO EVENTO")
	event.label_esq_one.SetName("eventstext")
	//packing
	event.vbox_esquerdo.PackStart(event.label_esq_one, true, true, 2)
	//entry esq one
	event.entry_esq_one, _ = gtk.EntryNew()
	event.entry_esq_one.SetProperty("xalign", 0.5)
	//packing
	event.vbox_esquerdo.PackStart(event.entry_esq_one, true, true, 2)
	//empty
	empty_box_two, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_two, true, true, uint(float64(height)*0.02))
	//label esq two
	event.label_esq_two, _ = gtk.LabelNew("PREÇO")
	event.label_esq_two.SetName("eventstext")
	//packing
	event.vbox_esquerdo.PackStart(event.label_esq_two, true, true, 2)
	//entry esq two
	event.entry_esq_two, _ = gtk.EntryNew()
	event.entry_esq_two.SetProperty("xalign", 0.5)
	//packing
	event.vbox_esquerdo.PackStart(event.entry_esq_two, true, true, 2)
	//empty
	empty_box_three, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_three, true, true, uint(float64(height)*0.02))
	//label three
	event.label_esq_three, _ = gtk.LabelNew("SELECIONE A DATA")
	event.label_esq_three.SetName("eventstext")
	//packing
	event.vbox_esquerdo.PackStart(event.label_esq_three, true, true, 2)
	//calendar
	event.calendar_esq, _ = gtk.CalendarNew()
	//packing
	event.vbox_esquerdo.PackStart(event.calendar_esq, true, true, 2)
	println(event.calendar_esq.GetName())
	//empty
	empty_box_four, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_four, true, true, uint(float64(height)*0.02))
	//label four
	event.label_esq_four, _ = gtk.LabelNew("DURAÇÃO EM HORAS")
	event.label_esq_four.SetName("eventstext")
	//packing
	event.vbox_esquerdo.PackStart(event.label_esq_four, true, true, 2)
	//entry esq three
	event.entry_esq_three, _ = gtk.EntryNew()
	event.entry_esq_three.SetProperty("xalign", 0.5)
	//packing
	event.vbox_esquerdo.PackStart(event.entry_esq_three, true, true, 2)
	//empty
	empty_box_five, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_five, true, true, uint(float64(height)*0.04))
	//hbox one
	event.hbox_esq_one, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	//packing
	event.vbox_esquerdo.PackStart(event.hbox_esq_one, true, true, 0)
	//label five
	event.label_esq_five, _ = gtk.LabelNew("HORAS")
	event.label_esq_five.SetName("eventstext")
	//packing
	event.hbox_esq_one.PackStart(event.label_esq_five, true, true, 20)
	//label six
	event.label_esq_six, _ = gtk.LabelNew("MINUTES")
	event.label_esq_six.SetName("eventstext")
	//packing
	event.hbox_esq_one.PackStart(event.label_esq_six, true, true, 20)
	//hbox two
	event.hbox_esq_two, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 20)
	//packing
	event.vbox_esquerdo.PackStart(event.hbox_esq_two, true, true, 0)
	//spinner
	event.spinner_one, _ = gtk.SpinButtonNewWithRange(0, 23, 1)
	event.spinner_one.SetProperty("xalign", 0.57)
	//spin
	event.spinner_one.SetName("spinner")
	//packing
	event.hbox_esq_two.PackStart(event.spinner_one, true, true, 20)
	//spinner
	event.spinner_two, _ = gtk.SpinButtonNewWithRange(0, 59, 1)
	event.spinner_two.SetProperty("xalign", 0.57)
	//spin
	event.spinner_two.SetName("spinner")
	//packing
	event.hbox_esq_two.PackStart(event.spinner_two, true, true, 0)
	//empty
	empty_box_six, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_six, true, true, uint(float64(height)*0.03))
	//button
	event.button_esq, _ = gtk.ButtonNewWithLabel("CRIAR")
	event.button_esq.Connect("clicked", func() {
		ano, mes, dia := event.calendar_esq.GetDate()
		println(ano)
		println(mes)
		println(dia)
	})
	event.button_esq.SetName("top-level")
	//packing
	event.vbox_esquerdo.PackStart(event.button_esq, true, true, 0)

	/**DIREITO**/
	//scroll window
	event.scroll_for_list, _ = gtk.ScrolledWindowNew(nil, nil)
	//hbox
	hbox_no_user, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//another one
	hbox_no_usertwo, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	hbox_no_user.PackStart(hbox_no_usertwo, false, false, uint(float64(height)*0.001))
	//packing
	hbox_no_user.PackStart(event.scroll_for_list, true, true, 0)
	//hbox
	hbox_no_userthree, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	hbox_no_user.PackStart(hbox_no_userthree, false, false, uint(float64(height)*0.04))
	//pack
	event.vbox_direito.PackStart(hbox_no_user, true, true, uint(float64(height)*0.05))
	//lista
	event.list_dir, _ = gtk.ListBoxNew()
	event.list_dir.SetName("tuvias")
	//create an row
	label, _ := gtk.LabelNew("alo")
	event.list_dir.Add(label)
	//pack the scroll
	event.scroll_for_list.Add(event.list_dir)
	//hbox
	event.hbox_dir_baixo, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.20))
	//packing
	event.vbox_direito.PackStart(event.hbox_dir_baixo, false, false, uint(float64(height)*0.006))
	//button dir one
	event.button_dir_one, _ = gtk.ButtonNewWithLabel("INFO")
	//ACAO DO BOTAO
	event.button_dir_one.Connect("clicked", func() {

		label, _ = gtk.LabelNew("lol")
		event.list_dir.Add(label)
		event.list_dir.ShowAll()
	})
	event.button_dir_one.SetName("buttonevents")
	//pack
	event.hbox_dir_baixo.PackStart(event.button_dir_one, true, true, 0)
	//button dir two
	event.button_dir_two, _ = gtk.ButtonNewWithLabel("DELETE")
	//ACAO DO BOTAO
	event.button_dir_two.Connect("clicked", func() {
		event.list_dir.Remove(event.list_dir.GetSelectedRow())
	})
	event.button_dir_two.SetName("buttonevents")
	//pack
	event.hbox_dir_baixo.PackStart(event.button_dir_two, true, true, uint(float64(width)*0.043))

}
