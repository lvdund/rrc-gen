package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17 struct {
	reducedMIMO_LayersFR2_2_DL_r17 int64 `lb:1,ub:8,madatory`
	reducedMIMO_LayersFR2_2_UL_r17 int64 `lb:1,ub:4,madatory`
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.reducedMIMO_LayersFR2_2_DL_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger reducedMIMO_LayersFR2_2_DL_r17", err)
	}
	if err = w.WriteInteger(ie.reducedMIMO_LayersFR2_2_UL_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger reducedMIMO_LayersFR2_2_UL_r17", err)
	}
	return nil
}

func (ie *MaxMIMO_LayerPreferenceFR2_2_r17_reducedMaxMIMO_LayersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_reducedMIMO_LayersFR2_2_DL_r17 int64
	if tmp_int_reducedMIMO_LayersFR2_2_DL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger reducedMIMO_LayersFR2_2_DL_r17", err)
	}
	ie.reducedMIMO_LayersFR2_2_DL_r17 = tmp_int_reducedMIMO_LayersFR2_2_DL_r17
	var tmp_int_reducedMIMO_LayersFR2_2_UL_r17 int64
	if tmp_int_reducedMIMO_LayersFR2_2_UL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger reducedMIMO_LayersFR2_2_UL_r17", err)
	}
	ie.reducedMIMO_LayersFR2_2_UL_r17 = tmp_int_reducedMIMO_LayersFR2_2_UL_r17
	return nil
}
