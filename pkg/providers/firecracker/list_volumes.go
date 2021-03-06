package firecracker

import (
	"github.com/solo-io/unik/pkg/types"
)

func (p *FirecrackerProvider) ListVolumes() ([]*types.Volume, error) {
	volumes := []*types.Volume{}
	for _, volume := range p.state.GetVolumes() {
		volumes = append(volumes, volume)
	}
	return volumes, nil
}
