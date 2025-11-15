package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1700_IEs struct {
	candidateCellListCPC_r17                  *CandidateCellListCPC_r17                                          `optional`
	twoPHRModeMCG_r17                         *CG_ConfigInfo_v1700_IEs_twoPHRModeMCG_r17                         `optional`
	lowMobilityEvaluationConnectedInPCell_r17 *CG_ConfigInfo_v1700_IEs_lowMobilityEvaluationConnectedInPCell_r17 `optional`
	nonCriticalExtension                      *CG_ConfigInfo_v1730_IEs                                           `optional`
}

func (ie *CG_ConfigInfo_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.candidateCellListCPC_r17 != nil, ie.twoPHRModeMCG_r17 != nil, ie.lowMobilityEvaluationConnectedInPCell_r17 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.candidateCellListCPC_r17 != nil {
		if err = ie.candidateCellListCPC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode candidateCellListCPC_r17", err)
		}
	}
	if ie.twoPHRModeMCG_r17 != nil {
		if err = ie.twoPHRModeMCG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode twoPHRModeMCG_r17", err)
		}
	}
	if ie.lowMobilityEvaluationConnectedInPCell_r17 != nil {
		if err = ie.lowMobilityEvaluationConnectedInPCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode lowMobilityEvaluationConnectedInPCell_r17", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var candidateCellListCPC_r17Present bool
	if candidateCellListCPC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPHRModeMCG_r17Present bool
	if twoPHRModeMCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lowMobilityEvaluationConnectedInPCell_r17Present bool
	if lowMobilityEvaluationConnectedInPCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if candidateCellListCPC_r17Present {
		ie.candidateCellListCPC_r17 = new(CandidateCellListCPC_r17)
		if err = ie.candidateCellListCPC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode candidateCellListCPC_r17", err)
		}
	}
	if twoPHRModeMCG_r17Present {
		ie.twoPHRModeMCG_r17 = new(CG_ConfigInfo_v1700_IEs_twoPHRModeMCG_r17)
		if err = ie.twoPHRModeMCG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode twoPHRModeMCG_r17", err)
		}
	}
	if lowMobilityEvaluationConnectedInPCell_r17Present {
		ie.lowMobilityEvaluationConnectedInPCell_r17 = new(CG_ConfigInfo_v1700_IEs_lowMobilityEvaluationConnectedInPCell_r17)
		if err = ie.lowMobilityEvaluationConnectedInPCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode lowMobilityEvaluationConnectedInPCell_r17", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1730_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
