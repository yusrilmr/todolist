package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/yusrilmr/todolist/backend/common"
	"github.com/yusrilmr/todolist/backend/controllers"
)

func SetLabelRoutes(router *mux.Router) *mux.Router {
	labelRouter := mux.NewRouter()
	labelRouter.HandleFunc("/labels", controllers.CreateLabel).Methods("POST")
	labelRouter.HandleFunc("/labels/{id}", controllers.UpdateLabel).Methods("PUT")
	labelRouter.HandleFunc("/labels/{id}", controllers.GetLabelById).Methods("GET")
	labelRouter.HandleFunc("/labels", controllers.GetLabels).Methods("GET")
	//labelRouter.HandleFunc("/labels/tasks/{id}", controllers.GetLabelsByTask).Methods("GET")
	labelRouter.HandleFunc("/labels/{id}", controllers.DeleteLabel).Methods("DELETE")
	router.PathPrefix("/labels").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni. Wrap(labelRouter),
	))
	return router
}