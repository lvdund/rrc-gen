package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RAN_AreaConfig struct {
	trackingAreaCode TrackingAreaCode `madatory`
	ran_AreaCodeList []RAN_AreaCode   `lb:1,ub:32,optional`
}

func (ie *RAN_AreaConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.ran_AreaCodeList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.trackingAreaCode.Encode(w); err != nil {
		return utils.WrapError("Encode trackingAreaCode", err)
	}
	if len(ie.ran_AreaCodeList) > 0 {
		tmp_ran_AreaCodeList := utils.NewSequence[*RAN_AreaCode]([]*RAN_AreaCode{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		for _, i := range ie.ran_AreaCodeList {
			tmp_ran_AreaCodeList.Value = append(tmp_ran_AreaCodeList.Value, &i)
		}
		if err = tmp_ran_AreaCodeList.Encode(w); err != nil {
			return utils.WrapError("Encode ran_AreaCodeList", err)
		}
	}
	return nil
}

func (ie *RAN_AreaConfig) Decode(r *uper.UperReader) error {
	var err error
	var ran_AreaCodeListPresent bool
	if ran_AreaCodeListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.trackingAreaCode.Decode(r); err != nil {
		return utils.WrapError("Decode trackingAreaCode", err)
	}
	if ran_AreaCodeListPresent {
		tmp_ran_AreaCodeList := utils.NewSequence[*RAN_AreaCode]([]*RAN_AreaCode{}, uper.Constraint{Lb: 1, Ub: 32}, false)
		fn_ran_AreaCodeList := func() *RAN_AreaCode {
			return new(RAN_AreaCode)
		}
		if err = tmp_ran_AreaCodeList.Decode(r, fn_ran_AreaCodeList); err != nil {
			return utils.WrapError("Decode ran_AreaCodeList", err)
		}
		ie.ran_AreaCodeList = []RAN_AreaCode{}
		for _, i := range tmp_ran_AreaCodeList.Value {
			ie.ran_AreaCodeList = append(ie.ran_AreaCodeList, *i)
		}
	}
	return nil
}
