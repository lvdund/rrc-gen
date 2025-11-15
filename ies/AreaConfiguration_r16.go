package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AreaConfiguration_r16 struct {
	areaConfig_r16          AreaConfig_r16            `madatory`
	interFreqTargetList_r16 []InterFreqTargetInfo_r16 `lb:1,ub:maxFreq,optional`
}

func (ie *AreaConfiguration_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.interFreqTargetList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.areaConfig_r16.Encode(w); err != nil {
		return utils.WrapError("Encode areaConfig_r16", err)
	}
	if len(ie.interFreqTargetList_r16) > 0 {
		tmp_interFreqTargetList_r16 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		for _, i := range ie.interFreqTargetList_r16 {
			tmp_interFreqTargetList_r16.Value = append(tmp_interFreqTargetList_r16.Value, &i)
		}
		if err = tmp_interFreqTargetList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqTargetList_r16", err)
		}
	}
	return nil
}

func (ie *AreaConfiguration_r16) Decode(r *uper.UperReader) error {
	var err error
	var interFreqTargetList_r16Present bool
	if interFreqTargetList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.areaConfig_r16.Decode(r); err != nil {
		return utils.WrapError("Decode areaConfig_r16", err)
	}
	if interFreqTargetList_r16Present {
		tmp_interFreqTargetList_r16 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		fn_interFreqTargetList_r16 := func() *InterFreqTargetInfo_r16 {
			return new(InterFreqTargetInfo_r16)
		}
		if err = tmp_interFreqTargetList_r16.Decode(r, fn_interFreqTargetList_r16); err != nil {
			return utils.WrapError("Decode interFreqTargetList_r16", err)
		}
		ie.interFreqTargetList_r16 = []InterFreqTargetInfo_r16{}
		for _, i := range tmp_interFreqTargetList_r16.Value {
			ie.interFreqTargetList_r16 = append(ie.interFreqTargetList_r16, *i)
		}
	}
	return nil
}
