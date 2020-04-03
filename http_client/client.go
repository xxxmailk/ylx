package http_client

import (
	"encoding/json"
	"encoding/xml"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func NewHttpClient(logEntry *logrus.Entry, headerFunc func(request *fasthttp.Request, log *logrus.Entry), preFunc func(sendBody interface{}, log *logrus.Entry)) *HttpClient {
	c := new(HttpClient)
	c.logger = logEntry
	c.headerFunc = headerFunc
	c.preFunc = preFunc
	return c
}

type HttpClient struct {
	header     []byte
	logger     *logrus.Entry
	headerFunc func(request *fasthttp.Request, log *logrus.Entry)
	preFunc    func(sendBody interface{}, log *logrus.Entry)
}

func (p *HttpClient) GetJson(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "GET", sendBody, rs)
}

func (p *HttpClient) PostJson(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "POST", sendBody, rs)
}

func (p *HttpClient) DeleteJson(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "DELETE", sendBody, rs)
}

func (p *HttpClient) PutJson(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "PUT", sendBody, rs)
}

func (p *HttpClient) GetXml(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "GET", sendBody, rs)
}

func (p *HttpClient) PostXml(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "POST", sendBody, rs)
}

func (p *HttpClient) DeleteXml(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "DELETE", sendBody, rs)
}

func (p *HttpClient) PutXml(url string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	return p.jsonClient(url, "PUT", sendBody, rs)
}

func (p *HttpClient) xmlClient(url, method string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	p.logger.Debugf("sending %s json request to %s", method, url)
	p.logger.Tracef("sending %s json request to %s request body: %v", method, url, sendBody)
	if p.preFunc != nil {
		p.preFunc(sendBody, p.logger)
	}
	req.Header.SetContentType("text/xml")
	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	if p.headerFunc != nil {
		p.headerFunc(req, p.logger)
	}
	requestBody, err := xml.Marshal(sendBody)
	if err != nil {
		p.logger.Errorf("marshal request xml failed, %w", err)
		return 500, nil
	}
	req.SetBody(requestBody)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		p.logger.Errorf("%s request: %s failed, %s", method, url, err.Error())
		return 500, nil
	}
	return resp.StatusCode(), xml.Unmarshal(resp.Body(), rs)
}

func (p *HttpClient) jsonClient(url, method string, sendBody interface{}, rs interface{}) (statCode int, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	p.logger.Debugf("sending %s json request to %s", method, url)
	p.logger.Tracef("sending %s json request to %s request body: %v", method, url, sendBody)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	if p.headerFunc != nil {
		p.headerFunc(req, p.logger)
	}

	if p.preFunc != nil {
		p.preFunc(sendBody, p.logger)
	}

	requestBody, err := json.Marshal(sendBody)
	if err != nil {
		p.logger.Errorf("marshal request json failed, %w", err)
		return 500, nil
	}
	req.SetBody(requestBody)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		p.logger.Errorf("%s request: %s failed, %s", method, url, err.Error())
		return 500, nil
	}
	return resp.StatusCode(), json.Unmarshal(resp.Body(), rs)
}
