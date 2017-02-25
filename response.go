package milter

/* Response represents commant sent from milter client */
type Response interface {
	Response() *Message
	Continue() bool
}

/* SimpleResponse interface to define list of pre-defined responses */
type SimpleResponse byte

/* Response returns a Message object reference */
func (r SimpleResponse) Response() *Message {
	return &Message{byte(r), nil}
}

/* Continue to process milter messages only if current code is Continue */
func (r SimpleResponse) Continue() bool {
	return byte(r) == Continue
}

/* Define standard responses with no data */
const (
	RespAccept   = SimpleResponse(Accept)
	RespContinue = SimpleResponse(Continue)
	RespDiscard  = SimpleResponse(Discard)
	RespReject   = SimpleResponse(Reject)
	RespTempFail = SimpleResponse(TempFail)
)

/* CustomResponse is a response object used to send data back to MTA */
type CustomResponse struct {
	Code byte
	Data []byte
}

/* Response returns message object with data */
func (c *CustomResponse) Response() *Message {
	return &Message{c.Code, c.Data}
}

/* Continue returns false if milter chain should be stopped, true otherwise */
func (c *CustomResponse) Continue() bool {
	for _, q := range []byte{Accept, Discard, Reject, TempFail} {
		if c.Code == q {
			return false
		}
	}
	return true
}

/* NewResponse generates a new CustomRespanse suitable for WritePacket */
func NewResponse(code byte, data []byte) *CustomResponse {
	return &CustomResponse{code, data}
}
