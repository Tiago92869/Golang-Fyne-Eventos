package principal

import (
	//"fmt"
	//"time"

	"fmt"
	"package/back-end"
	"strconv"
	"time"

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
	//date
	date *gtk.Label
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
}

/**TICKET INIT**/
var tick ticket

//SWITCH//
var switch_update int

/**INIT DO TICK**/
func InitTick() {
	//COMECA A 0
	switch_update = 0
	/**SCREEN SIZES**/
	//width
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	//height
	//iniciar css
	initcss()
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	/**BODY**/
	//hbox
	tick.hbox, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(width)*0.040))
	//vbox esquerda
	tick.vbox_esq, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//vbox direita
	tick.vbox_direito, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//empty hbox
	empty_hbox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 40)
	tick.hbox.PackStart(empty_hbox, false, false, 0)
	//pack on hbox
	tick.hbox.PackStart(tick.vbox_esq, true, true, 0)
	tick.hbox.PackStart(tick.vbox_direito, true, true, 0)

	/**ESQUERDA**/
	//empty
	empty_box_one, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_one, false, false, uint(float64(height)*0.04))
	//label esq one
	tick.label_esq_one, _ = gtk.LabelNew("DATA")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_one, true, true, 0)
	tick.label_esq_one.SetName("eventstext")
	//DATA
	tick.date, _ = gtk.LabelNew("Selecione um evento")
	//packing
	tick.vbox_esq.PackStart(tick.date, true, true, 0)
	//empty
	empty_box_four, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_four, true, true, uint(float64(height)*0.04))
	//label esq one
	tick.label_esq_two, _ = gtk.LabelNew("Nome do Evento")
	tick.label_esq_two.SetProperty("xalign", 0.5)
	//naming
	tick.label_esq_two.SetName("eventstext")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_two, true, true, 10)
	//label esq three
	tick.label_esq_three, _ = gtk.LabelNew("o-nome-fica-aqui")
	tick.label_esq_three.SetProperty("xalign", 0.5)
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_three, true, true, 0)

	//empty
	empty_box_two, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	tick.vbox_esq.PackStart(empty_box_two, true, true, uint(float64(height)*0.03))

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
	tick.vbox_esq.PackStart(empty_box_three, true, true, uint(float64(height)*0.04))

	//label esq six
	tick.label_esq_six, _ = gtk.LabelNew("Quantidade")
	//naming
	tick.label_esq_six.SetName("eventstext")
	//packing
	tick.vbox_esq.PackStart(tick.label_esq_six, true, true, uint(float64(height)*0.02))

	//spinner
	tick.spinner_button, _ = gtk.SpinButtonNewWithRange(1, 100, 1)
	tick.spinner_button.SetProperty("xalign", 0.55)
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
		index := tick.list_dir.GetSelectedRow().GetIndex()
		name_e := posi[index]
		var ev back.Evento
		ev = back.GetEvento(name_e)
		numb_tickets_string, _ := tick.spinner_button.GetText()
		numb_tickets, _ := strconv.Atoi(numb_tickets_string)
		for i := 0; i < numb_tickets; i++ {
			back.BuyTicket(ev.Nome)
		}
		tick.spinner_button.SetText("0")
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
	hbox_no_user.PackStart(hbox_no_usertwo, false, false, uint(float64(height)*0.010))
	//packing
	hbox_no_user.PackStart(tick.scroll_for_list, true, true, 0)
	//hbox
	hbox_no_userthree, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	hbox_no_user.PackStart(hbox_no_userthree, false, false, uint(float64(height)*0.010))
	//pack
	tick.vbox_direito.PackStart(hbox_no_user, true, true, uint(float64(height)*0.06))
	//lista
	tick.list_dir, _ = gtk.ListBoxNew()
	tick.list_dir.SetName("tuvias")
	//ENCHER A NOSSA LISTA
	for i := 0; i < len(back.Lista_eventos); i++ {
		//CRIAR A STRING COM O QUE VAMOS INSERIR DENTRO
		//nome
		nome := back.Lista_eventos[i].Nome
		//datas
		//ano
		ano := back.Lista_eventos[i].DataInicio.AnoI
		//mes
		mes := back.Lista_eventos[i].DataInicio.MesI
		//dia
		dia := back.Lista_eventos[i].DataInicio.DiaI
		//horas
		horas := back.Lista_eventos[i].DataInicio.HoraI
		//minutes
		minutes := back.Lista_eventos[i].DataInicio.MinutoI
		//preco
		preco := back.Lista_eventos[i].Preço
		//participantes
		participantes := back.Lista_eventos[i].Participantes
		//duracao
		dia_F := back.Lista_eventos[i].DataFim.DiaF
		//mes
		mes_F := back.Lista_eventos[i].DataFim.MesF
		//ano
		ano_F := back.Lista_eventos[i].DataFim.AnoF
		//horas
		horas_F := back.Lista_eventos[i].DataFim.HoraF
		//minutos fim
		minutes_F := back.Lista_eventos[i].DataFim.MinutoF
		/**COLOCAR DENTRO DA LISTA**/
		hbox_list, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		//images
		image_pessoa, _ := gtk.ImageNewFromPixbuf(buildiconscale("pessoa.png", int(float64(height)*0.04), int(float64(height)*0.04), true))
		image_data, _ := gtk.ImageNewFromPixbuf(buildiconscale("hora.png", int(float64(height)*0.04), int(float64(height)*0.04), true))
		image_dinheiro, _ := gtk.ImageNewFromPixbuf(buildiconscale("dinheiro.png", int(float64(height)*0.04), int(float64(height)*0.04), true))
		image_evento, _ := gtk.ImageNewFromPixbuf(buildiconscale("evento_list.png", int(float64(height)*0.04), int(float64(height)*0.04), true))
		//labels
		label_evento, _ := gtk.LabelNew(nome)
		pessoas_string := fmt.Sprintf("%d", participantes)
		label_pessoa, _ := gtk.LabelNew(pessoas_string)
		//criar string para colocar dentro
		datas := fmt.Sprintf("%02d/%02d/%02d %02d:%02d - %02d/%02d/%02d %02d:%02d", dia, mes, ano, horas, minutes, dia_F, mes_F, ano_F, horas_F, minutes_F)
		label_data, _ := gtk.LabelNew(datas)
		preco_string := fmt.Sprintf("%.2f", preco)
		label_money, _ := gtk.LabelNew(preco_string)
		//packing
		hbox_list.PackStart(image_evento, true, true, 10)
		hbox_list.PackStart(label_evento, true, true, 0)
		hbox_list.PackStart(image_data, true, true, 10)
		hbox_list.PackStart(label_data, true, true, 0)
		hbox_list.PackStart(image_pessoa, true, true, 10)
		hbox_list.PackStart(label_pessoa, true, true, 0)
		hbox_list.PackStart(image_dinheiro, true, true, 10)
		hbox_list.PackStart(label_money, true, true, 0)
		tick.list_dir.Add(hbox_list)
		tick.list_dir.ShowAll()
	}
	//pack the scroll
	tick.scroll_for_list.Add(tick.list_dir)
	//hbox
	tick.hbox_dir_baixo, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.3))
	//packing
	tick.vbox_direito.PackStart(tick.hbox_dir_baixo, false, false, uint(float64(width)*0.001))
	//button dir one
	tick.button_dir_one, _ = gtk.ButtonNewWithLabel("INFO")
	//naming
	tick.button_dir_one.SetName("buttontick")
	//ACAO DO BOTAO
	tick.button_dir_one.Connect("clicked", func() {
		index := tick.list_dir.GetSelectedRow().GetIndex()
		tick.list_dir.ShowAll()
		//bloquear janela
		event.hbox.SetSensitive(false)
		//CRIAR DIALOG MESSAGE
		dialog, _ := gtk.DialogNew()
		//content
		content, _ := dialog.GetContentArea()
		//CRIAR LABEL
		name_e := posi[index]
		var a back.Evento
		a = back.GetEvento(name_e)
		number_e := len(a.Bilhete)
		number := strconv.Itoa(number_e)
		label, _ := gtk.LabelNew("O numero de bilhetes vendidos é: " + number)
		//CENTRAR TEXTO
		label.SetProperty("xalign", 0.57)
		content.SetCenterWidget(label)
		//SETAR O TITULO DA JANELA
		dialog.SetTitle("WARNING")
		//SETAR POSI DA JANELA
		dialog.SetPosition(gtk.WIN_POS_CENTER)
		//MOSTRAR JANELA
		dialog.ShowAll()
		//SETAR UM ICON PARA A JANELA
		dialog.SetIconFromFile("evento.png")
		//IR BUSCAR UMA ACAO QUASO A JANELA SEJA DESTRUIDA
		dialog.Connect("destroy", func() {
			event.hbox.SetSensitive(true)
		})
	})

	//pack
	tick.hbox_dir_baixo.SetCenterWidget(tick.button_dir_one)

	//listener
	/**go func() {
		for range time.Tick(time.Second) {
			ano, mes, dia := tick.calendar.GetDate()
			fmt.Println("Escutei o ano->", ano, "o mes->", mes, "e o dia->", dia)
		}
	}()**/

}

/**EXPORTAR ITEMS**/
//head
func GetHeadTick() *gtk.Box {
	return tick.hbox
}

//FUNCAO PARA ATIVAR O TIMING
func Updatedata() {

	go func() {
		//MUDAR O SWITCH PARA LIGADO
		switch_update = 1
		//FUNCAO COM LOOP
		for range time.Tick(time.Second) {
			//IR BUSCAR A INDEX
			index := tick.list_dir.GetSelectedRow().GetIndex()
			//index
			if index != -1 {
				//IR BUSCAR QUANTIDADE
				quantidade_string, _ := tick.spinner_button.GetText()
				quantidade, _ := strconv.Atoi(quantidade_string)
				//IR BUSCAR EVENTO
				nome_evento := back.Lista_eventos[index].Nome
				//SETAR TEXTO
				tick.label_esq_three.SetText(nome_evento)
				var preco float64
				if quantidade != 0 {
					//IR BUSCAR PRECO
					preco = back.Lista_eventos[index].Preço * float64(quantidade)
				} else {
					//IR BUSCAR O PRECO
					preco = back.Lista_eventos[index].Preço
				}
				//STRING
				preco_string := fmt.Sprintf("%.01f", preco)
				tick.label_esq_five.SetText(preco_string)
				//ir buscar date
				//dia
				dia := back.Lista_eventos[index].DataInicio.DiaI
				//mes
				mes := back.Lista_eventos[index].DataInicio.MesI
				//ano
				ano := back.Lista_eventos[index].DataInicio.AnoI
				//horas
				horas := back.Lista_eventos[index].DataInicio.HoraI
				//minutes
				minutes := back.Lista_eventos[index].DataInicio.MinutoI
				//ir buscar date fim
				//dia f
				dia_f := back.Lista_eventos[index].DataFim.DiaF
				//mes f
				mes_f := back.Lista_eventos[index].DataFim.MesF
				//ano f
				ano_f := back.Lista_eventos[index].DataFim.AnoF
				//horas f
				horas_f := back.Lista_eventos[index].DataFim.HoraF
				//minutes f
				minutes_f := back.Lista_eventos[index].DataFim.MinutoF
				//date string
				date_string := fmt.Sprintf("%02d/%02d/%02d %02d:%02d - %02d/%02d/%02d %02d:%02d", dia, mes, ano, horas, minutes, dia_f, mes_f, ano_f, horas_f, minutes_f)
				//DATE LABEL
				tick.date.SetText(date_string)

			} else {
				//CASO NAO SEJA SELECIONADO NADA
				tick.label_esq_three.SetText("Selecione um evento")
				tick.label_esq_five.SetText("Selecione um evento")
			}
			//CONDICAO PARA SAIR
			if switch_update == 0 {
				break
			}
		}
	}()

}

//FAZER COM QUE O LOOP PARE
func Stop() {
	switch_update = 0
}
