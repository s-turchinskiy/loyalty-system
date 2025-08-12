package postgresql

import (
	"context"
	"encoding/json"
)

func (p *PostgreSQL) Ping(ctx context.Context) ([]byte, error) {

	err := p.db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(p.db.Stats(), "", "   ")

}
