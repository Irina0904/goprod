package main

type JsonResponse struct {
	Data       any `json:"data"`
	StatusCode int `json:"status_code"`
}