package principal

/**IMPORT DO GTK**/
/**IMPORT DO GDK**/
/**IMPORT DO GET WINDOW WIDTH & HEIGHT**/
import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"github.com/lxn/win"
)

/**STRUCT DE WIDGETS**/
type principal struct {
	/**CORPO DA JANELA**/
	vbox *gtk.Box
	/**PARTE DE CIMA DA JANELA**/
	//hbox de cima
	hbox_cima *gtk.Box
	//vbox de cima 1
	vbox_cima_one *gtk.Box
	//vbox de cima 2
	vbox_cima_two *gtk.Box
	//label de cima 1
	label_cima_one *gtk.Label
	//label de cima 2
	label_cima_two *gtk.Label
	//imagem de cima 1
	imagem_cima_one *gtk.Image
	//imagem de cima 2
	imagem_cima_two *gtk.Image

	/**PARTE DE BAIXO DA JANELA**/
	//hbox de baixo
	hbox_baixo *gtk.Box
	//vbox de baixo
	vbox_baixo *gtk.Box
	//label de baixo
	label_baixo *gtk.Label
	//imagem de baixo
	imagem_baixo *gtk.Image
}

/**FUNCAO PARA BUILD ICON**/
func buildicon(filename string) *gdk.Pixbuf {

	pix, err := gdk.PixbufNewFromFile(filename)

	if err != nil {
		log.Fatal("Erro ao dar load do icon")
	}

	return pix
}

/**FUNCAO PARA BUILD ICON COM SCALE**/
func buildiconscale(filename string, width int, height int, preserve bool) *gdk.Pixbuf {

	pix, err := gdk.PixbufNewFromFileAtScale(filename, width, height, true)

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

/**INICIO DO STRUCT**/
var princip principal

/**INIT DA JANELA**/
func Init() {

	/**INIT DO CSS LOADER**/
	//INICIALIZAR O CSS COM A FUNCAO JA DEFINIDA DE CSS
	initcss()
	/**SCREEN SIZES**/
	//width
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	//height
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	/**INIT DO CORPO**/
	princip.vbox, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)

	/**INIT DA PARTE DE CIMA**/
	//hbox de cima
	princip.hbox_cima, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
	//adicionar na vbox
	princip.vbox.PackStart(princip.hbox_cima, true, true, 0)
	//setar tamanho hbox
	princip.hbox_cima.SetSizeRequest(0*width, int(float64(height)*0.35))

	//vbox de cima one
	princip.vbox_cima_one, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	//adicionar na hbox
	princip.hbox_cima.PackStart(princip.vbox_cima_one, true, true, 0)
	//adicionar nome
	princip.vbox_cima_one.SetName("vboxescima")

	//vbox cima two
	princip.vbox_cima_two, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	//adicionar na hbox
	princip.hbox_cima.PackStart(princip.vbox_cima_two, true, true, 0)
	//adicionar nome
	princip.vbox_cima_two.SetName("vboxescima")

	//label cima one
	princip.label_cima_one, _ = gtk.LabelNew("              AQUI DÁ PARA TUDO            ")
	//setname
	princip.label_cima_one.SetName("janelacimatext")
	//adicionar espaco a cima
	empty_init, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	empty_init.SetSizeRequest(0, int(float64(height)*0.070))
	princip.vbox_cima_one.PackStart(empty_init, false, false, 0)
	//adicionar label na vbox cima one
	princip.vbox_cima_one.PackStart(princip.label_cima_one, false, false, 0)

	//label cima two
	princip.label_cima_two, _ = gtk.LabelNew("DESDE RESERVAR A CRIAR EVENTOS")
	//setname
	princip.label_cima_two.SetName("janelacimatext")
	//adicionar espaco a cima usando o empy de cima
	empty_inittwo, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	empty_inittwo.SetSizeRequest(0, int(float64(height)*0.070))
	princip.vbox_cima_two.PackStart(empty_inittwo, false, false, 0)
	//adicionar label na vbox cima two
	princip.vbox_cima_two.PackStart(princip.label_cima_two, false, false, 1)

	//foto cima one
	uau := buildiconscale("uau.png", int(float64(width)*0.15), int(float64(height)*0.15), true)
	princip.imagem_cima_one, _ = gtk.ImageNewFromPixbuf(uau)
	//inserir no container
	princip.vbox_cima_one.SetCenterWidget(princip.imagem_cima_one)

	//foto cima two
	ticket := buildiconscale("ticket.png", int(float64(width)*0.15), int(float64(height)*0.15), true)
	princip.imagem_cima_two, _ = gtk.ImageNewFromPixbuf(ticket)
	//inserir no container
	princip.vbox_cima_two.SetCenterWidget(princip.imagem_cima_two)

	/**PARTE DE BAIXO**/

	//vbox de baixo
	princip.hbox_baixo, _ = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//vbox de baixo
	princip.vbox_baixo, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	//adicionar na hbox
	princip.hbox_baixo.SetCenterWidget(princip.vbox_baixo)
	//criar vazio
	empty_up, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//alterar tamanho
	empty_up.SetSizeRequest(0, int(float64(height)*0.03))
	princip.vbox_baixo.PackStart(empty_up, true, true, 0)
	//criar label
	princip.label_baixo, _ = gtk.LabelNew("BEM-VINDO AO Ajax Eventos")
	//setname
	princip.label_baixo.SetName("janelacimatext")
	//adicionar na vbox
	princip.vbox_baixo.PackStart(princip.label_baixo, true, true, 0)
	//criar vazio
	empty_down, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//alterar tamanho
	empty_down.SetSizeRequest(0, int(float64(height)*0.03))
	princip.vbox_baixo.PackStart(empty_down, true, true, 0)
	//criar imagem
	estadio := buildiconscale("estadio.png", int(float64(width)*0.30), int(float64(height)*0.30), true)
	princip.imagem_baixo, _ = gtk.ImageNewFromPixbuf(estadio)
	//adicionar no centro da vbox
	princip.imagem_baixo.SetSizeRequest(20, 20)
	princip.vbox_baixo.PackStart(princip.imagem_baixo, false, false, 0)

	//adicionar parte de baixo ao corpo
	princip.vbox.PackStart(princip.hbox_baixo, true, true, 0)

}

//return da cabeca
func Gethead() *gtk.Box {
	return princip.vbox
}
