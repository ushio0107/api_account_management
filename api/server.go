package api

func (a *Api) Run() {
	r := a.NewRouter()

	r.Run()
}
