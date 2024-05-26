package historian

type Historian struct {
	Databases map[string]map[string][]Field
}

func (historian *Historian) Start() {
	//
}

func (historian *Historian) Insert(request Request) {

}
