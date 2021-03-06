package principal

import (
	"fmt"
	"package/back-end"
	"strconv"

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
	//label esq four
	label_esq_four *gtk.Label
	//entry esq three
	entry_esq_three *gtk.Entry
	//label esq seven
	label_esq_seven *gtk.Label
	//entry esq four
	entry_esq_four *gtk.Entry
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

//remover slice
func removeSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

/**GET HEAD**/
func GetHeadEventos() *gtk.Box {
	return event.hbox
}

//CRIAR ARRAY PARA SABER POSICAO DOS ITEMS (ISTO PODE SER ACESSADO POR TODOS OS DO PACKAGE COMUM A ESTE)
var posi []string

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
	event.label_esq_two, _ = gtk.LabelNew("PRE??O")
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
	event.vbox_esquerdo.PackStart(event.label_esq_three, true, true, 0)
	//calendar
	event.calendar_esq, _ = gtk.CalendarNew()
	//packing
	event.vbox_esquerdo.PackStart(event.calendar_esq, true, true, 0)
	//empty
	empty_box_four, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_four, true, true, uint(float64(height)*0.01))
	//label four
	event.label_esq_four, _ = gtk.LabelNew("DURA????O EM HORAS")
	event.label_esq_four.SetName("eventstext")
	//packing
	event.vbox_esquerdo.PackStart(event.label_esq_four, true, true, 2)
	//entry esq three
	event.entry_esq_three, _ = gtk.EntryNew()
	event.entry_esq_three.SetProperty("xalign", 0.5)
	//packing
	event.vbox_esquerdo.PackStart(event.entry_esq_three, true, true, 0)
	//label seven
	event.label_esq_seven, _ = gtk.LabelNew("NUMERO DE PARTICIPANTES")
	event.label_esq_seven.SetName("eventstext")
	//packing
	event.vbox_esquerdo.PackStart(event.label_esq_seven, true, true, 0)
	//entry
	event.entry_esq_four, _ = gtk.EntryNew()
	event.entry_esq_four.SetProperty("xalign", 0.5)
	//packing
	event.vbox_esquerdo.PackStart(event.entry_esq_four, true, true, 0)

	//empty
	empty_box_five, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_five, true, true, uint(float64(height)*0.01))
	//hbox one
	event.hbox_esq_one, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(event.hbox_esq_one, true, true, 0)
	//label five
	event.label_esq_five, _ = gtk.LabelNew("HORAS")
	event.label_esq_five.SetName("eventstext")
	//packing
	event.hbox_esq_one.PackStart(event.label_esq_five, true, true, 0)
	//label six
	event.label_esq_six, _ = gtk.LabelNew("MINUTES")
	event.label_esq_six.SetName("eventstext")
	//packing
	event.hbox_esq_one.PackStart(event.label_esq_six, true, true, 20)
	//hbox two
	event.hbox_esq_two, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
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
	event.hbox_esq_two.PackStart(event.spinner_two, true, true, 20)
	//empty
	empty_box_six, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//packing
	event.vbox_esquerdo.PackStart(empty_box_six, true, true, uint(float64(height)*0.03))
	//button
	event.button_esq, _ = gtk.ButtonNewWithLabel("CRIAR")
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
		preco := back.Lista_eventos[i].Pre??o
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
		event.list_dir.Add(hbox_list)
		//adicionar no vetor para saber a posicao correspondente
		posi = append(posi, nome)
		event.list_dir.ShowAll()
	}
	//pack the scroll
	event.scroll_for_list.Add(event.list_dir)
	//hbox
	event.hbox_dir_baixo, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, int(float64(height)*0.20))
	//packing
	event.vbox_direito.PackStart(event.hbox_dir_baixo, false, false, uint(float64(height)*0.006))
	//button dir one
	event.button_dir_one, _ = gtk.ButtonNewWithLabel("EDITAR")
	event.button_dir_one.SetName("buttonevents")
	//pack
	event.hbox_dir_baixo.PackStart(event.button_dir_one, true, true, 0)
	//button dir two
	event.button_dir_two, _ = gtk.ButtonNewWithLabel("DELETE")
	event.button_dir_two.SetName("buttonevents")
	//pack
	event.hbox_dir_baixo.PackStart(event.button_dir_two, true, true, uint(float64(width)*0.043))

	//TRIGGERS
	//trigger
	event.button_esq.Connect("clicked", func() {
		//BUSCAR ANO,MES,DIA
		ano, mes, dia := event.calendar_esq.GetDate()
		//buscar dados
		nome_evento, _ := event.entry_esq_one.GetText()
		//pre??o
		preco_string, _ := event.entry_esq_two.GetText()
		preco, _ := strconv.ParseFloat(preco_string, 1)
		//duracao
		duracao_string, _ := event.entry_esq_three.GetText()
		duracao, _ := strconv.Atoi(duracao_string)
		//numero de participantes
		participantes_string, _ := event.entry_esq_four.GetText()
		participantes, _ := strconv.Atoi(participantes_string)
		//horas
		horas_string, _ := event.spinner_one.GetText()
		horas, _ := strconv.Atoi(horas_string)
		//minutes
		minutes_string, _ := event.spinner_two.GetText()
		minutes, _ := strconv.Atoi(minutes_string)

		if nome_evento != "" && preco_string != "" && participantes_string != "" && horas_string != "0" && minutes_string != "0" {
			//CRIAR EVENTO
			sucesso := back.AdicionarEvento(nome_evento, int(ano), int(mes)+1, int(dia), horas, minutes, duracao, participantes, 0, preco)
			if sucesso != 1 {
				//ADICIONAR NA LISTA
				//CRIAR A STRING COM O QUE VAMOS INSERIR DENTRO
				last := len(back.Lista_eventos) - 1
				//nome
				nome := back.Lista_eventos[last].Nome
				//datas
				//ano
				ano := back.Lista_eventos[last].DataInicio.AnoI
				//mes
				mes := back.Lista_eventos[last].DataInicio.MesI
				//dia
				dia := back.Lista_eventos[last].DataInicio.DiaI
				//horas
				horas := back.Lista_eventos[last].DataInicio.HoraI
				//minutes
				minutes := back.Lista_eventos[last].DataInicio.MinutoI
				//preco
				preco := back.Lista_eventos[last].Pre??o
				//participantes
				participantes := back.Lista_eventos[last].Participantes
				//duracao
				dia_F := back.Lista_eventos[last].DataFim.DiaF
				//mes
				mes_F := back.Lista_eventos[last].DataFim.MesF
				//ano
				ano_F := back.Lista_eventos[last].DataFim.AnoF
				//horas
				horas_F := back.Lista_eventos[last].DataFim.HoraF
				//minutos fim
				minutes_F := back.Lista_eventos[last].DataFim.MinutoF
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
				event.list_dir.Add(hbox_list)
				event.list_dir.ShowAll()
				//adicionar no vetor para saber a posicao correspondente
				posi = append(posi, nome)
				//clear things
				event.entry_esq_one.SetText("")
				event.entry_esq_two.SetText("")
				event.entry_esq_three.SetText("")
				event.entry_esq_four.SetText("")
				event.spinner_one.SetText("0")
				event.spinner_two.SetText("0")
			} else {
				//bloquear janela
				event.hbox.SetSensitive(false)
				//CRIAR DIALOG MESSAGE
				dialog, _ := gtk.DialogNew()
				//content
				content, _ := dialog.GetContentArea()
				//CRIAR LABEL
				label, _ := gtk.LabelNew("Existem datas nesse dia!")
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
			}
		} else {
			//bloquear janela
			event.hbox.SetSensitive(false)
			//CRIAR DIALOG MESSAGE
			dialog, _ := gtk.DialogNew()
			//content
			content, _ := dialog.GetContentArea()
			//CRIAR LABEL
			label, _ := gtk.LabelNew("Preencha os campos com valores validos!")
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

		}

	})

	//ACAO DO BOTAO
	event.button_dir_two.Connect("clicked", func() {
		//get index
		index := event.list_dir.GetSelectedRow().GetIndex()
		//CASO ESTEJA ALGO SELECIONADO..
		if index != -1 {
			//delte do evento
			back.DeleteEvento(posi[index])
			//remover slice
			removeSlice(posi, index)
			//remover widget
			event.list_dir.Remove(event.list_dir.GetSelectedRow())
		}
	})

	//ACAO DO BOTAO
	event.button_dir_one.Connect("clicked", func() {
		//ir buscar a index
		index := event.list_dir.GetSelectedRow().GetIndex()
		if index != -1 {
			//bloquear janela
			event.hbox.SetSensitive(false)
			//CRIAR DIALOG MESSAGE
			dialog, _ := gtk.DialogNew()
			//content
			content, _ := dialog.GetContentArea()
			//SETAR O TITULO DA JANELA
			dialog.SetTitle("EDITAR")
			//setar tamanho
			dialog.SetSizeRequest(200, 200)
			//SETAR POSI DA JANELA
			dialog.SetPosition(gtk.WIN_POS_CENTER)
			//SETAR UM ICON PARA A JANELA
			dialog.SetIconFromFile("evento.png")
			//criar vbox
			vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 20)
			//criar label nome
			label_nome, _ := gtk.LabelNew("INTRODUZA O NOME")
			//setname
			label_nome.SetName("eventstext")
			//criar label numero de pessoas
			label_pessoas, _ := gtk.LabelNew("INTRODUZA NUMERO DE PESSOAS")
			//setname
			label_pessoas.SetName("eventstext")
			//criar label preco
			label_preco, _ := gtk.LabelNew("PRE??O")
			label_preco.SetName("eventstext")
			//CRIAR ENTRYS nome
			entry_nome, _ := gtk.EntryNew()
			//set property
			entry_nome.SetProperty("xalign", 0.5)
			//CRIAR ENTRYS pessoas
			entry_pessoas, _ := gtk.EntryNew()
			//set property
			entry_pessoas.SetProperty("xalign", 0.5)
			//CRIAR ENTRYS pre??o
			entry_preco, _ := gtk.EntryNew()
			//set property
			entry_preco.SetProperty("xalign", 0.5)
			//CRIAR BUTTON
			button_edit, _ := gtk.ButtonNewWithLabel("EDITAR")
			//set the name
			button_edit.SetName("buttonevents")
			//packing
			vbox.PackStart(label_nome, true, true, 0)
			vbox.PackStart(entry_nome, true, true, 0)
			vbox.PackStart(label_pessoas, true, true, 0)
			vbox.PackStart(entry_pessoas, true, true, 0)
			vbox.PackStart(label_preco, true, true, 0)
			vbox.PackStart(entry_preco, true, true, 0)
			vbox.PackStart(button_edit, true, true, 0)
			content.SetCenterWidget(vbox)
			dialog.ShowAll()

			//TRIGGERS
			//button editar
			button_edit.Connect("clicked", func() {
				//tornar sensivel
				event.hbox.SetSensitive(true)
				//get name
				nome, _ := entry_nome.GetText()
				//get participantes
				participantes_str, _ := entry_pessoas.GetText()
				participantes, _ := strconv.Atoi(participantes_str)
				//get preco
				preco_str, _ := entry_preco.GetText()
				preco, _ := strconv.ParseFloat(preco_str, 1)
				if nome != "" && participantes != 0 && preco != 0.0 {
					dialog.Close()
					//MUDAR O PRECO
					back.EditTicketPrice(posi[index], preco)
					//EDITAR O EVENTO
					back.EditEvento(posi[index], nome, participantes)
					//LIMPAR LISTA EVENTS
					for i := 0; i < len(posi); i++ {
						event.list_dir.Remove(event.list_dir.GetRowAtIndex(0))
					}
					//LIMPAR LISTA PARA O TICK
					for i := 0; i < len(posi); i++ {
						tick.list_dir.Remove(tick.list_dir.GetRowAtIndex(0))
					}
					//ESVAZIAR POSI
					posi = nil

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
						preco := back.Lista_eventos[i].Pre??o
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
						//ADICIONAR TUDO DNV NA LISTA PARA O EVENTS
						event.list_dir.Add(hbox_list)
						//adicionar no vetor para saber a posicao correspondente
						posi = append(posi, nome)
						event.list_dir.ShowAll()
					}

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
						preco := back.Lista_eventos[i].Pre??o
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
						//ADICIONAR TUDO DNV NA LISTA PARA O EVENTS
						tick.list_dir.Add(hbox_list)
						//adicionar no vetor para saber a posicao correspondente
						posi = append(posi, nome)
						tick.list_dir.ShowAll()
					}

				}
			})
			//CONECT
			dialog.Connect("destroy", func() {
				event.hbox.SetSensitive(true)
			})
		} else {

			//bloquear janela
			event.hbox.SetSensitive(false)
			//CRIAR DIALOG MESSAGE
			dialog, _ := gtk.DialogNew()
			//content
			content, _ := dialog.GetContentArea()
			//CRIAR LABEL
			label, _ := gtk.LabelNew("Selecione um item para editar!")
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

		}

	})

}
