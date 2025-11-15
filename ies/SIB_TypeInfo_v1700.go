package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB_TypeInfo_v1700 struct {
	sibType_r17   *SIB_TypeInfo_v1700_sibType_r17   `optional`
	valueTag_r17  *int64                            `lb:0,ub:31,optional`
	areaScope_r17 *SIB_TypeInfo_v1700_areaScope_r17 `optional`
}

func (ie *SIB_TypeInfo_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sibType_r17 != nil, ie.valueTag_r17 != nil, ie.areaScope_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sibType_r17 != nil {
		if err = ie.sibType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sibType_r17", err)
		}
	}
	if ie.valueTag_r17 != nil {
		if err = w.WriteInteger(*ie.valueTag_r17, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode valueTag_r17", err)
		}
	}
	if ie.areaScope_r17 != nil {
		if err = ie.areaScope_r17.Encode(w); err != nil {
			return utils.WrapError("Encode areaScope_r17", err)
		}
	}
	return nil
}

func (ie *SIB_TypeInfo_v1700) Decode(r *uper.UperReader) error {
	var err error
	var sibType_r17Present bool
	if sibType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var valueTag_r17Present bool
	if valueTag_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var areaScope_r17Present bool
	if areaScope_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sibType_r17Present {
		ie.sibType_r17 = new(SIB_TypeInfo_v1700_sibType_r17)
		if err = ie.sibType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sibType_r17", err)
		}
	}
	if valueTag_r17Present {
		var tmp_int_valueTag_r17 int64
		if tmp_int_valueTag_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode valueTag_r17", err)
		}
		ie.valueTag_r17 = &tmp_int_valueTag_r17
	}
	if areaScope_r17Present {
		ie.areaScope_r17 = new(SIB_TypeInfo_v1700_areaScope_r17)
		if err = ie.areaScope_r17.Decode(r); err != nil {
			return utils.WrapError("Decode areaScope_r17", err)
		}
	}
	return nil
}
