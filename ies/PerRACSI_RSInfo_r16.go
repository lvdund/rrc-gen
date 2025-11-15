package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRACSI_RSInfo_r16 struct {
	csi_RS_Index_r16                  CSI_RS_Index `madatory`
	numberOfPreamblesSentOnCSI_RS_r16 int64        `lb:1,ub:200,madatory`
}

func (ie *PerRACSI_RSInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.csi_RS_Index_r16.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_Index_r16", err)
	}
	if err = w.WriteInteger(ie.numberOfPreamblesSentOnCSI_RS_r16, &uper.Constraint{Lb: 1, Ub: 200}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfPreamblesSentOnCSI_RS_r16", err)
	}
	return nil
}

func (ie *PerRACSI_RSInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.csi_RS_Index_r16.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS_Index_r16", err)
	}
	var tmp_int_numberOfPreamblesSentOnCSI_RS_r16 int64
	if tmp_int_numberOfPreamblesSentOnCSI_RS_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 200}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfPreamblesSentOnCSI_RS_r16", err)
	}
	ie.numberOfPreamblesSentOnCSI_RS_r16 = tmp_int_numberOfPreamblesSentOnCSI_RS_r16
	return nil
}
