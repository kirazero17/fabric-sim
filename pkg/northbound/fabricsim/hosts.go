// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package fabricsim

import (
	"context"
	simapi "github.com/onosproject/onos-api/go/onos/fabricsim"
)

// GetHosts returns list of all simulated hosts
func (s *Server) GetHosts(ctx context.Context, request *simapi.GetHostsRequest) (*simapi.GetHostsResponse, error) {
	sims := s.Simulation.GetHostSimulators()
	hosts := make([]*simapi.Host, 0, len(sims))
	for _, sim := range sims {
		hosts = append(hosts, sim.Host)
	}
	return &simapi.GetHostsResponse{Hosts: hosts}, nil
}

// GetHost returns the specified simulated host
func (s *Server) GetHost(ctx context.Context, request *simapi.GetHostRequest) (*simapi.GetHostResponse, error) {
	sim, err := s.Simulation.GetHostSimulator(request.ID)
	if err != nil {
		return nil, err
	}
	return &simapi.GetHostResponse{Host: sim.Host}, nil
}

// AddHost creates and registers the specified simulated host
func (s *Server) AddHost(ctx context.Context, request *simapi.AddHostRequest) (*simapi.AddHostResponse, error) {
	if _, err := s.Simulation.AddHostSimulator(request.Host); err != nil {
		return nil, err
	}
	return &simapi.AddHostResponse{}, nil
}

// RemoveHost removes the specified simulated host
func (s *Server) RemoveHost(ctx context.Context, request *simapi.RemoveHostRequest) (*simapi.RemoveHostResponse, error) {
	if err := s.Simulation.RemoveHostSimulator(request.ID); err != nil {
		return nil, err
	}
	return &simapi.RemoveHostResponse{}, nil
}