package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EthernetHeaderCompression_r16_ehc_Downlink_r16 struct {
	drb_ContinueEHC_DL_r16 *EthernetHeaderCompression_r16_ehc_Downlink_r16_drb_ContinueEHC_DL_r16 `optional`
}

func (ie *EthernetHeaderCompression_r16_ehc_Downlink_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drb_ContinueEHC_DL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.drb_ContinueEHC_DL_r16 != nil {
		if err = ie.drb_ContinueEHC_DL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode drb_ContinueEHC_DL_r16", err)
		}
	}
	return nil
}

func (ie *EthernetHeaderCompression_r16_ehc_Downlink_r16) Decode(r *uper.UperReader) error {
	var err error
	var drb_ContinueEHC_DL_r16Present bool
	if drb_ContinueEHC_DL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if drb_ContinueEHC_DL_r16Present {
		ie.drb_ContinueEHC_DL_r16 = new(EthernetHeaderCompression_r16_ehc_Downlink_r16_drb_ContinueEHC_DL_r16)
		if err = ie.drb_ContinueEHC_DL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode drb_ContinueEHC_DL_r16", err)
		}
	}
	return nil
}
