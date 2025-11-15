package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUSCH_AlphaSet struct {
	p0_PUSCH_AlphaSetId P0_PUSCH_AlphaSetId `madatory`
	p0                  *int64              `lb:-16,ub:15,optional`
	alpha               *Alpha              `optional`
}

func (ie *P0_PUSCH_AlphaSet) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.p0 != nil, ie.alpha != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.p0_PUSCH_AlphaSetId.Encode(w); err != nil {
		return utils.WrapError("Encode p0_PUSCH_AlphaSetId", err)
	}
	if ie.p0 != nil {
		if err = w.WriteInteger(*ie.p0, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode p0", err)
		}
	}
	if ie.alpha != nil {
		if err = ie.alpha.Encode(w); err != nil {
			return utils.WrapError("Encode alpha", err)
		}
	}
	return nil
}

func (ie *P0_PUSCH_AlphaSet) Decode(r *uper.UperReader) error {
	var err error
	var p0Present bool
	if p0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var alphaPresent bool
	if alphaPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.p0_PUSCH_AlphaSetId.Decode(r); err != nil {
		return utils.WrapError("Decode p0_PUSCH_AlphaSetId", err)
	}
	if p0Present {
		var tmp_int_p0 int64
		if tmp_int_p0, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode p0", err)
		}
		ie.p0 = &tmp_int_p0
	}
	if alphaPresent {
		ie.alpha = new(Alpha)
		if err = ie.alpha.Decode(r); err != nil {
			return utils.WrapError("Decode alpha", err)
		}
	}
	return nil
}
