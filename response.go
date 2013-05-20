package elevator

import (
    "fmt"
    "bytes"
    "github.com/ugorji/go-msgpack"
)

type Response struct {
    Status   int
    Err_code int
    Err_msg  string
    Data     []string
}

// String represents the Response as a normalized string
func (r *Response) String() string {
    return fmt.Sprintf("<Response status:%d err_code:%d err_msg:%s data:%s",
        r.Status, r.Err_code, r.Err_msg, r.Data)
}

// NewResponse returns a pointer to a brand new allocated Response
func NewResponse(status int, err_code int, err_msg string, data []string) *Response {
    return &Response{
        Status:   status,
        Err_code: err_code,
        Err_msg:  err_msg,
        Data:     data,
    }
}

// NewSuccessResponse returns a pointer to a brand
// new allocated succesfull Response
func NewSuccessResponse(data []string) *Response {
    return &Response{
        Status: SUCCESS_STATUS,
        Data:   data,
    }
}

// NewFailureResponse returns a pointer to a brand
// new allocated failure Response
func NewFailureResponse(err_code int, err_msg string) *Response {
    return &Response{
        Status:   FAILURE_STATUS,
        Err_code: err_code,
        Err_msg:  err_msg,
    }
}

// ToArray transforms a Response to an array-like []interface{}
func (r *Response) ToArray() []interface{} {
    var response []interface{}

    response = append(response, r.Status, r.Err_code, r.Err_msg)

    for _, d := range r.Data {
        response = append(response, string(d))
    }

    return response
}

// PackInto method fulfills serializes the Response
// into a msgpacked response message
func (r *Response) PackInto(buffer *bytes.Buffer) error {
    enc := msgpack.NewEncoder(buffer)
    err := enc.Encode(r.ToArray())
    if err != nil {
        return err
    }

    return nil
}