package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17_codebookMode_NCJT_r17_Enum_mode1     uper.Enumerated = 0
	CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17_codebookMode_NCJT_r17_Enum_mode1And2 uper.Enumerated = 1
)

type CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17_codebookMode_NCJT_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17_codebookMode_NCJT_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17_codebookMode_NCJT_r17", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17_codebookMode_NCJT_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CA_ParametersNR_v1700_mTRP_CSI_EnhancementPerBC_r17_codebookMode_NCJT_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
