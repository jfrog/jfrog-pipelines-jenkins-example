package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type MathHandler struct {
}

func GetMathHandler() *MathHandler {
	return &MathHandler{
	}
}

func getNums(r *http.Request) (int, int, error) {
	values := r.URL.Query();

	num1str, num1ok := values["num1"]
	if !num1ok || len(num1str) < 1 {
		return -1, -1, errors.New("Missing num1 parameter.");
	}

	num1,num1Err := strconv.Atoi(num1str[0])
	if num1Err != nil {
		return -1, -1, errors.New("Parameter num1 is invalid.");
	}

	num2str, num2ok := values["num2"]
	if !num2ok || len(num2str) < 1 {
		return -1, -1, errors.New("Missing num2 parameter.");
	}

	num2,num2Err := strconv.Atoi(num2str[0])
	if  num2Err != nil {
		return -1, -1, errors.New("Parameter num2 is invalid.");
	}

	return num1, num2, nil;
}

func (hdlr *MathHandler) status(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"<html><h1>goci-example is ready!<h1><html>")
}

func (hdlr *MathHandler) addition(w http.ResponseWriter, r *http.Request){
	num1,num2,numErr := getNums(r)
	if numErr == nil {
		fmt.Fprintf(w,"%d",(num1 + num2))
	} else {
		http.Error(w,numErr.Error(),http.StatusBadRequest)
	}
}

func (hdlr *MathHandler) subtraction(w http.ResponseWriter, r *http.Request){
	num1,num2,numErr := getNums(r)
	if numErr == nil {
		fmt.Fprintf(w,"%d",(num1 - num2))
	} else {
		http.Error(w,numErr.Error(),http.StatusBadRequest)
	}
}

func (hdlr *MathHandler) multiplication(w http.ResponseWriter, r *http.Request){
	num1,num2,numErr := getNums(r)
	if numErr == nil {
		fmt.Fprintf(w,"%d",(num1 * num2))
	} else {
		http.Error(w,numErr.Error(),http.StatusBadRequest)
	}
}

func (hdlr *MathHandler) division(w http.ResponseWriter, r *http.Request){
	num1,num2,numErr := getNums(r)
	if numErr == nil {
		fmt.Fprintf(w,"%d",(num1 / num2))
	} else {
		http.Error(w,numErr.Error(),http.StatusBadRequest)
	}
}


