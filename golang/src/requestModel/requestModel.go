package requestModel

import "net/http"
import "io/ioutil"

type RequestModel struct  {
    urlParam string
    method string
    requestId string
    userAgent string
}

func ParseFromRequest(model *RequestModel, request *http.Request) {
   url,_ := ioutil.ReadAll(request.Body)
   model.urlParam = string(url) 
   model.method = request.Method
   model.userAgent = request.Header.Get("User-Agent")
   model.requestId = "1234" 
}

func (model *RequestModel) Init() {
   model.urlParam = "" 
   model.method = "" 
   model.requestId = "" 
   model.userAgent = "" 
}
