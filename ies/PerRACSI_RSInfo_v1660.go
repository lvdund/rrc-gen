package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRACSI_RSInfo_v1660 struct {
	csi_RS_Index_v1660 *int64 `lb:1,ub:96,optional`
}

func (ie *PerRACSI_RSInfo_v1660) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.csi_RS_Index_v1660 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.csi_RS_Index_v1660 != nil {
		if err = w.WriteInteger(*ie.csi_RS_Index_v1660, &uper.Constraint{Lb: 1, Ub: 96}, false); err != nil {
			return utils.WrapError("Encode csi_RS_Index_v1660", err)
		}
	}
	return nil
}

func (ie *PerRACSI_RSInfo_v1660) Decode(r *uper.UperReader) error {
	var err error
	var csi_RS_Index_v1660Present bool
	if csi_RS_Index_v1660Present, err = r.ReadBool(); err != nil {
		return err
	}
	if csi_RS_Index_v1660Present {
		var tmp_int_csi_RS_Index_v1660 int64
		if tmp_int_csi_RS_Index_v1660, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode csi_RS_Index_v1660", err)
		}
		ie.csi_RS_Index_v1660 = &tmp_int_csi_RS_Index_v1660
	}
	return nil
}
