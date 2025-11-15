package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17 struct {
	maxNumNZP_CSI_RS_r17        int64                                                                       `lb:2,ub:8,madatory`
	cSI_Report_mode_r17         MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_cSI_Report_mode_r17  `madatory`
	supportedComboAcrossCCs_r17 []CSI_MultiTRP_SupportedCombinations_r17                                    `lb:1,ub:16,madatory`
	codebookModeNCJT_r17        MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17_codebookModeNCJT_r17 `madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumNZP_CSI_RS_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumNZP_CSI_RS_r17", err)
	}
	if err = ie.cSI_Report_mode_r17.Encode(w); err != nil {
		return utils.WrapError("Encode cSI_Report_mode_r17", err)
	}
	tmp_supportedComboAcrossCCs_r17 := utils.NewSequence[*CSI_MultiTRP_SupportedCombinations_r17]([]*CSI_MultiTRP_SupportedCombinations_r17{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	for _, i := range ie.supportedComboAcrossCCs_r17 {
		tmp_supportedComboAcrossCCs_r17.Value = append(tmp_supportedComboAcrossCCs_r17.Value, &i)
	}
	if err = tmp_supportedComboAcrossCCs_r17.Encode(w); err != nil {
		return utils.WrapError("Encode supportedComboAcrossCCs_r17", err)
	}
	if err = ie.codebookModeNCJT_r17.Encode(w); err != nil {
		return utils.WrapError("Encode codebookModeNCJT_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumNZP_CSI_RS_r17 int64
	if tmp_int_maxNumNZP_CSI_RS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumNZP_CSI_RS_r17", err)
	}
	ie.maxNumNZP_CSI_RS_r17 = tmp_int_maxNumNZP_CSI_RS_r17
	if err = ie.cSI_Report_mode_r17.Decode(r); err != nil {
		return utils.WrapError("Decode cSI_Report_mode_r17", err)
	}
	tmp_supportedComboAcrossCCs_r17 := utils.NewSequence[*CSI_MultiTRP_SupportedCombinations_r17]([]*CSI_MultiTRP_SupportedCombinations_r17{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	fn_supportedComboAcrossCCs_r17 := func() *CSI_MultiTRP_SupportedCombinations_r17 {
		return new(CSI_MultiTRP_SupportedCombinations_r17)
	}
	if err = tmp_supportedComboAcrossCCs_r17.Decode(r, fn_supportedComboAcrossCCs_r17); err != nil {
		return utils.WrapError("Decode supportedComboAcrossCCs_r17", err)
	}
	ie.supportedComboAcrossCCs_r17 = []CSI_MultiTRP_SupportedCombinations_r17{}
	for _, i := range tmp_supportedComboAcrossCCs_r17.Value {
		ie.supportedComboAcrossCCs_r17 = append(ie.supportedComboAcrossCCs_r17, *i)
	}
	if err = ie.codebookModeNCJT_r17.Decode(r); err != nil {
		return utils.WrapError("Decode codebookModeNCJT_r17", err)
	}
	return nil
}
