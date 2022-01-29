package main

/** IMPORT PRINCIPAL **/
/**IMPORT DO GTK**/
/**IMPORT DO GDK**/
/**IMPORT DO GET WINDOW WIDTH & HEIGHT**/
import (
	"log"
	principal "package/templates"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/lxn/win"
)

/**FUNCAO PARA BUILD ICON**/
func buildicon(filename string) *gdk.Pixbuf {

	pix, err := gdk.PixbufNewFromFile(filename)

	if err != nil {
		log.Fatal("Erro ao dar load do icon")
	}

	return pix
}

/**FUNCAO PARA INICIAR O CSS**/
func initcss() {
	//provider
	Provider, _ := gtk.CssProviderNew()
	//LOAD
	if err := Provider.LoadFromPath("principal.css"); err != nil {
		log.Println(err)
	}
	//SCREEN
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatalf("Unable to get screen: %v\n", err)
	}
	//IR BUSCAR AS PROPRIEDADES
	gtk.AddProviderForScreen(screen, Provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

}

/**MAIN**/
func main() {
	/**START**/
	location := "Inicio"
	/**SCREEN SIZES**/
	//width
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	//height
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	//init
	gtk.Init(nil)
	//init da funcao princip
	principal.Init()
	principal.InitEvents()
	principal.InitTick()
	//init do css
	initcss()
	/**WINDOW CREATION**/
	//creation
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	//set title
	win.SetTitle("Ajax Eventos")
	//set fullscreen
	win.Maximize()
	//setar açao no caso de ser fechada
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	//win name para setar a cor da janela
	win.SetName("windowp")
	//setar icon
	icon := buildicon("evento.png")
	//setar icon
	win.SetIcon(icon)

	/**BODY**/
	//vbox
	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, int(float64(height)*0.02))
	//hbox btn
	hbox_btn, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//empacotar hbox na vbox
	vbox.PackStart(hbox_btn, true, true, 0)

	/**BTNS**/
	//fazer um empty space inicio
	empty_init, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	empty_init.SetSizeRequest(int(float64(width)*0.024), 0)
	//empacotar
	hbox_btn.PackStart(empty_init, false, false, 0)
	//btn data
	btn_data, _ := gtk.ButtonNewWithLabel("DATA")
	//setar a classe
	btn_data.SetName("top-level")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_data, true, true, 1)
	//btn bilhetes
	btn_bilhetes, _ := gtk.ButtonNewWithLabel("BILHETES")
	//connect
	btn_bilhetes.Connect("clicked", func() {
		switch location {
		//TORNAR INVISIVEL A CABECA CORRENTE E VISIVEL A INVISIVEL
		case "Inicio":
			principal.Gethead().Hide()
			principal.GetHeadTick().Show()
			location = "Bilhetes"
		case "Eventos":
			principal.GetHeadEventos().Hide()
			principal.GetHeadTick().Show()
			location = "Bilhetes"
		case "Data":
			println("NOTHING YET")
		default:
			break
		}
	})
	//setar a classe
	btn_bilhetes.SetName("top-level")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_bilhetes, true, true, 1)
	//btn inicio
	btn_inicio, _ := gtk.ButtonNewWithLabel("INICIO")
	//connect
	btn_inicio.Connect("clicked", func() {
		switch location {
		//TORNAR INVISIVEL A CABECA CORRENTE E VISIVEL A INVISIVEL
		case "Bilhetes":
			principal.GetHeadTick().Hide()
			principal.Gethead().Show()
			location = "Inicio"
		case "Eventos":
			principal.GetHeadEventos().Hide()
			principal.Gethead().Show()
			location = "Inicio"
		case "Data":
			println("NOTHING YET")
		default:
			break
		}
	})
	//setar a classe
	btn_inicio.SetName("top-level")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_inicio, true, true, 1)
	//btn eventos
	btn_eventos, _ := gtk.ButtonNewWithLabel("EVENTOS")
	//connect
	btn_eventos.Connect("clicked", func() {
		switch location {
		//TORNAR INVISIVEL A CABECA CORRENTE E VISIVEL A INVISIVEL
		case "Bilhetes":
			principal.GetHeadTick().Hide()
			principal.GetHeadEventos().Show()
			location = "Eventos"
		case "Inicio":
			principal.Gethead().Hide()
			principal.GetHeadEventos().Show()
			location = "Eventos"
		case "Data":
			println("NOTHING YET")
		default:
			break
		}
	})
	//setar a classe
	btn_eventos.SetName("top-level")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_eventos, true, true, 1)
	//btn sair
	btn_sair, _ := gtk.ButtonNewWithLabel("SAIR")
	//PARA SAIR DO PROGRAMA
	btn_sair.Connect("clicked", func() {
		gtk.MainQuit()
	})
	//setar a classe
	btn_sair.SetName("top-level")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_sair, true, true, 1)
	//EMPTY SPACE NO FINAL
	empty_end, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	empty_end.SetSizeRequest(int(float64(width)*0.024), 0)
	//empacotar
	hbox_btn.PackEnd(empty_end, false, false, 0)

	/**SEGUNDA PARTE DO CORPO**/
	hbox_baixo, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	hbox_baixo.SetSizeRequest(0*width, int(float64(height)*0.989))
	//EMPACOTAR
	vbox.PackStart(hbox_baixo, true, false, 1)

	//CRIAR VBOX ONDE VAI ESTAR O CONTEUDO
	vbox_baixo, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//settar tamanho
	vbox_baixo.SetSizeRequest(int(float64(width)*1.20), 0)
	//settar nome de teste
	vbox_baixo.SetName("vboxcentral")
	//COLOCAR NO MEIO
	hbox_baixo.SetCenterWidget(vbox_baixo)
	//packing
	//CABECA DO EVENTOS MAS ESCONDIDA
	vbox_baixo.Add(principal.GetHeadEventos())
	//CABECA DO TICKET MAS ESCONDIDA
	vbox_baixo.Add(principal.GetHeadTick())
	//CABECA DO INICIAR
	vbox_baixo.Add(principal.Gethead())

	//APENAS PARA APARECER
	println(width)
	println(height)

	/**ADICIONAR Á WINDOW O CONTENT**/
	win.Add(vbox)
	win.ShowAll()
	//ESCONDER EVENTOS
	principal.GetHeadEventos().Hide()
	//ESCONDER OS TICKETS
	principal.GetHeadTick().Hide()
	gtk.Main()
}
