package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressConfigurationList_r16 struct {
	iab_IP_AddressToAddModList_r16  []IAB_IP_AddressConfiguration_r16 `lb:1,ub:maxIAB_IP_Address_r16,optional`
	iab_IP_AddressToReleaseList_r16 []IAB_IP_AddressIndex_r16         `lb:1,ub:maxIAB_IP_Address_r16,optional`
}

func (ie *IAB_IP_AddressConfigurationList_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.iab_IP_AddressToAddModList_r16) > 0, len(ie.iab_IP_AddressToReleaseList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.iab_IP_AddressToAddModList_r16) > 0 {
		tmp_iab_IP_AddressToAddModList_r16 := utils.NewSequence[*IAB_IP_AddressConfiguration_r16]([]*IAB_IP_AddressConfiguration_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		for _, i := range ie.iab_IP_AddressToAddModList_r16 {
			tmp_iab_IP_AddressToAddModList_r16.Value = append(tmp_iab_IP_AddressToAddModList_r16.Value, &i)
		}
		if err = tmp_iab_IP_AddressToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_IP_AddressToAddModList_r16", err)
		}
	}
	if len(ie.iab_IP_AddressToReleaseList_r16) > 0 {
		tmp_iab_IP_AddressToReleaseList_r16 := utils.NewSequence[*IAB_IP_AddressIndex_r16]([]*IAB_IP_AddressIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		for _, i := range ie.iab_IP_AddressToReleaseList_r16 {
			tmp_iab_IP_AddressToReleaseList_r16.Value = append(tmp_iab_IP_AddressToReleaseList_r16.Value, &i)
		}
		if err = tmp_iab_IP_AddressToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode iab_IP_AddressToReleaseList_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressConfigurationList_r16) Decode(r *uper.UperReader) error {
	var err error
	var iab_IP_AddressToAddModList_r16Present bool
	if iab_IP_AddressToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var iab_IP_AddressToReleaseList_r16Present bool
	if iab_IP_AddressToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if iab_IP_AddressToAddModList_r16Present {
		tmp_iab_IP_AddressToAddModList_r16 := utils.NewSequence[*IAB_IP_AddressConfiguration_r16]([]*IAB_IP_AddressConfiguration_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		fn_iab_IP_AddressToAddModList_r16 := func() *IAB_IP_AddressConfiguration_r16 {
			return new(IAB_IP_AddressConfiguration_r16)
		}
		if err = tmp_iab_IP_AddressToAddModList_r16.Decode(r, fn_iab_IP_AddressToAddModList_r16); err != nil {
			return utils.WrapError("Decode iab_IP_AddressToAddModList_r16", err)
		}
		ie.iab_IP_AddressToAddModList_r16 = []IAB_IP_AddressConfiguration_r16{}
		for _, i := range tmp_iab_IP_AddressToAddModList_r16.Value {
			ie.iab_IP_AddressToAddModList_r16 = append(ie.iab_IP_AddressToAddModList_r16, *i)
		}
	}
	if iab_IP_AddressToReleaseList_r16Present {
		tmp_iab_IP_AddressToReleaseList_r16 := utils.NewSequence[*IAB_IP_AddressIndex_r16]([]*IAB_IP_AddressIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		fn_iab_IP_AddressToReleaseList_r16 := func() *IAB_IP_AddressIndex_r16 {
			return new(IAB_IP_AddressIndex_r16)
		}
		if err = tmp_iab_IP_AddressToReleaseList_r16.Decode(r, fn_iab_IP_AddressToReleaseList_r16); err != nil {
			return utils.WrapError("Decode iab_IP_AddressToReleaseList_r16", err)
		}
		ie.iab_IP_AddressToReleaseList_r16 = []IAB_IP_AddressIndex_r16{}
		for _, i := range tmp_iab_IP_AddressToReleaseList_r16.Value {
			ie.iab_IP_AddressToReleaseList_r16 = append(ie.iab_IP_AddressToReleaseList_r16, *i)
		}
	}
	return nil
}
