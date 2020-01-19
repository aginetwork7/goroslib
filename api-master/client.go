package api_master

import (
	"fmt"

	"github.com/aler9/goroslib/xmlrpc"
)

type Client struct {
	url      string
	callerId string
}

func NewClient(host string, port uint16, callerId string) (*Client, error) {
	return &Client{
		url:      fmt.Sprintf("http://%s:%d/", host, port),
		callerId: callerId,
	}, nil
}

func (c *Client) Close() error {
	return nil
}

func (c *Client) GetSystemState() (*SystemState, error) {
	req := getSystemStateReq{
		c.callerId,
	}
	var res getSystemStateRes
	err := xmlrpc.Client(c.url, "getSystemState", req, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 1 {
		return nil, fmt.Errorf("server returned an error (%d): %s", res.Code, res.StatusMessage)
	}

	return &res.State, nil
}

func (c *Client) GetTopicTypes() ([]TopicType, error) {
	req := getTopicTypesReq{
		c.callerId,
	}
	var res getTopicTypesRes
	err := xmlrpc.Client(c.url, "getTopicTypes", req, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 1 {
		return nil, fmt.Errorf("server returned an error (%d): %s", res.Code, res.StatusMessage)
	}

	return res.Types, nil
}

func (c *Client) lookup(method string, name string) (string, error) {
	req := lookupReq{
		c.callerId,
		name,
	}
	var res lookupRes
	err := xmlrpc.Client(c.url, method, req, &res)
	if err != nil {
		return "", err
	}

	if res.Code != 1 {
		return "", fmt.Errorf("server returned an error (%d): %s", res.Code, res.StatusMessage)
	}

	return res.Uri, nil
}

func (c *Client) LookupNode(name string) (string, error) {
	return c.lookup("lookupNode", name)
}

func (c *Client) LookupService(name string) (string, error) {
	return c.lookup("lookupService", name)
}

func (c *Client) Register(method string, topic string, topicType string, callerUrl string) ([]string, error) {
	req := registerReq{
		c.callerId,
		topic,
		topicType,
		callerUrl,
	}
	var res registerRes
	err := xmlrpc.Client(c.url, method, req, &res)
	if err != nil {
		return nil, err
	}

	if res.Code != 1 {
		return nil, fmt.Errorf("server returned an error (%d): %s", res.Code, res.StatusMessage)
	}

	return res.Uris, nil
}

func (c *Client) unregister(method string, topic string, callerUrl string) error {
	req := unregisterReq{
		c.callerId,
		topic,
		callerUrl,
	}
	var res unregisterRes
	err := xmlrpc.Client(c.url, method, req, &res)
	if err != nil {
		return err
	}

	if res.NumUnregistered == 0 {
		return fmt.Errorf("unregister failed")
	}

	return nil
}

func (c *Client) RegisterSubscriber(topic string, topicType string, callerUrl string) ([]string, error) {
	return c.Register("registerSubscriber", topic, topicType, callerUrl)
}

func (c *Client) UnregisterSubscriber(topic string, callerUrl string) error {
	return c.unregister("unregisterSubscriber", topic, callerUrl)
}

func (c *Client) RegisterPublisher(topic string, topicType string, callerUrl string) ([]string, error) {
	return c.Register("registerPublisher", topic, topicType, callerUrl)
}

func (c *Client) UnregisterPublisher(topic string, callerUrl string) error {
	return c.unregister("unregisterPublisher", topic, callerUrl)
}

func (c *Client) RegisterService(service string, serviceUrl string, callerUrl string) error {
	req := serviceRegisterReq{
		c.callerId,
		service,
		serviceUrl,
		callerUrl,
	}
	var res serviceRegisterRes
	err := xmlrpc.Client(c.url, "registerService", req, &res)
	if err != nil {
		return err
	}

	if res.Code != 1 {
		return fmt.Errorf("server returned an error (%d): %s", res.Code, res.StatusMessage)
	}

	return nil
}

func (c *Client) UnregisterService(service string, serviceUrl string) error {
	req := serviceUnregisterReq{
		c.callerId,
		service,
		serviceUrl,
	}
	var res serviceUnregisterRes
	err := xmlrpc.Client(c.url, "unregisterService", req, &res)
	if err != nil {
		return err
	}

	if res.NumUnregistered == 0 {
		return fmt.Errorf("unregister failed")
	}

	return nil
}
