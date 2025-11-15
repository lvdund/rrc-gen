package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB17_r17 struct {
	segmentNumber_r17    int64                     `lb:0,ub:63,madatory`
	segmentType_r17      SIB17_r17_segmentType_r17 `madatory`
	segmentContainer_r17 []byte                    `madatory`
}

func (ie *SIB17_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.segmentNumber_r17, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger segmentNumber_r17", err)
	}
	if err = ie.segmentType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode segmentType_r17", err)
	}
	if err = w.WriteOctetString(ie.segmentContainer_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString segmentContainer_r17", err)
	}
	return nil
}

func (ie *SIB17_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_segmentNumber_r17 int64
	if tmp_int_segmentNumber_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger segmentNumber_r17", err)
	}
	ie.segmentNumber_r17 = tmp_int_segmentNumber_r17
	if err = ie.segmentType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode segmentType_r17", err)
	}
	var tmp_os_segmentContainer_r17 []byte
	if tmp_os_segmentContainer_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString segmentContainer_r17", err)
	}
	ie.segmentContainer_r17 = tmp_os_segmentContainer_r17
	return nil
}
