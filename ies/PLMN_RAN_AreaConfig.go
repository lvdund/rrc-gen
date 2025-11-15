package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_RAN_AreaConfig struct {
	plmn_Identity *PLMN_Identity   `optional`
	ran_Area      []RAN_AreaConfig `lb:1,ub:16,madatory`
}

func (ie *PLMN_RAN_AreaConfig) Encode(w *uper.UperWriter) error {
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
	tmp_ran_Area := utils.NewSequence[*RAN_AreaConfig]([]*RAN_AreaConfig{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	for _, i := range ie.ran_Area {
		tmp_ran_Area.Value = append(tmp_ran_Area.Value, &i)
	}
	if err = tmp_ran_Area.Encode(w); err != nil {
		return utils.WrapError("Encode ran_Area", err)
	}
	return nil
}

func (ie *PLMN_RAN_AreaConfig) Decode(r *uper.UperReader) error {
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
	tmp_ran_Area := utils.NewSequence[*RAN_AreaConfig]([]*RAN_AreaConfig{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	fn_ran_Area := func() *RAN_AreaConfig {
		return new(RAN_AreaConfig)
	}
	if err = tmp_ran_Area.Decode(r, fn_ran_Area); err != nil {
		return utils.WrapError("Decode ran_Area", err)
	}
	ie.ran_Area = []RAN_AreaConfig{}
	for _, i := range tmp_ran_Area.Value {
		ie.ran_Area = append(ie.ran_Area, *i)
	}
	return nil
}
