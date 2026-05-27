/*
GoLive http frame work this is the Context file that holds
varias function to send read and detect errors
most of these ,methods will need a http status code 200 404 etc...
*/

// TODO: make error and valid send json make the key for the map the param

package goLive

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

//Context custom struct to have http types tied to methods rather than passed to functions
type Context struct{
	Writer http.ResponseWriter
	Request *http.Request
	params map[string]string
}

// SendJSON used for sending map/json data
func (c *Context) SendJSON(status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	return json.NewEncoder(c.Writer).Encode(data)
}

// PrettyJSON is jut SendJSON with indenting used for sending map/json data
func (c *Context) PrettyJSON(status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	encoder := json.NewEncoder(c.Writer)
	encoder.SetIndent("", " ")
	return encoder.Encode(data)
}

// ReadJSON reads json
func (c *Context) ReadJSON(data any) error {
	defer c.Request.Body.Close()

	if c.Request.Header.Get("Content-Type") != "application/json" {
		return errors.New("invalid content type")
	}

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // optional: prevents extra fields in JSON
	return decoder.Decode(data)
}

// ReadForm is used to read forms Request ie "username=David"
func (c *Context) ReadForm(data string) ( string ){
 err := c.Request.ParseForm()
	if err != nil {
		http.Error(c.Writer, "Bad form data", http.StatusBadRequest)
		return "Error" 
	}

	resp := c.Request.FormValue(data)

	return resp 
}

//SendSTRING sends a simple text only string to the client good for fast tests json key is "response 
func (c *Context) SendSTRING(status int, data string)error{
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(data))
	return err
}

//Valid Sends a json response with custom message json key is "response"
func (c *Context)Valid(status int, validMsg string)error{
	resp := map[string]string{
		"response": validMsg,
	}
	if status < 200 || status > 399{
		return ErrInvalidRedirectCode
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	err := json.NewEncoder(c.Writer).Encode(resp)
	return err
}

//Error sends error into the request for more custom and simple error handling in the http methods
func (c *Context) Error(status int, errorMsg string) error {
	// _, err := c.Writer.Write([]byte(errorMsg))
	resp := map[string]string{
		"response": errorMsg,
	}
	if status < 400 || status > 599{
		return ErrInvalidRedirectCode
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	err := json.NewEncoder(c.Writer).Encode(resp)
	return err
}

//Redirect redirects the clients request to the needed url or other path
func (c *Context) Redirect(status int, redirectURL string) error {
	if status < 300 || status > 308 {
		return ErrInvalidRedirectCode
	}
	http.Redirect(c.Writer, c.Request, redirectURL, status)
	c.Writer.WriteHeader(status)
	return nil
}

//Param gets the value of path url param
func (c *Context)Param(key string)string{
	if c.params == nil{
		return ""
	}
	return c.params[key]
}

//QueryGet gets the Query from the url ?=
func (c *Context)QueryGet(data string)string{
	foundQuery := c.Request.URL.Query().Get(data)
	return foundQuery
}

func (c *Context)ReciveFile(){
	// TODO: read name
}

// SendFile FIX: need more tweaking not working as intened
func (c *Context)SendFile(filepath string)error{
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.Error(c.Writer, "File not found", http.StatusNotFound)
		return err
	}
	http.ServeFile(c.Writer, c.Request, filepath)
	return nil
}


