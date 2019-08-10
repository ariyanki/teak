package test

import (
	"net/http"
)

var (
	//BillersResp biller response
	BillersResp = []map[string]interface{}{
		{"id": 1, "label": "ariyanki", "description": "test biller telkomsel"},
	}
	//BillersVarsResp biller variable response
	BillersVarsResp = []map[string]interface{}{
		{"name": "biller:ariyanki", "value": `{"Message":"Hello World"}`},
	}
	//InvalidBillersVars InvalidBillersVars
	InvalidBillersVars = `{"ID":,"Message":}`
	//CacheData cache data
	CacheData = map[string]string{
		"biller:biller_label": `{
		"production":false,
		"debug":true,
		"response_code_mapping":{
		   "00":[
			  {
				 "biller_rc":"0000",
				 "message":"Success"
			  }
		   ],
		   "20":[
			  {
				 "biller_rc":"9006",
				 "message":"Nomorexpiredataubelumaktif"
			  }
		   ],
		   "22":[
			  {
				 "biller_rc":"9012",
				 "message":"DuplicateTransactionID/RejectFromBilling"
			  },
			  {
				 "biller_rc":"",
				 "message":"Duplicateentry“idx”key1"
			  }
		   ]
		},
		"product_mapping":{
		   "1":{
			  "price":4900,
			  "product_id":"PKREG60WS"
		   }
		},
		"api_config":{
		   "connection_type":"http",
		   "content_type":"json",
		   "auth":{
			  "type":"Bearer Token",
			  "request":{
				 "url":"",
				 "method":"post",
				 "header":{
					"Content-Type":"application/json"
				 },
				 "timeout":60,
				 "body":"{\"username\": \"andi123\",\"password\": \"asd123\"}"
			  },
			  "response_map":{
				 "token":"payload.token"
			  }
		   },
		   "purchase":{
			  "mobile":{
				 "request":{
					"url":"",
					"method":"post",
					"header":{
					   "Content-Type":"application/json"
					},
					"timeout":60,
					"body":"{\"product_id\":\"b5af8f85-119e-450e-ae52-27a3cc6a15c5\",\"data\":{\"transaction_id\":[transaction_id],\"customer_number\":\"[customer_id]\"}}"
				 },
				 "response_map":{
					"remote_transaction_id":"data.remote_transaction_id",
					"serial_number":"data.serial_number",
					"biller_rescode":"data.rc_biller",
					"message":"data.biller_message"
				 }
			  }
		   },
		   "stock":{
			  "request":{
				 "url":"",
				 "method":"post",
				 "header":{
					"Content-Type":"application/json"
				 },
				 "timeout":60,
				 "body":"{\"product_id\":\"b5af8f85-119e-450e-ae52-27a3cc6a15c5\",\"data\":{\"transaction_id\":[transaction_id],\"customer_number\":\"[customer_id]\"}}"
			  },
			  "response_map":{
				 "remote_transaction_id":"data.remote_transaction_id",
				 "serial_number":"data.serial_number",
				 "biller_rescode":"data.rc_biller",
				 "message":"data.biller_message"
			  }
		   }
		}
	 }`,
		"product:1": `{
		"id":1,
		"created_at":"2018-06-26T19:05:15+07:00",
		"updated_at":"2019-05-04T10:10:13+07:00",
		"deleted_at":null,
		"code":"1",
		"label":"Tri Rp50,000",
		"product_type":{
		   "id":1,
		   "label":"mobile"
		},
		"type_id":1,
		"operator":{
		   "id":1,
		   "label":"tri",
		   "code":""
		},
		"operator_id":1,
		"nominal":50000,
		"price":48000,
		"status":true,
		"billers":[

		]
	 }`,
	 "partner:1":`"{\"id\":1,\"created_at\":\"2018-04-30T17:58:45+07:00\",\"updated_at\":\"2019-08-01T20:49:21+07:00\",\"deleted_at\":null,\"username\":\"tester\",\"password\":\"$2a$08$uoUNRdFg/cF6Y9MitOt6VuYHQ/1UqfblUrB1tGYsRPg93WCaOieXa\",\"name\":\"tester\",\"callback_url\":\"http://localhost/test.php\",\"package\":{\"id\":0,\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"deleted_at\":null,\"label\":\"\",\"blacklist\":false,\"products\":null},\"package_id\":84,\"email\":\"\",\"payment_type\":\"deposit\",\"is_active\":true,\"hosts\":[{\"partner_id\":1,\"ip_address\":\"175.103.43.138\"},{\"partner_id\":1,\"ip_address\":\"202.4.189.18\"},{\"partner_id\":1,\"ip_address\":\"175.103.43.138\"}],\"banks\":[],\"balance\":{\"id\":1,\"created_at\":\"2018-04-30T18:08:40+07:00\",\"updated_at\":\"2019-05-27T13:17:51+07:00\",\"deleted_at\":null,\"partner_id\":1,\"amount\":2263000,\"histories\":null}}"`,
	}
	//MessageDataStr message data
	MessageDataStr = `{
		"command_type":"purchase",
		"transaction_id":6855945,
		"transaction_biller_id":6855950,
		"remote_transaction_id":"",
		"remote_product_id":"1",
		"customer_id":"085363783000",
		"partner_id":6,
		"partner_name":"abc",
		"product_type_label":"mobile",
		"product_id":1,
		"biller_id":1,
		"biller_label":"biller_label",
		"counter":1,
		"rawdata":""
	 }`
	//RespBody RespBody
	RespBody = `{
		"transaction_id": 6855945,
		"customer_number": "085363783000",
		"status": "sukses",
		"data": {
		  "serial_number": "0051003619319900",
		  "remote_transaction_id": "110931978314",
		  "rc_biller": "0000",
		  "biller_message": "Pengisian Voucher sebesar 25000 ke nomor 085363783000 pada tanggal 18/07/2019 09:54:44 telah berhasil dengan no ref <0051003619319900>"
		}
	  }`
	//RespKraken RespKraken
	RespKraken = `{
		"transaction_id": 6855945,
		"transaction_biller_id": 6855950,
		"remote_transaction_id": "12456",
		"remote_product_id": "PKREG60WS",
		"customer_id": "085363783000",
		"response_code": "00",
		"price": 4900,
		"amount": 50000,
		"data": {
			"serial_number": "0051003619319900",
			"biller_rescode": "0000",
			"message": "Pengisian Voucher sebesar 25000 ke nomor 085363783000 pada tanggal 18/07/2019 09:54:44 telah berhasil dengan no ref <0051003619319900>",
			"rawdata": "{\"Message\":\"Service Unavailable\"}\n"
		}
	}`
	//DummyHandlerLogin DummyHandlerLogin
	DummyHandlerLogin = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body := `{
				"code": 200,
				"payload": {
					"token": "a.b.c"
				}
			}`
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		}
	}
	//DummyHandlerProcess DummyHandlerProcess
	DummyHandlerProcess = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body := RespBody
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		}
	}
	//DummyHandlerProcessNot200 DummyHandlerProcessNot200
	DummyHandlerProcessNot200 = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.WriteHeader(http.StatusForbidden)
		}
	}
	//DummyHandlerProcessReadBodyError DummyHandlerProcessReadBodyError
	DummyHandlerProcessReadBodyError = func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.WriteHeader(http.StatusOK)
		}
	}
)
