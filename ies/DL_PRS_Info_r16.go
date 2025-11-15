package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DL_PRS_Info_r16 struct {
	dl_PRS_ID_r16            int64  `lb:0,ub:255,madatory`
	dl_PRS_ResourceSetId_r16 int64  `lb:0,ub:7,madatory`
	dl_PRS_ResourceId_r16    *int64 `lb:0,ub:63,optional`
}

func (ie *DL_PRS_Info_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_PRS_ResourceId_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.dl_PRS_ID_r16, &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("WriteInteger dl_PRS_ID_r16", err)
	}
	if err = w.WriteInteger(ie.dl_PRS_ResourceSetId_r16, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger dl_PRS_ResourceSetId_r16", err)
	}
	if ie.dl_PRS_ResourceId_r16 != nil {
		if err = w.WriteInteger(*ie.dl_PRS_ResourceId_r16, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode dl_PRS_ResourceId_r16", err)
		}
	}
	return nil
}

func (ie *DL_PRS_Info_r16) Decode(r *uper.UperReader) error {
	var err error
	var dl_PRS_ResourceId_r16Present bool
	if dl_PRS_ResourceId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_dl_PRS_ID_r16 int64
	if tmp_int_dl_PRS_ID_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("ReadInteger dl_PRS_ID_r16", err)
	}
	ie.dl_PRS_ID_r16 = tmp_int_dl_PRS_ID_r16
	var tmp_int_dl_PRS_ResourceSetId_r16 int64
	if tmp_int_dl_PRS_ResourceSetId_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger dl_PRS_ResourceSetId_r16", err)
	}
	ie.dl_PRS_ResourceSetId_r16 = tmp_int_dl_PRS_ResourceSetId_r16
	if dl_PRS_ResourceId_r16Present {
		var tmp_int_dl_PRS_ResourceId_r16 int64
		if tmp_int_dl_PRS_ResourceId_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode dl_PRS_ResourceId_r16", err)
		}
		ie.dl_PRS_ResourceId_r16 = &tmp_int_dl_PRS_ResourceId_r16
	}
	return nil
}
