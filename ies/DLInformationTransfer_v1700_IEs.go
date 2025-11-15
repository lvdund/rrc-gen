package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransfer_v1700_IEs struct {
	dedicatedInfoF1c_r17 *DedicatedInfoF1c_r17                             `optional`
	rxTxTimeDiff_gNB_r17 *RxTxTimeDiff_r17                                 `optional`
	ta_PDC_r17           *DLInformationTransfer_v1700_IEs_ta_PDC_r17       `optional`
	sib9Fallback_r17     *DLInformationTransfer_v1700_IEs_sib9Fallback_r17 `optional`
	nonCriticalExtension interface{}                                       `optional`
}

func (ie *DLInformationTransfer_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dedicatedInfoF1c_r17 != nil, ie.rxTxTimeDiff_gNB_r17 != nil, ie.ta_PDC_r17 != nil, ie.sib9Fallback_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dedicatedInfoF1c_r17 != nil {
		if err = ie.dedicatedInfoF1c_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dedicatedInfoF1c_r17", err)
		}
	}
	if ie.rxTxTimeDiff_gNB_r17 != nil {
		if err = ie.rxTxTimeDiff_gNB_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rxTxTimeDiff_gNB_r17", err)
		}
	}
	if ie.ta_PDC_r17 != nil {
		if err = ie.ta_PDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ta_PDC_r17", err)
		}
	}
	if ie.sib9Fallback_r17 != nil {
		if err = ie.sib9Fallback_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sib9Fallback_r17", err)
		}
	}
	return nil
}

func (ie *DLInformationTransfer_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var dedicatedInfoF1c_r17Present bool
	if dedicatedInfoF1c_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rxTxTimeDiff_gNB_r17Present bool
	if rxTxTimeDiff_gNB_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ta_PDC_r17Present bool
	if ta_PDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sib9Fallback_r17Present bool
	if sib9Fallback_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dedicatedInfoF1c_r17Present {
		ie.dedicatedInfoF1c_r17 = new(DedicatedInfoF1c_r17)
		if err = ie.dedicatedInfoF1c_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dedicatedInfoF1c_r17", err)
		}
	}
	if rxTxTimeDiff_gNB_r17Present {
		ie.rxTxTimeDiff_gNB_r17 = new(RxTxTimeDiff_r17)
		if err = ie.rxTxTimeDiff_gNB_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rxTxTimeDiff_gNB_r17", err)
		}
	}
	if ta_PDC_r17Present {
		ie.ta_PDC_r17 = new(DLInformationTransfer_v1700_IEs_ta_PDC_r17)
		if err = ie.ta_PDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ta_PDC_r17", err)
		}
	}
	if sib9Fallback_r17Present {
		ie.sib9Fallback_r17 = new(DLInformationTransfer_v1700_IEs_sib9Fallback_r17)
		if err = ie.sib9Fallback_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sib9Fallback_r17", err)
		}
	}
	return nil
}
