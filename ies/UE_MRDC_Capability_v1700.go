package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v1700 struct {
	measAndMobParametersMRDC_v1700 MeasAndMobParametersMRDC_v1700 `madatory`
	nonCriticalExtension           *UE_MRDC_Capability_v1730      `optional`
}

func (ie *UE_MRDC_Capability_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measAndMobParametersMRDC_v1700.Encode(w); err != nil {
		return utils.WrapError("Encode measAndMobParametersMRDC_v1700", err)
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v1700) Decode(r *uper.UperReader) error {
	var err error
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measAndMobParametersMRDC_v1700.Decode(r); err != nil {
		return utils.WrapError("Decode measAndMobParametersMRDC_v1700", err)
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_MRDC_Capability_v1730)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
