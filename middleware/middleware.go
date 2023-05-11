package middleware

import (
	"myapp/data"

	"github.com/hilsonxhero/napoleon"
)

type Middleware struct {
	App    *napoleon.Napoleon
	Models data.Models
}
