package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterNR struct {
	frequencyBandListFilter *FreqBandList                       `optional`
	nonCriticalExtension    *UE_CapabilityRequestFilterNR_v1540 `optional`
}

func (ie *UE_CapabilityRequestFilterNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.frequencyBandListFilter != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.frequencyBandListFilter != nil {
		if err = ie.frequencyBandListFilter.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyBandListFilter", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterNR) Decode(r *uper.UperReader) error {
	var err error
	var frequencyBandListFilterPresent bool
	if frequencyBandListFilterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyBandListFilterPresent {
		ie.frequencyBandListFilter = new(FreqBandList)
		if err = ie.frequencyBandListFilter.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyBandListFilter", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_CapabilityRequestFilterNR_v1540)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
