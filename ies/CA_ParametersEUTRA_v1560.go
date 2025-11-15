package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersEUTRA_v1560 struct {
	fd_MIMO_TotalWeightedLayers *int64 `lb:2,ub:128,optional`
}

func (ie *CA_ParametersEUTRA_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fd_MIMO_TotalWeightedLayers != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.fd_MIMO_TotalWeightedLayers != nil {
		if err = w.WriteInteger(*ie.fd_MIMO_TotalWeightedLayers, &uper.Constraint{Lb: 2, Ub: 128}, false); err != nil {
			return utils.WrapError("Encode fd_MIMO_TotalWeightedLayers", err)
		}
	}
	return nil
}

func (ie *CA_ParametersEUTRA_v1560) Decode(r *uper.UperReader) error {
	var err error
	var fd_MIMO_TotalWeightedLayersPresent bool
	if fd_MIMO_TotalWeightedLayersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if fd_MIMO_TotalWeightedLayersPresent {
		var tmp_int_fd_MIMO_TotalWeightedLayers int64
		if tmp_int_fd_MIMO_TotalWeightedLayers, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode fd_MIMO_TotalWeightedLayers", err)
		}
		ie.fd_MIMO_TotalWeightedLayers = &tmp_int_fd_MIMO_TotalWeightedLayers
	}
	return nil
}
