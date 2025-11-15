package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBSInterestIndication_r17_IEs struct {
	mbs_FreqList_r17         *CarrierFreqListMBS_r17                         `optional`
	mbs_Priority_r17         *MBSInterestIndication_r17_IEs_mbs_Priority_r17 `optional`
	mbs_ServiceList_r17      *MBS_ServiceList_r17                            `optional`
	lateNonCriticalExtension *[]byte                                         `optional`
	nonCriticalExtension     interface{}                                     `optional`
}

func (ie *MBSInterestIndication_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mbs_FreqList_r17 != nil, ie.mbs_Priority_r17 != nil, ie.mbs_ServiceList_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mbs_FreqList_r17 != nil {
		if err = ie.mbs_FreqList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mbs_FreqList_r17", err)
		}
	}
	if ie.mbs_Priority_r17 != nil {
		if err = ie.mbs_Priority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mbs_Priority_r17", err)
		}
	}
	if ie.mbs_ServiceList_r17 != nil {
		if err = ie.mbs_ServiceList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mbs_ServiceList_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MBSInterestIndication_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	var mbs_FreqList_r17Present bool
	if mbs_FreqList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mbs_Priority_r17Present bool
	if mbs_Priority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mbs_ServiceList_r17Present bool
	if mbs_ServiceList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mbs_FreqList_r17Present {
		ie.mbs_FreqList_r17 = new(CarrierFreqListMBS_r17)
		if err = ie.mbs_FreqList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbs_FreqList_r17", err)
		}
	}
	if mbs_Priority_r17Present {
		ie.mbs_Priority_r17 = new(MBSInterestIndication_r17_IEs_mbs_Priority_r17)
		if err = ie.mbs_Priority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbs_Priority_r17", err)
		}
	}
	if mbs_ServiceList_r17Present {
		ie.mbs_ServiceList_r17 = new(MBS_ServiceList_r17)
		if err = ie.mbs_ServiceList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbs_ServiceList_r17", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
