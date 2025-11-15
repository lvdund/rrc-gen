package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EthernetHeaderCompression_r16_ehc_Common_r16 struct {
	ehc_CID_Length_r16 EthernetHeaderCompression_r16_ehc_Common_r16_ehc_CID_Length_r16 `madatory`
}

func (ie *EthernetHeaderCompression_r16_ehc_Common_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ehc_CID_Length_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ehc_CID_Length_r16", err)
	}
	return nil
}

func (ie *EthernetHeaderCompression_r16_ehc_Common_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ehc_CID_Length_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ehc_CID_Length_r16", err)
	}
	return nil
}
