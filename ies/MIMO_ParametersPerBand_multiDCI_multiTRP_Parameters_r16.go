package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16 struct {
	overlapPDSCHsFullyFreqTime_r16       *int64                                                                                        `lb:1,ub:2,optional`
	overlapPDSCHsInTimePartiallyFreq_r16 *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_overlapPDSCHsInTimePartiallyFreq_r16 `optional`
	outOfOrderOperationDL_r16            *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16            `optional`
	outOfOrderOperationUL_r16            *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationUL_r16            `optional`
	separateCRS_RateMatching_r16         *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_separateCRS_RateMatching_r16         `optional`
	defaultQCL_PerCORESETPoolIndex_r16   *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_defaultQCL_PerCORESETPoolIndex_r16   `optional`
	maxNumberActivatedTCI_States_r16     *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16     `optional`
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.overlapPDSCHsFullyFreqTime_r16 != nil, ie.overlapPDSCHsInTimePartiallyFreq_r16 != nil, ie.outOfOrderOperationDL_r16 != nil, ie.outOfOrderOperationUL_r16 != nil, ie.separateCRS_RateMatching_r16 != nil, ie.defaultQCL_PerCORESETPoolIndex_r16 != nil, ie.maxNumberActivatedTCI_States_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.overlapPDSCHsFullyFreqTime_r16 != nil {
		if err = w.WriteInteger(*ie.overlapPDSCHsFullyFreqTime_r16, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode overlapPDSCHsFullyFreqTime_r16", err)
		}
	}
	if ie.overlapPDSCHsInTimePartiallyFreq_r16 != nil {
		if err = ie.overlapPDSCHsInTimePartiallyFreq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode overlapPDSCHsInTimePartiallyFreq_r16", err)
		}
	}
	if ie.outOfOrderOperationDL_r16 != nil {
		if err = ie.outOfOrderOperationDL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode outOfOrderOperationDL_r16", err)
		}
	}
	if ie.outOfOrderOperationUL_r16 != nil {
		if err = ie.outOfOrderOperationUL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode outOfOrderOperationUL_r16", err)
		}
	}
	if ie.separateCRS_RateMatching_r16 != nil {
		if err = ie.separateCRS_RateMatching_r16.Encode(w); err != nil {
			return utils.WrapError("Encode separateCRS_RateMatching_r16", err)
		}
	}
	if ie.defaultQCL_PerCORESETPoolIndex_r16 != nil {
		if err = ie.defaultQCL_PerCORESETPoolIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode defaultQCL_PerCORESETPoolIndex_r16", err)
		}
	}
	if ie.maxNumberActivatedTCI_States_r16 != nil {
		if err = ie.maxNumberActivatedTCI_States_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberActivatedTCI_States_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var overlapPDSCHsFullyFreqTime_r16Present bool
	if overlapPDSCHsFullyFreqTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var overlapPDSCHsInTimePartiallyFreq_r16Present bool
	if overlapPDSCHsInTimePartiallyFreq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var outOfOrderOperationDL_r16Present bool
	if outOfOrderOperationDL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var outOfOrderOperationUL_r16Present bool
	if outOfOrderOperationUL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var separateCRS_RateMatching_r16Present bool
	if separateCRS_RateMatching_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var defaultQCL_PerCORESETPoolIndex_r16Present bool
	if defaultQCL_PerCORESETPoolIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberActivatedTCI_States_r16Present bool
	if maxNumberActivatedTCI_States_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if overlapPDSCHsFullyFreqTime_r16Present {
		var tmp_int_overlapPDSCHsFullyFreqTime_r16 int64
		if tmp_int_overlapPDSCHsFullyFreqTime_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode overlapPDSCHsFullyFreqTime_r16", err)
		}
		ie.overlapPDSCHsFullyFreqTime_r16 = &tmp_int_overlapPDSCHsFullyFreqTime_r16
	}
	if overlapPDSCHsInTimePartiallyFreq_r16Present {
		ie.overlapPDSCHsInTimePartiallyFreq_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_overlapPDSCHsInTimePartiallyFreq_r16)
		if err = ie.overlapPDSCHsInTimePartiallyFreq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode overlapPDSCHsInTimePartiallyFreq_r16", err)
		}
	}
	if outOfOrderOperationDL_r16Present {
		ie.outOfOrderOperationDL_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16)
		if err = ie.outOfOrderOperationDL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode outOfOrderOperationDL_r16", err)
		}
	}
	if outOfOrderOperationUL_r16Present {
		ie.outOfOrderOperationUL_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationUL_r16)
		if err = ie.outOfOrderOperationUL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode outOfOrderOperationUL_r16", err)
		}
	}
	if separateCRS_RateMatching_r16Present {
		ie.separateCRS_RateMatching_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_separateCRS_RateMatching_r16)
		if err = ie.separateCRS_RateMatching_r16.Decode(r); err != nil {
			return utils.WrapError("Decode separateCRS_RateMatching_r16", err)
		}
	}
	if defaultQCL_PerCORESETPoolIndex_r16Present {
		ie.defaultQCL_PerCORESETPoolIndex_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_defaultQCL_PerCORESETPoolIndex_r16)
		if err = ie.defaultQCL_PerCORESETPoolIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode defaultQCL_PerCORESETPoolIndex_r16", err)
		}
	}
	if maxNumberActivatedTCI_States_r16Present {
		ie.maxNumberActivatedTCI_States_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16)
		if err = ie.maxNumberActivatedTCI_States_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberActivatedTCI_States_r16", err)
		}
	}
	return nil
}
