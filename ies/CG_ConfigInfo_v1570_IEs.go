package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1570_IEs struct {
	sftdFrequencyList_NR    *SFTD_FrequencyList_NR    `optional`
	sftdFrequencyList_EUTRA *SFTD_FrequencyList_EUTRA `optional`
	nonCriticalExtension    *CG_ConfigInfo_v1590_IEs  `optional`
}

func (ie *CG_ConfigInfo_v1570_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sftdFrequencyList_NR != nil, ie.sftdFrequencyList_EUTRA != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sftdFrequencyList_NR != nil {
		if err = ie.sftdFrequencyList_NR.Encode(w); err != nil {
			return utils.WrapError("Encode sftdFrequencyList_NR", err)
		}
	}
	if ie.sftdFrequencyList_EUTRA != nil {
		if err = ie.sftdFrequencyList_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode sftdFrequencyList_EUTRA", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1570_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sftdFrequencyList_NRPresent bool
	if sftdFrequencyList_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sftdFrequencyList_EUTRAPresent bool
	if sftdFrequencyList_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sftdFrequencyList_NRPresent {
		ie.sftdFrequencyList_NR = new(SFTD_FrequencyList_NR)
		if err = ie.sftdFrequencyList_NR.Decode(r); err != nil {
			return utils.WrapError("Decode sftdFrequencyList_NR", err)
		}
	}
	if sftdFrequencyList_EUTRAPresent {
		ie.sftdFrequencyList_EUTRA = new(SFTD_FrequencyList_EUTRA)
		if err = ie.sftdFrequencyList_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode sftdFrequencyList_EUTRA", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1590_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
