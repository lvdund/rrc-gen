package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17 struct {
	minBeamApplicationTime_r17  MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17_minBeamApplicationTime_r17 `madatory`
	maxActivatedDL_TCIPerCC_r17 int64                                                                                `lb:2,ub:8,madatory`
	maxActivatedUL_TCIPerCC_r17 int64                                                                                `lb:2,ub:8,madatory`
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.minBeamApplicationTime_r17.Encode(w); err != nil {
		return utils.WrapError("Encode minBeamApplicationTime_r17", err)
	}
	if err = w.WriteInteger(ie.maxActivatedDL_TCIPerCC_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxActivatedDL_TCIPerCC_r17", err)
	}
	if err = w.WriteInteger(ie.maxActivatedUL_TCIPerCC_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxActivatedUL_TCIPerCC_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.minBeamApplicationTime_r17.Decode(r); err != nil {
		return utils.WrapError("Decode minBeamApplicationTime_r17", err)
	}
	var tmp_int_maxActivatedDL_TCIPerCC_r17 int64
	if tmp_int_maxActivatedDL_TCIPerCC_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxActivatedDL_TCIPerCC_r17", err)
	}
	ie.maxActivatedDL_TCIPerCC_r17 = tmp_int_maxActivatedDL_TCIPerCC_r17
	var tmp_int_maxActivatedUL_TCIPerCC_r17 int64
	if tmp_int_maxActivatedUL_TCIPerCC_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxActivatedUL_TCIPerCC_r17", err)
	}
	ie.maxActivatedUL_TCIPerCC_r17 = tmp_int_maxActivatedUL_TCIPerCC_r17
	return nil
}
