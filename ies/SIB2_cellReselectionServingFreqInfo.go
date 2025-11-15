package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_cellReselectionServingFreqInfo struct {
	s_NonIntraSearchP          *ReselectionThreshold       `optional`
	s_NonIntraSearchQ          *ReselectionThresholdQ      `optional`
	threshServingLowP          ReselectionThreshold        `madatory`
	threshServingLowQ          *ReselectionThresholdQ      `optional`
	cellReselectionPriority    CellReselectionPriority     `madatory`
	cellReselectionSubPriority *CellReselectionSubPriority `optional`
}

func (ie *SIB2_cellReselectionServingFreqInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.s_NonIntraSearchP != nil, ie.s_NonIntraSearchQ != nil, ie.threshServingLowQ != nil, ie.cellReselectionSubPriority != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.s_NonIntraSearchP != nil {
		if err = ie.s_NonIntraSearchP.Encode(w); err != nil {
			return utils.WrapError("Encode s_NonIntraSearchP", err)
		}
	}
	if ie.s_NonIntraSearchQ != nil {
		if err = ie.s_NonIntraSearchQ.Encode(w); err != nil {
			return utils.WrapError("Encode s_NonIntraSearchQ", err)
		}
	}
	if err = ie.threshServingLowP.Encode(w); err != nil {
		return utils.WrapError("Encode threshServingLowP", err)
	}
	if ie.threshServingLowQ != nil {
		if err = ie.threshServingLowQ.Encode(w); err != nil {
			return utils.WrapError("Encode threshServingLowQ", err)
		}
	}
	if err = ie.cellReselectionPriority.Encode(w); err != nil {
		return utils.WrapError("Encode cellReselectionPriority", err)
	}
	if ie.cellReselectionSubPriority != nil {
		if err = ie.cellReselectionSubPriority.Encode(w); err != nil {
			return utils.WrapError("Encode cellReselectionSubPriority", err)
		}
	}
	return nil
}

func (ie *SIB2_cellReselectionServingFreqInfo) Decode(r *uper.UperReader) error {
	var err error
	var s_NonIntraSearchPPresent bool
	if s_NonIntraSearchPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var s_NonIntraSearchQPresent bool
	if s_NonIntraSearchQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var threshServingLowQPresent bool
	if threshServingLowQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellReselectionSubPriorityPresent bool
	if cellReselectionSubPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if s_NonIntraSearchPPresent {
		ie.s_NonIntraSearchP = new(ReselectionThreshold)
		if err = ie.s_NonIntraSearchP.Decode(r); err != nil {
			return utils.WrapError("Decode s_NonIntraSearchP", err)
		}
	}
	if s_NonIntraSearchQPresent {
		ie.s_NonIntraSearchQ = new(ReselectionThresholdQ)
		if err = ie.s_NonIntraSearchQ.Decode(r); err != nil {
			return utils.WrapError("Decode s_NonIntraSearchQ", err)
		}
	}
	if err = ie.threshServingLowP.Decode(r); err != nil {
		return utils.WrapError("Decode threshServingLowP", err)
	}
	if threshServingLowQPresent {
		ie.threshServingLowQ = new(ReselectionThresholdQ)
		if err = ie.threshServingLowQ.Decode(r); err != nil {
			return utils.WrapError("Decode threshServingLowQ", err)
		}
	}
	if err = ie.cellReselectionPriority.Decode(r); err != nil {
		return utils.WrapError("Decode cellReselectionPriority", err)
	}
	if cellReselectionSubPriorityPresent {
		ie.cellReselectionSubPriority = new(CellReselectionSubPriority)
		if err = ie.cellReselectionSubPriority.Decode(r); err != nil {
			return utils.WrapError("Decode cellReselectionSubPriority", err)
		}
	}
	return nil
}
