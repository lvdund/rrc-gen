package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersEUTRA_v1570 struct {
	dl_1024QAM_TotalWeightedLayers *int64 `lb:0,ub:10,optional`
}

func (ie *CA_ParametersEUTRA_v1570) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_1024QAM_TotalWeightedLayers != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_1024QAM_TotalWeightedLayers != nil {
		if err = w.WriteInteger(*ie.dl_1024QAM_TotalWeightedLayers, &uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode dl_1024QAM_TotalWeightedLayers", err)
		}
	}
	return nil
}

func (ie *CA_ParametersEUTRA_v1570) Decode(r *uper.UperReader) error {
	var err error
	var dl_1024QAM_TotalWeightedLayersPresent bool
	if dl_1024QAM_TotalWeightedLayersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_1024QAM_TotalWeightedLayersPresent {
		var tmp_int_dl_1024QAM_TotalWeightedLayers int64
		if tmp_int_dl_1024QAM_TotalWeightedLayers, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode dl_1024QAM_TotalWeightedLayers", err)
		}
		ie.dl_1024QAM_TotalWeightedLayers = &tmp_int_dl_1024QAM_TotalWeightedLayers
	}
	return nil
}
