package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSSCH_Config_r16 struct {
	sl_PSSCH_DMRS_TimePatternList_r16 []int64                             `lb:1,ub:3,e_lb:2,e_ub:4,optional`
	sl_BetaOffsets2ndSCI_r16          []SL_BetaOffsets_r16                `lb:4,ub:4,optional`
	sl_Scaling_r16                    *SL_PSSCH_Config_r16_sl_Scaling_r16 `optional`
}

func (ie *SL_PSSCH_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_PSSCH_DMRS_TimePatternList_r16) > 0, len(ie.sl_BetaOffsets2ndSCI_r16) > 0, ie.sl_Scaling_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_PSSCH_DMRS_TimePatternList_r16) > 0 {
		tmp_sl_PSSCH_DMRS_TimePatternList_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		for _, i := range ie.sl_PSSCH_DMRS_TimePatternList_r16 {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 2, Ub: 4}, false)
			tmp_sl_PSSCH_DMRS_TimePatternList_r16.Value = append(tmp_sl_PSSCH_DMRS_TimePatternList_r16.Value, &tmp_ie)
		}
		if err = tmp_sl_PSSCH_DMRS_TimePatternList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSSCH_DMRS_TimePatternList_r16", err)
		}
	}
	if len(ie.sl_BetaOffsets2ndSCI_r16) > 0 {
		tmp_sl_BetaOffsets2ndSCI_r16 := utils.NewSequence[*SL_BetaOffsets_r16]([]*SL_BetaOffsets_r16{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		for _, i := range ie.sl_BetaOffsets2ndSCI_r16 {
			tmp_sl_BetaOffsets2ndSCI_r16.Value = append(tmp_sl_BetaOffsets2ndSCI_r16.Value, &i)
		}
		if err = tmp_sl_BetaOffsets2ndSCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_BetaOffsets2ndSCI_r16", err)
		}
	}
	if ie.sl_Scaling_r16 != nil {
		if err = ie.sl_Scaling_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Scaling_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSSCH_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PSSCH_DMRS_TimePatternList_r16Present bool
	if sl_PSSCH_DMRS_TimePatternList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_BetaOffsets2ndSCI_r16Present bool
	if sl_BetaOffsets2ndSCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Scaling_r16Present bool
	if sl_Scaling_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PSSCH_DMRS_TimePatternList_r16Present {
		tmp_sl_PSSCH_DMRS_TimePatternList_r16 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 3}, false)
		fn_sl_PSSCH_DMRS_TimePatternList_r16 := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 2, Ub: 4}, false)
			return &ie
		}
		if err = tmp_sl_PSSCH_DMRS_TimePatternList_r16.Decode(r, fn_sl_PSSCH_DMRS_TimePatternList_r16); err != nil {
			return utils.WrapError("Decode sl_PSSCH_DMRS_TimePatternList_r16", err)
		}
		ie.sl_PSSCH_DMRS_TimePatternList_r16 = []int64{}
		for _, i := range tmp_sl_PSSCH_DMRS_TimePatternList_r16.Value {
			ie.sl_PSSCH_DMRS_TimePatternList_r16 = append(ie.sl_PSSCH_DMRS_TimePatternList_r16, int64(i.Value))
		}
	}
	if sl_BetaOffsets2ndSCI_r16Present {
		tmp_sl_BetaOffsets2ndSCI_r16 := utils.NewSequence[*SL_BetaOffsets_r16]([]*SL_BetaOffsets_r16{}, uper.Constraint{Lb: 4, Ub: 4}, false)
		fn_sl_BetaOffsets2ndSCI_r16 := func() *SL_BetaOffsets_r16 {
			return new(SL_BetaOffsets_r16)
		}
		if err = tmp_sl_BetaOffsets2ndSCI_r16.Decode(r, fn_sl_BetaOffsets2ndSCI_r16); err != nil {
			return utils.WrapError("Decode sl_BetaOffsets2ndSCI_r16", err)
		}
		ie.sl_BetaOffsets2ndSCI_r16 = []SL_BetaOffsets_r16{}
		for _, i := range tmp_sl_BetaOffsets2ndSCI_r16.Value {
			ie.sl_BetaOffsets2ndSCI_r16 = append(ie.sl_BetaOffsets2ndSCI_r16, *i)
		}
	}
	if sl_Scaling_r16Present {
		ie.sl_Scaling_r16 = new(SL_PSSCH_Config_r16_sl_Scaling_r16)
		if err = ie.sl_Scaling_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Scaling_r16", err)
		}
	}
	return nil
}
