package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v1610 struct {
	measAndMobParametersMRDC_v1610 *MeasAndMobParametersMRDC_v1610 `optional`
	generalParametersMRDC_v1610    *GeneralParametersMRDC_v1610    `optional`
	pdcp_ParametersMRDC_v1610      *PDCP_ParametersMRDC_v1610      `optional`
	nonCriticalExtension           *UE_MRDC_Capability_v1700       `optional`
}

func (ie *UE_MRDC_Capability_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC_v1610 != nil, ie.generalParametersMRDC_v1610 != nil, ie.pdcp_ParametersMRDC_v1610 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC_v1610 != nil {
		if err = ie.measAndMobParametersMRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC_v1610", err)
		}
	}
	if ie.generalParametersMRDC_v1610 != nil {
		if err = ie.generalParametersMRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode generalParametersMRDC_v1610", err)
		}
	}
	if ie.pdcp_ParametersMRDC_v1610 != nil {
		if err = ie.pdcp_ParametersMRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_ParametersMRDC_v1610", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v1610) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDC_v1610Present bool
	if measAndMobParametersMRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var generalParametersMRDC_v1610Present bool
	if generalParametersMRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_ParametersMRDC_v1610Present bool
	if pdcp_ParametersMRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersMRDC_v1610Present {
		ie.measAndMobParametersMRDC_v1610 = new(MeasAndMobParametersMRDC_v1610)
		if err = ie.measAndMobParametersMRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC_v1610", err)
		}
	}
	if generalParametersMRDC_v1610Present {
		ie.generalParametersMRDC_v1610 = new(GeneralParametersMRDC_v1610)
		if err = ie.generalParametersMRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode generalParametersMRDC_v1610", err)
		}
	}
	if pdcp_ParametersMRDC_v1610Present {
		ie.pdcp_ParametersMRDC_v1610 = new(PDCP_ParametersMRDC_v1610)
		if err = ie.pdcp_ParametersMRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_ParametersMRDC_v1610", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_MRDC_Capability_v1700)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
