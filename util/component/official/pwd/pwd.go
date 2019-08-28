package pwd

import (
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"hedu-basic/util/ecode"
	"io/ioutil"
	"net/http"
	"strings"
)

type Envelope struct {
	Body Body `xml:"Body"`
}

type Body struct {
	ChangePasswordResponse ChangePasswordResponse `xml:"ChangePasswordResponse"`
}

type ChangePasswordResponse struct {
	ChangePasswordResult int `xml:"ChangePasswordResult"`
}

// 修改密码
// 服务来自http://218.198.80.19/service.asmx?op=ChangePasswordFailed
func ChangePassWord(username string, password string) int {
	b := buildRequestBody(username, password)
	request, _ := http.NewRequest(http.MethodPost, "http://218.198.80.19/service.asmx", strings.NewReader(b))
	request.Header.Set("Content-Type", "text/xml; charset=utf-8")
	request.Header.Set("SOAPAction", `http://tempuri.org/zfxk/Service/ChangePassword`)
	response, err := new(http.Client).Do(request)
	if err != nil {
		logrus.WithField("error", err).Warnln("official service unavailability")
		return ecode.OfficialServiceUnavailability
	}
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	var state = new(Envelope)
	xml.Unmarshal(result, state)
	if state.Body.ChangePasswordResponse.ChangePasswordResult != 0 {
		return ecode.ChangePasswordFailed
	}
	return ecode.NoError
}

// 构造请求体
func buildRequestBody(username, password string) string {

	return `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <ChangePassword xmlns="http://tempuri.org/zfxk/Service">
      <UserType>学生</UserType>
      <Account>` + username + `</Account>
      <Newpassword>` + password + `</Newpassword>
    </ChangePassword>
  </soap:Body>
</soap:Envelope>`
}
