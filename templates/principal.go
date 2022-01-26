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

	//vbox de cima one
	princip.vbox_cima_one, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	//setar tamanho hbox
	princip.hbox_cima.SetSizeRequest(0*width, int(float64(height)*0.20))
	//adicionar na hbox
	princip.hbox_cima.PackStart(princip.vbox_cima_one, true, false, 0)

	//vbox cima two
	princip.vbox_cima_two, _ = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)
	//adicionar na hbox
	princip.hbox_cima.PackStart(princip.vbox_cima_two, true, true, 0)

	//label cima one
	princip.label_cima_one, _ = gtk.LabelNew("AQUI DÁ PARA TUDO")
	//adicionar label na vbox cima one
	princip.vbox_cima_one.PackStart(princip.label_cima_one, false, false, 1)

	//label cima two
	princip.label_cima_two, _ = gtk.LabelNew("DESDE RESERVAR A CRIAR EVENTOS")
	//adicionar label na vbox cima two
	princip.vbox_cima_two.PackStart(princip.label_cima_two, false, false, 1)

}
