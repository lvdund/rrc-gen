package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PTRS_Config_r16 struct {
	sl_PTRS_FreqDensity_r16 []int64                                   `lb:2,ub:2,e_lb:0,e_ub:0,optional`
	sl_PTRS_TimeDensity_r16 []int64                                   `lb:3,ub:3,e_lb:0,e_ub:0,optional`
	sl_PTRS_RE_Offset_r16   *SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16 `optional`
}

func (ie *SL_PTRS_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_PTRS_FreqDensity_r16) > 0, len(ie.sl_PTRS_TimeDensity_r16) > 0, ie.sl_PTRS_RE_Offset_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_PTRS_FreqDensity_r16) > 0 {
		tmp_sl_PTRS_FreqDensity_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		for _, i := range ie.sl_PTRS_FreqDensity_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_sl_PTRS_FreqDensity_r16.Value = append(tmp_sl_PTRS_FreqDensity_r16.Value, &tmp_ie)
		}
		if err = tmp_sl_PTRS_FreqDensity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PTRS_FreqDensity_r16", err)
		}
	}
	if len(ie.sl_PTRS_TimeDensity_r16) > 0 {
		tmp_sl_PTRS_TimeDensity_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.sl_PTRS_TimeDensity_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 0}, false)
			tmp_sl_PTRS_TimeDensity_r16.Value = append(tmp_sl_PTRS_TimeDensity_r16.Value, &tmp_ie)
		}
		if err = tmp_sl_PTRS_TimeDensity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PTRS_TimeDensity_r16", err)
		}
	}
	if ie.sl_PTRS_RE_Offset_r16 != nil {
		if err = ie.sl_PTRS_RE_Offset_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PTRS_RE_Offset_r16", err)
		}
	}
	return nil
}

func (ie *SL_PTRS_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PTRS_FreqDensity_r16Present bool
	if sl_PTRS_FreqDensity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PTRS_TimeDensity_r16Present bool
	if sl_PTRS_TimeDensity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PTRS_RE_Offset_r16Present bool
	if sl_PTRS_RE_Offset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PTRS_FreqDensity_r16Present {
		tmp_sl_PTRS_FreqDensity_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 2, Ub: 2}, false)
		fn_sl_PTRS_FreqDensity_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_sl_PTRS_FreqDensity_r16.Decode(r, fn_sl_PTRS_FreqDensity_r16); err != nil {
			return utils.WrapError("Decode sl_PTRS_FreqDensity_r16", err)
		}
		ie.sl_PTRS_FreqDensity_r16 = []int64{}
		for _, i := range tmp_sl_PTRS_FreqDensity_r16.Value {
			ie.sl_PTRS_FreqDensity_r16 = append(ie.sl_PTRS_FreqDensity_r16, int64(i.Value))
		}
	}
	if sl_PTRS_TimeDensity_r16Present {
		tmp_sl_PTRS_TimeDensity_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_sl_PTRS_TimeDensity_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 0}, false)
			return &ie
		}
		if err = tmp_sl_PTRS_TimeDensity_r16.Decode(r, fn_sl_PTRS_TimeDensity_r16); err != nil {
			return utils.WrapError("Decode sl_PTRS_TimeDensity_r16", err)
		}
		ie.sl_PTRS_TimeDensity_r16 = []int64{}
		for _, i := range tmp_sl_PTRS_TimeDensity_r16.Value {
			ie.sl_PTRS_TimeDensity_r16 = append(ie.sl_PTRS_TimeDensity_r16, int64(i.Value))
		}
	}
	if sl_PTRS_RE_Offset_r16Present {
		ie.sl_PTRS_RE_Offset_r16 = new(SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16)
		if err = ie.sl_PTRS_RE_Offset_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PTRS_RE_Offset_r16", err)
		}
	}
	return nil
}
