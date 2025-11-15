package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v1560 struct {
	receivedFilters                    *[]byte                              `optional`
	measAndMobParametersMRDC_v1560     *MeasAndMobParametersMRDC_v1560      `optional`
	fdd_Add_UE_MRDC_Capabilities_v1560 *UE_MRDC_CapabilityAddXDD_Mode_v1560 `optional`
	tdd_Add_UE_MRDC_Capabilities_v1560 *UE_MRDC_CapabilityAddXDD_Mode_v1560 `optional`
	nonCriticalExtension               *UE_MRDC_Capability_v1610            `optional`
}

func (ie *UE_MRDC_Capability_v1560) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.receivedFilters != nil, ie.measAndMobParametersMRDC_v1560 != nil, ie.fdd_Add_UE_MRDC_Capabilities_v1560 != nil, ie.tdd_Add_UE_MRDC_Capabilities_v1560 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.receivedFilters != nil {
		if err = w.WriteOctetString(*ie.receivedFilters, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode receivedFilters", err)
		}
	}
	if ie.measAndMobParametersMRDC_v1560 != nil {
		if err = ie.measAndMobParametersMRDC_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_v1560", err)
		}
	}
	if ie.fdd_Add_UE_MRDC_Capabilities_v1560 != nil {
		if err = ie.fdd_Add_UE_MRDC_Capabilities_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if ie.tdd_Add_UE_MRDC_Capabilities_v1560 != nil {
		if err = ie.tdd_Add_UE_MRDC_Capabilities_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v1560) Decode(r *uper.UperReader) error {
	var err error
	var receivedFiltersPresent bool
	if receivedFiltersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersMRDC_v1560Present bool
	if measAndMobParametersMRDC_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fdd_Add_UE_MRDC_Capabilities_v1560Present bool
	if fdd_Add_UE_MRDC_Capabilities_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_Add_UE_MRDC_Capabilities_v1560Present bool
	if tdd_Add_UE_MRDC_Capabilities_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if receivedFiltersPresent {
		var tmp_os_receivedFilters []byte
		if tmp_os_receivedFilters, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode receivedFilters", err)
		}
		ie.receivedFilters = &tmp_os_receivedFilters
	}
	if measAndMobParametersMRDC_v1560Present {
		ie.measAndMobParametersMRDC_v1560 = new(MeasAndMobParametersMRDC_v1560)
		if err = ie.measAndMobParametersMRDC_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_v1560", err)
		}
	}
	if fdd_Add_UE_MRDC_Capabilities_v1560Present {
		ie.fdd_Add_UE_MRDC_Capabilities_v1560 = new(UE_MRDC_CapabilityAddXDD_Mode_v1560)
		if err = ie.fdd_Add_UE_MRDC_Capabilities_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if tdd_Add_UE_MRDC_Capabilities_v1560Present {
		ie.tdd_Add_UE_MRDC_Capabilities_v1560 = new(UE_MRDC_CapabilityAddXDD_Mode_v1560)
		if err = ie.tdd_Add_UE_MRDC_Capabilities_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_Add_UE_MRDC_Capabilities_v1560", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_MRDC_Capability_v1610)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
