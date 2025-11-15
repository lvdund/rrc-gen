package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeComplete_v1720_IEs struct {
	uplinkTxDirectCurrentMoreCarrierList_r17 *UplinkTxDirectCurrentMoreCarrierList_r17 `optional`
	nonCriticalExtension                     interface{}                               `optional`
}

func (ie *RRCResumeComplete_v1720_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uplinkTxDirectCurrentMoreCarrierList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uplinkTxDirectCurrentMoreCarrierList_r17 != nil {
		if err = ie.uplinkTxDirectCurrentMoreCarrierList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkTxDirectCurrentMoreCarrierList_r17", err)
		}
	}
	return nil
}

func (ie *RRCResumeComplete_v1720_IEs) Decode(r *uper.UperReader) error {
	var err error
	var uplinkTxDirectCurrentMoreCarrierList_r17Present bool
	if uplinkTxDirectCurrentMoreCarrierList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if uplinkTxDirectCurrentMoreCarrierList_r17Present {
		ie.uplinkTxDirectCurrentMoreCarrierList_r17 = new(UplinkTxDirectCurrentMoreCarrierList_r17)
		if err = ie.uplinkTxDirectCurrentMoreCarrierList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkTxDirectCurrentMoreCarrierList_r17", err)
		}
	}
	return nil
}
