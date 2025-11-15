package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17 struct {
	maxNumPeriodicSRS_r17          int64 `lb:1,ub:8,madatory`
	maxNumAperiodicSRS_r17         int64 `lb:1,ub:8,madatory`
	maxNumSP_SRS_r17               int64 `lb:0,ub:8,madatory`
	numSRS_ResourcePerCC_r17       int64 `lb:1,ub:16,madatory`
	numSRS_ResourceNonCodebook_r17 int64 `lb:1,ub:2,madatory`
}

func (ie *MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumPeriodicSRS_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumPeriodicSRS_r17", err)
	}
	if err = w.WriteInteger(ie.maxNumAperiodicSRS_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumAperiodicSRS_r17", err)
	}
	if err = w.WriteInteger(ie.maxNumSP_SRS_r17, &uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumSP_SRS_r17", err)
	}
	if err = w.WriteInteger(ie.numSRS_ResourcePerCC_r17, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteInteger numSRS_ResourcePerCC_r17", err)
	}
	if err = w.WriteInteger(ie.numSRS_ResourceNonCodebook_r17, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger numSRS_ResourceNonCodebook_r17", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumPeriodicSRS_r17 int64
	if tmp_int_maxNumPeriodicSRS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumPeriodicSRS_r17", err)
	}
	ie.maxNumPeriodicSRS_r17 = tmp_int_maxNumPeriodicSRS_r17
	var tmp_int_maxNumAperiodicSRS_r17 int64
	if tmp_int_maxNumAperiodicSRS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumAperiodicSRS_r17", err)
	}
	ie.maxNumAperiodicSRS_r17 = tmp_int_maxNumAperiodicSRS_r17
	var tmp_int_maxNumSP_SRS_r17 int64
	if tmp_int_maxNumSP_SRS_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumSP_SRS_r17", err)
	}
	ie.maxNumSP_SRS_r17 = tmp_int_maxNumSP_SRS_r17
	var tmp_int_numSRS_ResourcePerCC_r17 int64
	if tmp_int_numSRS_ResourcePerCC_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadInteger numSRS_ResourcePerCC_r17", err)
	}
	ie.numSRS_ResourcePerCC_r17 = tmp_int_numSRS_ResourcePerCC_r17
	var tmp_int_numSRS_ResourceNonCodebook_r17 int64
	if tmp_int_numSRS_ResourceNonCodebook_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger numSRS_ResourceNonCodebook_r17", err)
	}
	ie.numSRS_ResourceNonCodebook_r17 = tmp_int_numSRS_ResourceNonCodebook_r17
	return nil
}
