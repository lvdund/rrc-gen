package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosMeasGapPreConfigToAddModList_r17 struct {
	Value []PosGapConfig_r17 `lb:1,ub:maxNrofPreConfigPosGapId_r17,madatory`
}

func (ie *PosMeasGapPreConfigToAddModList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PosGapConfig_r17]([]*PosGapConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPreConfigPosGapId_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PosMeasGapPreConfigToAddModList_r17", err)
	}
	return nil
}

func (ie *PosMeasGapPreConfigToAddModList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PosGapConfig_r17]([]*PosGapConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPreConfigPosGapId_r17}, false)
	fn := func() *PosGapConfig_r17 {
		return new(PosGapConfig_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PosMeasGapPreConfigToAddModList_r17", err)
	}
	ie.Value = []PosGapConfig_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
