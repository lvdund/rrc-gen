package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB_TypeInfo struct {
	type_sib  SIB_TypeInfo_type_sib   `madatory`
	valueTag  *int64                  `lb:0,ub:31,optional`
	areaScope *SIB_TypeInfo_areaScope `optional`
}

func (ie *SIB_TypeInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.valueTag != nil, ie.areaScope != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.type_sib.Encode(w); err != nil {
		return utils.WrapError("Encode type_sib", err)
	}
	if ie.valueTag != nil {
		if err = w.WriteInteger(*ie.valueTag, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode valueTag", err)
		}
	}
	if ie.areaScope != nil {
		if err = ie.areaScope.Encode(w); err != nil {
			return utils.WrapError("Encode areaScope", err)
		}
	}
	return nil
}

func (ie *SIB_TypeInfo) Decode(r *uper.UperReader) error {
	var err error
	var valueTagPresent bool
	if valueTagPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var areaScopePresent bool
	if areaScopePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.type_sib.Decode(r); err != nil {
		return utils.WrapError("Decode type_sib", err)
	}
	if valueTagPresent {
		var tmp_int_valueTag int64
		if tmp_int_valueTag, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode valueTag", err)
		}
		ie.valueTag = &tmp_int_valueTag
	}
	if areaScopePresent {
		ie.areaScope = new(SIB_TypeInfo_areaScope)
		if err = ie.areaScope.Decode(r); err != nil {
			return utils.WrapError("Decode areaScope", err)
		}
	}
	return nil
}
