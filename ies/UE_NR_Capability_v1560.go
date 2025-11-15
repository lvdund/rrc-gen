package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1560 struct {
	nrdc_Parameters      *NRDC_Parameters        `optional`
	receivedFilters      *[]byte                 `optional`
	nonCriticalExtension *UE_NR_Capability_v1570 `optional`
}

func (ie *UE_NR_Capability_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrdc_Parameters != nil, ie.receivedFilters != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrdc_Parameters != nil {
		if err = ie.nrdc_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode nrdc_Parameters", err)
		}
	}
	if ie.receivedFilters != nil {
		if err = w.WriteOctetString(*ie.receivedFilters, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode receivedFilters", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1560) Decode(r *uper.UperReader) error {
	var err error
	var nrdc_ParametersPresent bool
	if nrdc_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var receivedFiltersPresent bool
	if receivedFiltersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nrdc_ParametersPresent {
		ie.nrdc_Parameters = new(NRDC_Parameters)
		if err = ie.nrdc_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode nrdc_Parameters", err)
		}
	}
	if receivedFiltersPresent {
		var tmp_os_receivedFilters []byte
		if tmp_os_receivedFilters, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode receivedFilters", err)
		}
		ie.receivedFilters = &tmp_os_receivedFilters
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1570)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
