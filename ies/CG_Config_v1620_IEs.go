package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1620_IEs struct {
	ueAssistanceInformationSCG_r16 *[]byte              `optional`
	nonCriticalExtension           *CG_Config_v1630_IEs `optional`
}

func (ie *CG_Config_v1620_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ueAssistanceInformationSCG_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ueAssistanceInformationSCG_r16 != nil {
		if err = w.WriteOctetString(*ie.ueAssistanceInformationSCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ueAssistanceInformationSCG_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1620_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ueAssistanceInformationSCG_r16Present bool
	if ueAssistanceInformationSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ueAssistanceInformationSCG_r16Present {
		var tmp_os_ueAssistanceInformationSCG_r16 []byte
		if tmp_os_ueAssistanceInformationSCG_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ueAssistanceInformationSCG_r16", err)
		}
		ie.ueAssistanceInformationSCG_r16 = &tmp_os_ueAssistanceInformationSCG_r16
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1630_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
