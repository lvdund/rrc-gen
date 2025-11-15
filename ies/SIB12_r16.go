package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB12_r16 struct {
	segmentNumber_r16    int64                     `lb:0,ub:63,madatory`
	segmentType_r16      SIB12_r16_segmentType_r16 `madatory`
	segmentContainer_r16 []byte                    `madatory`
}

func (ie *SIB12_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.segmentNumber_r16, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger segmentNumber_r16", err)
	}
	if err = ie.segmentType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode segmentType_r16", err)
	}
	if err = w.WriteOctetString(ie.segmentContainer_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString segmentContainer_r16", err)
	}
	return nil
}

func (ie *SIB12_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_segmentNumber_r16 int64
	if tmp_int_segmentNumber_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger segmentNumber_r16", err)
	}
	ie.segmentNumber_r16 = tmp_int_segmentNumber_r16
	if err = ie.segmentType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode segmentType_r16", err)
	}
	var tmp_os_segmentContainer_r16 []byte
	if tmp_os_segmentContainer_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString segmentContainer_r16", err)
	}
	ie.segmentContainer_r16 = tmp_os_segmentContainer_r16
	return nil
}
