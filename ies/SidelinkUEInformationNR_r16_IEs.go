package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkUEInformationNR_r16_IEs struct {
	sl_RxInterestedFreqList_r16 *SL_InterestedFreqList_r16         `optional`
	sl_TxResourceReqList_r16    *SL_TxResourceReqList_r16          `optional`
	sl_FailureList_r16          *SL_FailureList_r16                `optional`
	lateNonCriticalExtension    *[]byte                            `optional`
	nonCriticalExtension        *SidelinkUEInformationNR_v1700_IEs `optional`
}

func (ie *SidelinkUEInformationNR_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_RxInterestedFreqList_r16 != nil, ie.sl_TxResourceReqList_r16 != nil, ie.sl_FailureList_r16 != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_RxInterestedFreqList_r16 != nil {
		if err = ie.sl_RxInterestedFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RxInterestedFreqList_r16", err)
		}
	}
	if ie.sl_TxResourceReqList_r16 != nil {
		if err = ie.sl_TxResourceReqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxResourceReqList_r16", err)
		}
	}
	if ie.sl_FailureList_r16 != nil {
		if err = ie.sl_FailureList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_FailureList_r16", err)
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

func (ie *SidelinkUEInformationNR_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sl_RxInterestedFreqList_r16Present bool
	if sl_RxInterestedFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxResourceReqList_r16Present bool
	if sl_TxResourceReqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_FailureList_r16Present bool
	if sl_FailureList_r16Present, err = r.ReadBool(); err != nil {
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
	if sl_RxInterestedFreqList_r16Present {
		ie.sl_RxInterestedFreqList_r16 = new(SL_InterestedFreqList_r16)
		if err = ie.sl_RxInterestedFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RxInterestedFreqList_r16", err)
		}
	}
	if sl_TxResourceReqList_r16Present {
		ie.sl_TxResourceReqList_r16 = new(SL_TxResourceReqList_r16)
		if err = ie.sl_TxResourceReqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxResourceReqList_r16", err)
		}
	}
	if sl_FailureList_r16Present {
		ie.sl_FailureList_r16 = new(SL_FailureList_r16)
		if err = ie.sl_FailureList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_FailureList_r16", err)
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
		ie.nonCriticalExtension = new(SidelinkUEInformationNR_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
