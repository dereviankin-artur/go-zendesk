package mock

import (
	"github.com/dereviankin-artur/go-zendesk/zendesk"
)

var _ zendesk.API = (*Client)(nil)
