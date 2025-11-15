package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEPositioningAssistanceInfo_r17_IEs struct {
	ue_TxTEG_AssociationList_r17 *UE_TxTEG_AssociationList_r17          `optional`
	lateNonCriticalExtension     *[]byte                                `optional`
	nonCriticalExtension         *UEPositioningAssistanceInfo_v1720_IEs `optional`
}

func (ie *UEPositioningAssistanceInfo_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ue_TxTEG_AssociationList_r17 != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ue_TxTEG_AssociationList_r17 != nil {
		if err = ie.ue_TxTEG_AssociationList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ue_TxTEG_AssociationList_r17", err)
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

func (ie *UEPositioningAssistanceInfo_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ue_TxTEG_AssociationList_r17Present bool
	if ue_TxTEG_AssociationList_r17Present, err = r.ReadBool(); err != nil {
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
	if ue_TxTEG_AssociationList_r17Present {
		ie.ue_TxTEG_AssociationList_r17 = new(UE_TxTEG_AssociationList_r17)
		if err = ie.ue_TxTEG_AssociationList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ue_TxTEG_AssociationList_r17", err)
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
		ie.nonCriticalExtension = new(UEPositioningAssistanceInfo_v1720_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
