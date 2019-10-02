package ovo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type CoreGateway struct {
	Client Client
}

func (gateway *CoreGateway) PushToPay(req *PushToPayRequest, lang string) (PushToPayResponse, error) {
	path := gateway.Client.APIEnvType.String() + "/pos"
	headers := gateway.PopulateHeader()
	reqParam := gateway.PopulateOVOParam(req)
	resp := PushToPayResponse{}
	jsonReq, _ := json.Marshal(reqParam)

	err := gateway.Client.Call("POST", path, headers, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Checkout : ", err)
		return resp, err
	}

	if resp.ResponseCode != "00" {
		message, _ := GetDetailResponse(lang, resp.ResponseCode)
		return resp, errors.New(message)
	}

	return resp, nil
}

func (gateway *CoreGateway) ReversalPushToPay(req *PushToPayRequest, lang string) (PushToPayResponse, error) {
	path := gateway.Client.APIEnvType.String() + "/pos"
	headers := gateway.PopulateHeader()
	reqParam := gateway.PopulateOVOParam(req)
	resp := PushToPayResponse{}
	jsonReq, _ := json.Marshal(reqParam)

	err := gateway.Client.Call("POST", path, headers, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Checkout : ", err)
		return resp, err
	}

	if resp.ResponseCode != "00" {
		message, _ := GetDetailResponse(lang, resp.ResponseCode)
		return resp, errors.New(message)
	}

	return resp, nil
}

func (gateway *CoreGateway) VoidPushToPay(req *PushToPayRequest, lang string) (PushToPayResponse, error) {
	path := gateway.Client.APIEnvType.String() + "/pos"
	headers := gateway.PopulateHeader()
	reqParam := gateway.PopulateOVOParam(req)
	resp := PushToPayResponse{}
	jsonReq, _ := json.Marshal(reqParam)

	err := gateway.Client.Call("POST", path, headers, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Checkout : ", err)
		return resp, err
	}

	if resp.ResponseCode != "00" {
		message, _ := GetDetailResponse(lang, resp.ResponseCode)
		return resp, errors.New(message)
	}

	return resp, nil
}

func (gateway *CoreGateway) CheckPaymentStatus(req *PushToPayRequest, lang string) (PushToPayResponse, error) {
	path := gateway.Client.APIEnvType.String() + "/pos"
	headers := gateway.PopulateHeader()
	reqParam := gateway.PopulateOVOParam(req)
	resp := PushToPayResponse{}
	jsonReq, _ := json.Marshal(reqParam)

	err := gateway.Client.Call("POST", path, headers, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Checkout : ", err)
		return resp, err
	}

	if resp.ResponseCode != "00" {
		message, _ := GetDetailResponse(lang, resp.ResponseCode)
		return resp, errors.New(message)
	}

	return resp, nil
}

func (gateway *CoreGateway) PopulateHeader() []Header {
	head := []Header{
		Header{Key: "app-id", Value: gateway.Client.AppID},
		Header{Key: "random", Value: strconv.FormatInt(time.Now().Unix(), 10)},
		Header{Key: "hmac", Value: EncodeHMACSHA256(gateway.Client.AppID, strconv.FormatInt(time.Now().Unix(), 10), gateway.Client.Key)},
	}

	return head
}

func (gateway *CoreGateway) PopulateOVOParam(req *PushToPayRequest) *PushToPayRequest {
	req.TID = gateway.Client.TID
	req.MID = gateway.Client.MID
	req.MerchantID = gateway.Client.MerchantID
	req.StoreCode = gateway.Client.StoreCode
	req.AppSource = gateway.Client.AppSource

	return req
}
