package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReferenceTimeInfo_r16 struct {
	time_r16         ReferenceTime_r16                       `madatory`
	uncertainty_r16  *int64                                  `lb:0,ub:32767,optional`
	timeInfoType_r16 *ReferenceTimeInfo_r16_timeInfoType_r16 `optional`
	referenceSFN_r16 *int64                                  `lb:0,ub:1023,optional`
}

func (ie *ReferenceTimeInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uncertainty_r16 != nil, ie.timeInfoType_r16 != nil, ie.referenceSFN_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.time_r16.Encode(w); err != nil {
		return utils.WrapError("Encode time_r16", err)
	}
	if ie.uncertainty_r16 != nil {
		if err = w.WriteInteger(*ie.uncertainty_r16, &uper.Constraint{Lb: 0, Ub: 32767}, false); err != nil {
			return utils.WrapError("Encode uncertainty_r16", err)
		}
	}
	if ie.timeInfoType_r16 != nil {
		if err = ie.timeInfoType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode timeInfoType_r16", err)
		}
	}
	if ie.referenceSFN_r16 != nil {
		if err = w.WriteInteger(*ie.referenceSFN_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode referenceSFN_r16", err)
		}
	}
	return nil
}

func (ie *ReferenceTimeInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var uncertainty_r16Present bool
	if uncertainty_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeInfoType_r16Present bool
	if timeInfoType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var referenceSFN_r16Present bool
	if referenceSFN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.time_r16.Decode(r); err != nil {
		return utils.WrapError("Decode time_r16", err)
	}
	if uncertainty_r16Present {
		var tmp_int_uncertainty_r16 int64
		if tmp_int_uncertainty_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 32767}, false); err != nil {
			return utils.WrapError("Decode uncertainty_r16", err)
		}
		ie.uncertainty_r16 = &tmp_int_uncertainty_r16
	}
	if timeInfoType_r16Present {
		ie.timeInfoType_r16 = new(ReferenceTimeInfo_r16_timeInfoType_r16)
		if err = ie.timeInfoType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode timeInfoType_r16", err)
		}
	}
	if referenceSFN_r16Present {
		var tmp_int_referenceSFN_r16 int64
		if tmp_int_referenceSFN_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode referenceSFN_r16", err)
		}
		ie.referenceSFN_r16 = &tmp_int_referenceSFN_r16
	}
	return nil
}
