package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CGI_InfoNR_noSIB1 struct {
	ssb_SubcarrierOffset int64            `lb:0,ub:15,madatory`
	pdcch_ConfigSIB1     PDCCH_ConfigSIB1 `madatory`
}

func (ie *CGI_InfoNR_noSIB1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.ssb_SubcarrierOffset, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger ssb_SubcarrierOffset", err)
	}
	if err = ie.pdcch_ConfigSIB1.Encode(w); err != nil {
		return utils.WrapError("Encode pdcch_ConfigSIB1", err)
	}
	return nil
}

func (ie *CGI_InfoNR_noSIB1) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_ssb_SubcarrierOffset int64
	if tmp_int_ssb_SubcarrierOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger ssb_SubcarrierOffset", err)
	}
	ie.ssb_SubcarrierOffset = tmp_int_ssb_SubcarrierOffset
	if err = ie.pdcch_ConfigSIB1.Decode(r); err != nil {
		return utils.WrapError("Decode pdcch_ConfigSIB1", err)
	}
	return nil
}
