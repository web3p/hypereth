// Copyright 2018 AMIS Technologies
// This file is part of the hypereth library.
//
// The hypereth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The hypereth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the hypereth library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "context"

// StartMining starts mining operation.
func (ec *Client) StartMining(ctx context.Context) error {
	var r []byte
	return ec.c.CallContext(ctx, &r, "miner_start", nil)
}

// StopMining stops mining.
func (ec *Client) StopMining(ctx context.Context) error {
	return ec.c.CallContext(ctx, nil, "miner_stop", nil)
}
