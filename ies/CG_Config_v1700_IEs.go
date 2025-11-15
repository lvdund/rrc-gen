package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1700_IEs struct {
	candidateCellInfoListCPC_r17 *CandidateCellInfoListCPC_r17          `optional`
	twoPHRModeSCG_r17            *CG_Config_v1700_IEs_twoPHRModeSCG_r17 `optional`
	nonCriticalExtension         *CG_Config_v1730_IEs                   `optional`
}

func (ie *CG_Config_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.candidateCellInfoListCPC_r17 != nil, ie.twoPHRModeSCG_r17 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.candidateCellInfoListCPC_r17 != nil {
		if err = ie.candidateCellInfoListCPC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode candidateCellInfoListCPC_r17", err)
		}
	}
	if ie.twoPHRModeSCG_r17 != nil {
		if err = ie.twoPHRModeSCG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode twoPHRModeSCG_r17", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var candidateCellInfoListCPC_r17Present bool
	if candidateCellInfoListCPC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPHRModeSCG_r17Present bool
	if twoPHRModeSCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if candidateCellInfoListCPC_r17Present {
		ie.candidateCellInfoListCPC_r17 = new(CandidateCellInfoListCPC_r17)
		if err = ie.candidateCellInfoListCPC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode candidateCellInfoListCPC_r17", err)
		}
	}
	if twoPHRModeSCG_r17Present {
		ie.twoPHRModeSCG_r17 = new(CG_Config_v1700_IEs_twoPHRModeSCG_r17)
		if err = ie.twoPHRModeSCG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode twoPHRModeSCG_r17", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1730_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
