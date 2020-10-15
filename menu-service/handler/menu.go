package handler

import "net/http"

func AddMenu(w http.ResponseWriter, *http.Request){
	utils.WrapAPISuccess(w,r,"success", http.StatusOK)
}