package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AreaConfiguration_v1700 struct {
	areaConfig_r17          *AreaConfig_r16           `optional`
	interFreqTargetList_r17 []InterFreqTargetInfo_r16 `lb:1,ub:maxFreq,optional`
}

func (ie *AreaConfiguration_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.areaConfig_r17 != nil, len(ie.interFreqTargetList_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.areaConfig_r17 != nil {
		if err = ie.areaConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode areaConfig_r17", err)
		}
	}
	if len(ie.interFreqTargetList_r17) > 0 {
		tmp_interFreqTargetList_r17 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		for _, i := range ie.interFreqTargetList_r17 {
			tmp_interFreqTargetList_r17.Value = append(tmp_interFreqTargetList_r17.Value, &i)
		}
		if err = tmp_interFreqTargetList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode interFreqTargetList_r17", err)
		}
	}
	return nil
}

func (ie *AreaConfiguration_v1700) Decode(r *uper.UperReader) error {
	var err error
	var areaConfig_r17Present bool
	if areaConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var interFreqTargetList_r17Present bool
	if interFreqTargetList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if areaConfig_r17Present {
		ie.areaConfig_r17 = new(AreaConfig_r16)
		if err = ie.areaConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode areaConfig_r17", err)
		}
	}
	if interFreqTargetList_r17Present {
		tmp_interFreqTargetList_r17 := utils.NewSequence[*InterFreqTargetInfo_r16]([]*InterFreqTargetInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
		fn_interFreqTargetList_r17 := func() *InterFreqTargetInfo_r16 {
			return new(InterFreqTargetInfo_r16)
		}
		if err = tmp_interFreqTargetList_r17.Decode(r, fn_interFreqTargetList_r17); err != nil {
			return utils.WrapError("Decode interFreqTargetList_r17", err)
		}
		ie.interFreqTargetList_r17 = []InterFreqTargetInfo_r16{}
		for _, i := range tmp_interFreqTargetList_r17.Value {
			ie.interFreqTargetList_r17 = append(ie.interFreqTargetList_r17, *i)
		}
	}
	return nil
}
