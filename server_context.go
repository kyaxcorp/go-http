package server

import (
	"context"

	"github.com/kyaxcorp/go-helper/_context"
)

func (s *Server) SetContext(ctx context.Context) {
	if ctx == nil {
		ctx = _context.GetDefaultContext()
	}
	s.parentCtx = ctx
}
