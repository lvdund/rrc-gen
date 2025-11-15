package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NRDC_Parameters struct {
	measAndMobParametersNRDC     *MeasAndMobParametersMRDC       `optional`
	generalParametersNRDC        *GeneralParametersMRDC_XDD_Diff `optional`
	fdd_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	tdd_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	fr1_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	fr2_Add_UE_NRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	dummy2                       *[]byte                         `optional`
	dummy                        interface{}                     `optional`
}

func (ie *NRDC_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersNRDC != nil, ie.generalParametersNRDC != nil, ie.fdd_Add_UE_NRDC_Capabilities != nil, ie.tdd_Add_UE_NRDC_Capabilities != nil, ie.fr1_Add_UE_NRDC_Capabilities != nil, ie.fr2_Add_UE_NRDC_Capabilities != nil, ie.dummy2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersNRDC != nil {
		if err = ie.measAndMobParametersNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersNRDC", err)
		}
	}
	if ie.generalParametersNRDC != nil {
		if err = ie.generalParametersNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode generalParametersNRDC", err)
		}
	}
	if ie.fdd_Add_UE_NRDC_Capabilities != nil {
		if err = ie.fdd_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.tdd_Add_UE_NRDC_Capabilities != nil {
		if err = ie.tdd_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.fr1_Add_UE_NRDC_Capabilities != nil {
		if err = ie.fr1_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.fr2_Add_UE_NRDC_Capabilities != nil {
		if err = ie.fr2_Add_UE_NRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fr2_Add_UE_NRDC_Capabilities", err)
		}
	}
	if ie.dummy2 != nil {
		if err = w.WriteOctetString(*ie.dummy2, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode dummy2", err)
		}
	}
	return nil
}

func (ie *NRDC_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersNRDCPresent bool
	if measAndMobParametersNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var generalParametersNRDCPresent bool
	if generalParametersNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fdd_Add_UE_NRDC_CapabilitiesPresent bool
	if fdd_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_Add_UE_NRDC_CapabilitiesPresent bool
	if tdd_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_Add_UE_NRDC_CapabilitiesPresent bool
	if fr1_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_Add_UE_NRDC_CapabilitiesPresent bool
	if fr2_Add_UE_NRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy2Present bool
	if dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measAndMobParametersNRDCPresent {
		ie.measAndMobParametersNRDC = new(MeasAndMobParametersMRDC)
		if err = ie.measAndMobParametersNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersNRDC", err)
		}
	}
	if generalParametersNRDCPresent {
		ie.generalParametersNRDC = new(GeneralParametersMRDC_XDD_Diff)
		if err = ie.generalParametersNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode generalParametersNRDC", err)
		}
	}
	if fdd_Add_UE_NRDC_CapabilitiesPresent {
		ie.fdd_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.fdd_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if tdd_Add_UE_NRDC_CapabilitiesPresent {
		ie.tdd_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.tdd_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_Add_UE_NRDC_Capabilities", err)
		}
	}
	if fr1_Add_UE_NRDC_CapabilitiesPresent {
		ie.fr1_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.fr1_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_Add_UE_NRDC_Capabilities", err)
		}
	}
	if fr2_Add_UE_NRDC_CapabilitiesPresent {
		ie.fr2_Add_UE_NRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.fr2_Add_UE_NRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_Add_UE_NRDC_Capabilities", err)
		}
	}
	if dummy2Present {
		var tmp_os_dummy2 []byte
		if tmp_os_dummy2, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode dummy2", err)
		}
		ie.dummy2 = &tmp_os_dummy2
	}
	return nil
}
