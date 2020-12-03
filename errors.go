package events

import "errors"

//ErrEventDataDoesNotImplementProperties error event data does not implement properties
var ErrEventDataDoesNotImplementProperties = errors.New("event data does not implement properties")
