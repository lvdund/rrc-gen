package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UERadioAccessCapabilityInformation_IEs struct {
	ue_RadioAccessCapabilityInfo []byte      `madatory`
	nonCriticalExtension         interface{} `optional`
}

func (ie *UERadioAccessCapabilityInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.ue_RadioAccessCapabilityInfo, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString ue_RadioAccessCapabilityInfo", err)
	}
	return nil
}

func (ie *UERadioAccessCapabilityInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var tmp_os_ue_RadioAccessCapabilityInfo []byte
	if tmp_os_ue_RadioAccessCapabilityInfo, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString ue_RadioAccessCapabilityInfo", err)
	}
	ie.ue_RadioAccessCapabilityInfo = tmp_os_ue_RadioAccessCapabilityInfo
	return nil
}
