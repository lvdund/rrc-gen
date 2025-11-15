package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeComplete_IEs struct {
	dedicatedNAS_Message      *DedicatedNAS_Message        `optional`
	selectedPLMN_Identity     *int64                       `lb:1,ub:maxPLMN,optional`
	uplinkTxDirectCurrentList *UplinkTxDirectCurrentList   `optional`
	lateNonCriticalExtension  *[]byte                      `optional`
	nonCriticalExtension      *RRCResumeComplete_v1610_IEs `optional`
}

func (ie *RRCResumeComplete_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dedicatedNAS_Message != nil, ie.selectedPLMN_Identity != nil, ie.uplinkTxDirectCurrentList != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dedicatedNAS_Message != nil {
		if err = ie.dedicatedNAS_Message.Encode(w); err != nil {
			return utils.WrapError("Encode dedicatedNAS_Message", err)
		}
	}
	if ie.selectedPLMN_Identity != nil {
		if err = w.WriteInteger(*ie.selectedPLMN_Identity, &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			return utils.WrapError("Encode selectedPLMN_Identity", err)
		}
	}
	if ie.uplinkTxDirectCurrentList != nil {
		if err = ie.uplinkTxDirectCurrentList.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxDirectCurrentList", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCResumeComplete_IEs) Decode(r *uper.UperReader) error {
	var err error
	var dedicatedNAS_MessagePresent bool
	if dedicatedNAS_MessagePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var selectedPLMN_IdentityPresent bool
	if selectedPLMN_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkTxDirectCurrentListPresent bool
	if uplinkTxDirectCurrentListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dedicatedNAS_MessagePresent {
		ie.dedicatedNAS_Message = new(DedicatedNAS_Message)
		if err = ie.dedicatedNAS_Message.Decode(r); err != nil {
			return utils.WrapError("Decode dedicatedNAS_Message", err)
		}
	}
	if selectedPLMN_IdentityPresent {
		var tmp_int_selectedPLMN_Identity int64
		if tmp_int_selectedPLMN_Identity, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			return utils.WrapError("Decode selectedPLMN_Identity", err)
		}
		ie.selectedPLMN_Identity = &tmp_int_selectedPLMN_Identity
	}
	if uplinkTxDirectCurrentListPresent {
		ie.uplinkTxDirectCurrentList = new(UplinkTxDirectCurrentList)
		if err = ie.uplinkTxDirectCurrentList.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxDirectCurrentList", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCResumeComplete_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
