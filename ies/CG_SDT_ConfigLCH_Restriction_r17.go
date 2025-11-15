package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_SDT_ConfigLCH_Restriction_r17 struct {
	logicalChannelIdentity_r17      LogicalChannelIdentity                                            `madatory`
	configuredGrantType1Allowed_r17 *CG_SDT_ConfigLCH_Restriction_r17_configuredGrantType1Allowed_r17 `optional`
	allowedCG_List_r17              []ConfiguredGrantConfigIndexMAC_r16                               `lb:0,ub:maxNrofConfiguredGrantConfigMAC_1_r16,optional`
}

func (ie *CG_SDT_ConfigLCH_Restriction_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.configuredGrantType1Allowed_r17 != nil, len(ie.allowedCG_List_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.logicalChannelIdentity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode logicalChannelIdentity_r17", err)
	}
	if ie.configuredGrantType1Allowed_r17 != nil {
		if err = ie.configuredGrantType1Allowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode configuredGrantType1Allowed_r17", err)
		}
	}
	if len(ie.allowedCG_List_r17) > 0 {
		tmp_allowedCG_List_r17 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		for _, i := range ie.allowedCG_List_r17 {
			tmp_allowedCG_List_r17.Value = append(tmp_allowedCG_List_r17.Value, &i)
		}
		if err = tmp_allowedCG_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode allowedCG_List_r17", err)
		}
	}
	return nil
}

func (ie *CG_SDT_ConfigLCH_Restriction_r17) Decode(r *uper.UperReader) error {
	var err error
	var configuredGrantType1Allowed_r17Present bool
	if configuredGrantType1Allowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedCG_List_r17Present bool
	if allowedCG_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.logicalChannelIdentity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode logicalChannelIdentity_r17", err)
	}
	if configuredGrantType1Allowed_r17Present {
		ie.configuredGrantType1Allowed_r17 = new(CG_SDT_ConfigLCH_Restriction_r17_configuredGrantType1Allowed_r17)
		if err = ie.configuredGrantType1Allowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode configuredGrantType1Allowed_r17", err)
		}
	}
	if allowedCG_List_r17Present {
		tmp_allowedCG_List_r17 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		fn_allowedCG_List_r17 := func() *ConfiguredGrantConfigIndexMAC_r16 {
			return new(ConfiguredGrantConfigIndexMAC_r16)
		}
		if err = tmp_allowedCG_List_r17.Decode(r, fn_allowedCG_List_r17); err != nil {
			return utils.WrapError("Decode allowedCG_List_r17", err)
		}
		ie.allowedCG_List_r17 = []ConfiguredGrantConfigIndexMAC_r16{}
		for _, i := range tmp_allowedCG_List_r17.Value {
			ie.allowedCG_List_r17 = append(ie.allowedCG_List_r17, *i)
		}
	}
	return nil
}
