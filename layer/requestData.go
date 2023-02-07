package layer

import "github.com/TheDeveloper10/rem"

type IRequestBody interface {
	// validate data
	Validate() error
}

type IRequestURL interface {
	// validate `rem.KeyValue` then return an error or fill structure
	Process(rem.KeyValue) error
}

type IRequestQuery interface {
	// validate `rem.KeyValues` then return an error or fill structure
	Process(rem.KeyValues) error
}