package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HRNN_r16 struct {
	hrnn_r16 *[]byte `lb:1,ub:maxHRNN_Len_r16,optional`
}

func (ie *HRNN_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.hrnn_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.hrnn_r16 != nil {
		if err = w.WriteOctetString(*ie.hrnn_r16, &uper.Constraint{Lb: 1, Ub: maxHRNN_Len_r16}, false); err != nil {
			return utils.WrapError("Encode hrnn_r16", err)
		}
	}
	return nil
}

func (ie *HRNN_r16) Decode(r *uper.UperReader) error {
	var err error
	var hrnn_r16Present bool
	if hrnn_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if hrnn_r16Present {
		var tmp_os_hrnn_r16 []byte
		if tmp_os_hrnn_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: maxHRNN_Len_r16}, false); err != nil {
			return utils.WrapError("Decode hrnn_r16", err)
		}
		ie.hrnn_r16 = &tmp_os_hrnn_r16
	}
	return nil
}
