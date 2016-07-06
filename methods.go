package meli

//Get
func (c *Client) Get(path string, params map[string]string) ([]byte, error) {
	return execute("GET", path, params, nil)
}

//Post
func (c *Client) Post(path string, body interface{}, params map[string]string) ([]byte, error) {
	return execute("POST", path, params, body)
}

//Put
func (c *Client) Put(path string, body interface{}, params map[string]string) ([]byte, error) {
	return execute("PUT", path, params, body)
}

//Delete
func (c *Client) Delete(path string, params map[string]string) ([]byte, error) {
	return execute("DELETE", path, params, nil)
}
