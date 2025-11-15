package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1620_IEs struct {
	ueAssistanceInformationSourceSCG_r16 *[]byte                  `optional`
	nonCriticalExtension                 *CG_ConfigInfo_v1640_IEs `optional`
}

func (ie *CG_ConfigInfo_v1620_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ueAssistanceInformationSourceSCG_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ueAssistanceInformationSourceSCG_r16 != nil {
		if err = w.WriteOctetString(*ie.ueAssistanceInformationSourceSCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ueAssistanceInformationSourceSCG_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1620_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ueAssistanceInformationSourceSCG_r16Present bool
	if ueAssistanceInformationSourceSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ueAssistanceInformationSourceSCG_r16Present {
		var tmp_os_ueAssistanceInformationSourceSCG_r16 []byte
		if tmp_os_ueAssistanceInformationSourceSCG_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ueAssistanceInformationSourceSCG_r16", err)
		}
		ie.ueAssistanceInformationSourceSCG_r16 = &tmp_os_ueAssistanceInformationSourceSCG_r16
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1640_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
