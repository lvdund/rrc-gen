package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR struct {
	dummy                                     *CA_ParametersNR_dummy                                     `optional`
	parallelTxSRS_PUCCH_PUSCH                 *CA_ParametersNR_parallelTxSRS_PUCCH_PUSCH                 `optional`
	parallelTxPRACH_SRS_PUCCH_PUSCH           *CA_ParametersNR_parallelTxPRACH_SRS_PUCCH_PUSCH           `optional`
	simultaneousRxTxInterBandCA               *CA_ParametersNR_simultaneousRxTxInterBandCA               `optional`
	simultaneousRxTxSUL                       *CA_ParametersNR_simultaneousRxTxSUL                       `optional`
	diffNumerologyAcrossPUCCH_Group           *CA_ParametersNR_diffNumerologyAcrossPUCCH_Group           `optional`
	diffNumerologyWithinPUCCH_GroupSmallerSCS *CA_ParametersNR_diffNumerologyWithinPUCCH_GroupSmallerSCS `optional`
	supportedNumberTAG                        *CA_ParametersNR_supportedNumberTAG                        `optional`
}

func (ie *CA_ParametersNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dummy != nil, ie.parallelTxSRS_PUCCH_PUSCH != nil, ie.parallelTxPRACH_SRS_PUCCH_PUSCH != nil, ie.simultaneousRxTxInterBandCA != nil, ie.simultaneousRxTxSUL != nil, ie.diffNumerologyAcrossPUCCH_Group != nil, ie.diffNumerologyWithinPUCCH_GroupSmallerSCS != nil, ie.supportedNumberTAG != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.parallelTxSRS_PUCCH_PUSCH != nil {
		if err = ie.parallelTxSRS_PUCCH_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode parallelTxSRS_PUCCH_PUSCH", err)
		}
	}
	if ie.parallelTxPRACH_SRS_PUCCH_PUSCH != nil {
		if err = ie.parallelTxPRACH_SRS_PUCCH_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode parallelTxPRACH_SRS_PUCCH_PUSCH", err)
		}
	}
	if ie.simultaneousRxTxInterBandCA != nil {
		if err = ie.simultaneousRxTxInterBandCA.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxTxInterBandCA", err)
		}
	}
	if ie.simultaneousRxTxSUL != nil {
		if err = ie.simultaneousRxTxSUL.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxTxSUL", err)
		}
	}
	if ie.diffNumerologyAcrossPUCCH_Group != nil {
		if err = ie.diffNumerologyAcrossPUCCH_Group.Encode(w); err != nil {
			return utils.WrapError("Encode diffNumerologyAcrossPUCCH_Group", err)
		}
	}
	if ie.diffNumerologyWithinPUCCH_GroupSmallerSCS != nil {
		if err = ie.diffNumerologyWithinPUCCH_GroupSmallerSCS.Encode(w); err != nil {
			return utils.WrapError("Encode diffNumerologyWithinPUCCH_GroupSmallerSCS", err)
		}
	}
	if ie.supportedNumberTAG != nil {
		if err = ie.supportedNumberTAG.Encode(w); err != nil {
			return utils.WrapError("Encode supportedNumberTAG", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR) Decode(r *uper.UperReader) error {
	var err error
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var parallelTxSRS_PUCCH_PUSCHPresent bool
	if parallelTxSRS_PUCCH_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var parallelTxPRACH_SRS_PUCCH_PUSCHPresent bool
	if parallelTxPRACH_SRS_PUCCH_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousRxTxInterBandCAPresent bool
	if simultaneousRxTxInterBandCAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousRxTxSULPresent bool
	if simultaneousRxTxSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var diffNumerologyAcrossPUCCH_GroupPresent bool
	if diffNumerologyAcrossPUCCH_GroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var diffNumerologyWithinPUCCH_GroupSmallerSCSPresent bool
	if diffNumerologyWithinPUCCH_GroupSmallerSCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedNumberTAGPresent bool
	if supportedNumberTAGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dummyPresent {
		ie.dummy = new(CA_ParametersNR_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	if parallelTxSRS_PUCCH_PUSCHPresent {
		ie.parallelTxSRS_PUCCH_PUSCH = new(CA_ParametersNR_parallelTxSRS_PUCCH_PUSCH)
		if err = ie.parallelTxSRS_PUCCH_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode parallelTxSRS_PUCCH_PUSCH", err)
		}
	}
	if parallelTxPRACH_SRS_PUCCH_PUSCHPresent {
		ie.parallelTxPRACH_SRS_PUCCH_PUSCH = new(CA_ParametersNR_parallelTxPRACH_SRS_PUCCH_PUSCH)
		if err = ie.parallelTxPRACH_SRS_PUCCH_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode parallelTxPRACH_SRS_PUCCH_PUSCH", err)
		}
	}
	if simultaneousRxTxInterBandCAPresent {
		ie.simultaneousRxTxInterBandCA = new(CA_ParametersNR_simultaneousRxTxInterBandCA)
		if err = ie.simultaneousRxTxInterBandCA.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxTxInterBandCA", err)
		}
	}
	if simultaneousRxTxSULPresent {
		ie.simultaneousRxTxSUL = new(CA_ParametersNR_simultaneousRxTxSUL)
		if err = ie.simultaneousRxTxSUL.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxTxSUL", err)
		}
	}
	if diffNumerologyAcrossPUCCH_GroupPresent {
		ie.diffNumerologyAcrossPUCCH_Group = new(CA_ParametersNR_diffNumerologyAcrossPUCCH_Group)
		if err = ie.diffNumerologyAcrossPUCCH_Group.Decode(r); err != nil {
			return utils.WrapError("Decode diffNumerologyAcrossPUCCH_Group", err)
		}
	}
	if diffNumerologyWithinPUCCH_GroupSmallerSCSPresent {
		ie.diffNumerologyWithinPUCCH_GroupSmallerSCS = new(CA_ParametersNR_diffNumerologyWithinPUCCH_GroupSmallerSCS)
		if err = ie.diffNumerologyWithinPUCCH_GroupSmallerSCS.Decode(r); err != nil {
			return utils.WrapError("Decode diffNumerologyWithinPUCCH_GroupSmallerSCS", err)
		}
	}
	if supportedNumberTAGPresent {
		ie.supportedNumberTAG = new(CA_ParametersNR_supportedNumberTAG)
		if err = ie.supportedNumberTAG.Decode(r); err != nil {
			return utils.WrapError("Decode supportedNumberTAG", err)
		}
	}
	return nil
}
