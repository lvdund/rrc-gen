package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NTN_NeighCellConfigList_r17 struct {
	Value []NTN_NeighCellConfig_r17 `lb:1,ub:maxCellNTN_r17,madatory`
}

func (ie *NTN_NeighCellConfigList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*NTN_NeighCellConfig_r17]([]*NTN_NeighCellConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxCellNTN_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NTN_NeighCellConfigList_r17", err)
	}
	return nil
}

func (ie *NTN_NeighCellConfigList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*NTN_NeighCellConfig_r17]([]*NTN_NeighCellConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxCellNTN_r17}, false)
	fn := func() *NTN_NeighCellConfig_r17 {
		return new(NTN_NeighCellConfig_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NTN_NeighCellConfigList_r17", err)
	}
	ie.Value = []NTN_NeighCellConfig_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
