package cryptcheck

import "log"

// debug displays only if fDebug is set
func (c *Client) debug(str string, a ...interface{}) {
	if c.level >= 2 {
		log.Printf(str, a...)
	}
}

// debug displays only if fVerbose is set
func (c *Client) verbose(str string, a ...interface{}) {
	if c.level >= 1 {
		log.Printf(str, a...)
	}
}
