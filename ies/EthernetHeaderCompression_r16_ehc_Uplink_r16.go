package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EthernetHeaderCompression_r16_ehc_Uplink_r16 struct {
	maxCID_EHC_UL_r16      int64                                                                `lb:1,ub:32767,madatory`
	drb_ContinueEHC_UL_r16 *EthernetHeaderCompression_r16_ehc_Uplink_r16_drb_ContinueEHC_UL_r16 `optional`
}

func (ie *EthernetHeaderCompression_r16_ehc_Uplink_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drb_ContinueEHC_UL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.maxCID_EHC_UL_r16, &uper.Constraint{Lb: 1, Ub: 32767}, false); err != nil {
		return utils.WrapError("WriteInteger maxCID_EHC_UL_r16", err)
	}
	if ie.drb_ContinueEHC_UL_r16 != nil {
		if err = ie.drb_ContinueEHC_UL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode drb_ContinueEHC_UL_r16", err)
		}
	}
	return nil
}

func (ie *EthernetHeaderCompression_r16_ehc_Uplink_r16) Decode(r *uper.UperReader) error {
	var err error
	var drb_ContinueEHC_UL_r16Present bool
	if drb_ContinueEHC_UL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_maxCID_EHC_UL_r16 int64
	if tmp_int_maxCID_EHC_UL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32767}, false); err != nil {
		return utils.WrapError("ReadInteger maxCID_EHC_UL_r16", err)
	}
	ie.maxCID_EHC_UL_r16 = tmp_int_maxCID_EHC_UL_r16
	if drb_ContinueEHC_UL_r16Present {
		ie.drb_ContinueEHC_UL_r16 = new(EthernetHeaderCompression_r16_ehc_Uplink_r16_drb_ContinueEHC_UL_r16)
		if err = ie.drb_ContinueEHC_UL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode drb_ContinueEHC_UL_r16", err)
		}
	}
	return nil
}
