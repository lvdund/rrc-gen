package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_RAN_AreaCellList struct {
	Value []PLMN_RAN_AreaCell `lb:1,ub:maxPLMNIdentities,madatory`
}

func (ie *PLMN_RAN_AreaCellList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PLMN_RAN_AreaCell]([]*PLMN_RAN_AreaCell{}, uper.Constraint{Lb: 1, Ub: maxPLMNIdentities}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PLMN_RAN_AreaCellList", err)
	}
	return nil
}

func (ie *PLMN_RAN_AreaCellList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PLMN_RAN_AreaCell]([]*PLMN_RAN_AreaCell{}, uper.Constraint{Lb: 1, Ub: maxPLMNIdentities}, false)
	fn := func() *PLMN_RAN_AreaCell {
		return new(PLMN_RAN_AreaCell)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PLMN_RAN_AreaCellList", err)
	}
	ie.Value = []PLMN_RAN_AreaCell{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
