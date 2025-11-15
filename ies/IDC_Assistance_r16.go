package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IDC_Assistance_r16 struct {
	affectedCarrierFreqList_r16     *AffectedCarrierFreqList_r16     `optional`
	affectedCarrierFreqCombList_r16 *AffectedCarrierFreqCombList_r16 `optional`
}

func (ie *IDC_Assistance_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.affectedCarrierFreqList_r16 != nil, ie.affectedCarrierFreqCombList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.affectedCarrierFreqList_r16 != nil {
		if err = ie.affectedCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode affectedCarrierFreqList_r16", err)
		}
	}
	if ie.affectedCarrierFreqCombList_r16 != nil {
		if err = ie.affectedCarrierFreqCombList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode affectedCarrierFreqCombList_r16", err)
		}
	}
	return nil
}

func (ie *IDC_Assistance_r16) Decode(r *uper.UperReader) error {
	var err error
	var affectedCarrierFreqList_r16Present bool
	if affectedCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var affectedCarrierFreqCombList_r16Present bool
	if affectedCarrierFreqCombList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if affectedCarrierFreqList_r16Present {
		ie.affectedCarrierFreqList_r16 = new(AffectedCarrierFreqList_r16)
		if err = ie.affectedCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode affectedCarrierFreqList_r16", err)
		}
	}
	if affectedCarrierFreqCombList_r16Present {
		ie.affectedCarrierFreqCombList_r16 = new(AffectedCarrierFreqCombList_r16)
		if err = ie.affectedCarrierFreqCombList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode affectedCarrierFreqCombList_r16", err)
		}
	}
	return nil
}
