package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16 struct {
	slotOffset_r16 *int64 `lb:1,ub:32,optional`
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.slotOffset_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.slotOffset_r16 != nil {
		if err = w.WriteInteger(*ie.slotOffset_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode slotOffset_r16", err)
		}
	}
	return nil
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_aperiodic_r16) Decode(r *uper.UperReader) error {
	var err error
	var slotOffset_r16Present bool
	if slotOffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if slotOffset_r16Present {
		var tmp_int_slotOffset_r16 int64
		if tmp_int_slotOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode slotOffset_r16", err)
		}
		ie.slotOffset_r16 = &tmp_int_slotOffset_r16
	}
	return nil
}
