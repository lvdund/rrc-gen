package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PosMeasGapPreConfigToReleaseList_r17 struct {
	Value []MeasPosPreConfigGapId_r17 `lb:1,ub:maxNrofPreConfigPosGapId_r17,madatory`
}

func (ie *PosMeasGapPreConfigToReleaseList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasPosPreConfigGapId_r17]([]*MeasPosPreConfigGapId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPreConfigPosGapId_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PosMeasGapPreConfigToReleaseList_r17", err)
	}
	return nil
}

func (ie *PosMeasGapPreConfigToReleaseList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasPosPreConfigGapId_r17]([]*MeasPosPreConfigGapId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPreConfigPosGapId_r17}, false)
	fn := func() *MeasPosPreConfigGapId_r17 {
		return new(MeasPosPreConfigGapId_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PosMeasGapPreConfigToReleaseList_r17", err)
	}
	ie.Value = []MeasPosPreConfigGapId_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
