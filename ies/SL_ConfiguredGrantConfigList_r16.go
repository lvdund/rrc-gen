package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfiguredGrantConfigList_r16 struct {
	sl_ConfiguredGrantConfigToReleaseList_r16 []SL_ConfigIndexCG_r16         `lb:1,ub:maxNrofCG_SL_r16,optional`
	sl_ConfiguredGrantConfigToAddModList_r16  []SL_ConfiguredGrantConfig_r16 `lb:1,ub:maxNrofCG_SL_r16,optional`
}

func (ie *SL_ConfiguredGrantConfigList_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_ConfiguredGrantConfigToReleaseList_r16) > 0, len(ie.sl_ConfiguredGrantConfigToAddModList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_ConfiguredGrantConfigToReleaseList_r16) > 0 {
		tmp_sl_ConfiguredGrantConfigToReleaseList_r16 := utils.NewSequence[*SL_ConfigIndexCG_r16]([]*SL_ConfigIndexCG_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCG_SL_r16}, false)
		for _, i := range ie.sl_ConfiguredGrantConfigToReleaseList_r16 {
			tmp_sl_ConfiguredGrantConfigToReleaseList_r16.Value = append(tmp_sl_ConfiguredGrantConfigToReleaseList_r16.Value, &i)
		}
		if err = tmp_sl_ConfiguredGrantConfigToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ConfiguredGrantConfigToReleaseList_r16", err)
		}
	}
	if len(ie.sl_ConfiguredGrantConfigToAddModList_r16) > 0 {
		tmp_sl_ConfiguredGrantConfigToAddModList_r16 := utils.NewSequence[*SL_ConfiguredGrantConfig_r16]([]*SL_ConfiguredGrantConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCG_SL_r16}, false)
		for _, i := range ie.sl_ConfiguredGrantConfigToAddModList_r16 {
			tmp_sl_ConfiguredGrantConfigToAddModList_r16.Value = append(tmp_sl_ConfiguredGrantConfigToAddModList_r16.Value, &i)
		}
		if err = tmp_sl_ConfiguredGrantConfigToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ConfiguredGrantConfigToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *SL_ConfiguredGrantConfigList_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ConfiguredGrantConfigToReleaseList_r16Present bool
	if sl_ConfiguredGrantConfigToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ConfiguredGrantConfigToAddModList_r16Present bool
	if sl_ConfiguredGrantConfigToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ConfiguredGrantConfigToReleaseList_r16Present {
		tmp_sl_ConfiguredGrantConfigToReleaseList_r16 := utils.NewSequence[*SL_ConfigIndexCG_r16]([]*SL_ConfigIndexCG_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCG_SL_r16}, false)
		fn_sl_ConfiguredGrantConfigToReleaseList_r16 := func() *SL_ConfigIndexCG_r16 {
			return new(SL_ConfigIndexCG_r16)
		}
		if err = tmp_sl_ConfiguredGrantConfigToReleaseList_r16.Decode(r, fn_sl_ConfiguredGrantConfigToReleaseList_r16); err != nil {
			return utils.WrapError("Decode sl_ConfiguredGrantConfigToReleaseList_r16", err)
		}
		ie.sl_ConfiguredGrantConfigToReleaseList_r16 = []SL_ConfigIndexCG_r16{}
		for _, i := range tmp_sl_ConfiguredGrantConfigToReleaseList_r16.Value {
			ie.sl_ConfiguredGrantConfigToReleaseList_r16 = append(ie.sl_ConfiguredGrantConfigToReleaseList_r16, *i)
		}
	}
	if sl_ConfiguredGrantConfigToAddModList_r16Present {
		tmp_sl_ConfiguredGrantConfigToAddModList_r16 := utils.NewSequence[*SL_ConfiguredGrantConfig_r16]([]*SL_ConfiguredGrantConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCG_SL_r16}, false)
		fn_sl_ConfiguredGrantConfigToAddModList_r16 := func() *SL_ConfiguredGrantConfig_r16 {
			return new(SL_ConfiguredGrantConfig_r16)
		}
		if err = tmp_sl_ConfiguredGrantConfigToAddModList_r16.Decode(r, fn_sl_ConfiguredGrantConfigToAddModList_r16); err != nil {
			return utils.WrapError("Decode sl_ConfiguredGrantConfigToAddModList_r16", err)
		}
		ie.sl_ConfiguredGrantConfigToAddModList_r16 = []SL_ConfiguredGrantConfig_r16{}
		for _, i := range tmp_sl_ConfiguredGrantConfigToAddModList_r16.Value {
			ie.sl_ConfiguredGrantConfigToAddModList_r16 = append(ie.sl_ConfiguredGrantConfigToAddModList_r16, *i)
		}
	}
	return nil
}
