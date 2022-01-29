package principal

import "github.com/gotk3/gotk3/gtk"

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
	println("hello")
}
