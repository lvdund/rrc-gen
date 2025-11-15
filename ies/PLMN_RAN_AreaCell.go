package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_RAN_AreaCell struct {
	plmn_Identity *PLMN_Identity `optional`
	ran_AreaCells []CellIdentity `lb:1,ub:32,madatory`
}

func (ie *PLMN_RAN_AreaCell) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.plmn_Identity != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.plmn_Identity != nil {
		if err = ie.plmn_Identity.Encode(w); err != nil {
			return utils.WrapError("Encode plmn_Identity", err)
		}
	}
	tmp_ran_AreaCells := utils.NewSequence[*CellIdentity]([]*CellIdentity{}, uper.Constraint{Lb: 1, Ub: 32}, false)
	for _, i := range ie.ran_AreaCells {
		tmp_ran_AreaCells.Value = append(tmp_ran_AreaCells.Value, &i)
	}
	if err = tmp_ran_AreaCells.Encode(w); err != nil {
		return utils.WrapError("Encode ran_AreaCells", err)
	}
	return nil
}

func (ie *PLMN_RAN_AreaCell) Decode(r *uper.UperReader) error {
	var err error
	var plmn_IdentityPresent bool
	if plmn_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if plmn_IdentityPresent {
		ie.plmn_Identity = new(PLMN_Identity)
		if err = ie.plmn_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode plmn_Identity", err)
		}
	}
	tmp_ran_AreaCells := utils.NewSequence[*CellIdentity]([]*CellIdentity{}, uper.Constraint{Lb: 1, Ub: 32}, false)
	fn_ran_AreaCells := func() *CellIdentity {
		return new(CellIdentity)
	}
	if err = tmp_ran_AreaCells.Decode(r, fn_ran_AreaCells); err != nil {
		return utils.WrapError("Decode ran_AreaCells", err)
	}
	ie.ran_AreaCells = []CellIdentity{}
	for _, i := range tmp_ran_AreaCells.Value {
		ie.ran_AreaCells = append(ie.ran_AreaCells, *i)
	}
	return nil
}
