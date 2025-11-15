package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeComplete_v1640_IEs struct {
	uplinkTxDirectCurrentTwoCarrierList_r16 *UplinkTxDirectCurrentTwoCarrierList_r16 `optional`
	nonCriticalExtension                    *RRCResumeComplete_v1700_IEs             `optional`
}

func (ie *RRCResumeComplete_v1640_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uplinkTxDirectCurrentTwoCarrierList_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uplinkTxDirectCurrentTwoCarrierList_r16 != nil {
		if err = ie.uplinkTxDirectCurrentTwoCarrierList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxDirectCurrentTwoCarrierList_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCResumeComplete_v1640_IEs) Decode(r *uper.UperReader) error {
	var err error
	var uplinkTxDirectCurrentTwoCarrierList_r16Present bool
	if uplinkTxDirectCurrentTwoCarrierList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if uplinkTxDirectCurrentTwoCarrierList_r16Present {
		ie.uplinkTxDirectCurrentTwoCarrierList_r16 = new(UplinkTxDirectCurrentTwoCarrierList_r16)
		if err = ie.uplinkTxDirectCurrentTwoCarrierList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxDirectCurrentTwoCarrierList_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCResumeComplete_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
