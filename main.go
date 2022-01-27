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

	/**SCREEN SIZES**/
	//width
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	//height
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	//init
	gtk.Init(nil)

	/**WINDOW CREATION**/
	//creation
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	//set title
	win.SetTitle("EVENTGOR")
	//set fullscreen
	win.Maximize()
	//setar açao no caso de ser fechada
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	/**BODY**/
	//vbox
	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//hbox btn
	hbox_btn, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//empacotar hbox na vbox
	vbox.PackStart(hbox_btn, true, true, 0)

	/**BTNS**/
	//btn data
	btn_data, _ := gtk.ButtonNewWithLabel("DATA")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_data, true, true, 1)
	//btn bilhetes
	btn_bilhetes, _ := gtk.ButtonNewWithLabel("BILHETES")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_bilhetes, true, true, 1)
	//btn inicio
	btn_inicio, _ := gtk.ButtonNewWithLabel("INICIO")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_inicio, true, true, 1)
	//btn eventos
	btn_eventos, _ := gtk.ButtonNewWithLabel("EVENTOS")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_eventos, true, true, 1)
	//btn sair
	btn_sair, _ := gtk.ButtonNewWithLabel("SAIR")
	//EMPACOTAR NA HBOX
	hbox_btn.PackStart(btn_sair, true, true, 1)

	/**SEGUNDA PARTE DO CORPO**/
	hbox_baixo, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	hbox_baixo.SetSizeRequest(0*width, int(float64(height)*1.05))
	//EMPACOTAR
	vbox.PackStart(hbox_baixo, true, false, 1)

	//APENAS PARA APARECER
	println(width)
	println(height)
	principal.Init()

	/**ADICIONAR Á WINDOW O CONTENT**/
	win.Add(vbox)
	win.ShowAll()
	gtk.Main()
}
