package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MobilityFromNRCommand_IEs struct {
	targetRAT_Type             MobilityFromNRCommand_IEs_targetRAT_Type `madatory`
	targetRAT_MessageContainer []byte                                   `madatory`
	nas_SecurityParamFromNR    *[]byte                                  `optional`
	lateNonCriticalExtension   *[]byte                                  `optional`
	nonCriticalExtension       *MobilityFromNRCommand_v1610_IEs         `optional`
}

func (ie *MobilityFromNRCommand_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nas_SecurityParamFromNR != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.targetRAT_Type.Encode(w); err != nil {
		return utils.WrapError("Encode targetRAT_Type", err)
	}
	if err = w.WriteOctetString(ie.targetRAT_MessageContainer, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString targetRAT_MessageContainer", err)
	}
	if ie.nas_SecurityParamFromNR != nil {
		if err = w.WriteOctetString(*ie.nas_SecurityParamFromNR, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode nas_SecurityParamFromNR", err)
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

func (ie *MobilityFromNRCommand_IEs) Decode(r *uper.UperReader) error {
	var err error
	var nas_SecurityParamFromNRPresent bool
	if nas_SecurityParamFromNRPresent, err = r.ReadBool(); err != nil {
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
	if err = ie.targetRAT_Type.Decode(r); err != nil {
		return utils.WrapError("Decode targetRAT_Type", err)
	}
	var tmp_os_targetRAT_MessageContainer []byte
	if tmp_os_targetRAT_MessageContainer, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString targetRAT_MessageContainer", err)
	}
	ie.targetRAT_MessageContainer = tmp_os_targetRAT_MessageContainer
	if nas_SecurityParamFromNRPresent {
		var tmp_os_nas_SecurityParamFromNR []byte
		if tmp_os_nas_SecurityParamFromNR, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode nas_SecurityParamFromNR", err)
		}
		ie.nas_SecurityParamFromNR = &tmp_os_nas_SecurityParamFromNR
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(MobilityFromNRCommand_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
