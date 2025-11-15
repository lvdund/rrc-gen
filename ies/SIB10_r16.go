package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB10_r16 struct {
	hrnn_List_r16            *HRNN_List_r16 `optional`
	lateNonCriticalExtension *[]byte        `optional`
}

func (ie *SIB10_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.hrnn_List_r16 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.hrnn_List_r16 != nil {
		if err = ie.hrnn_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode hrnn_List_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB10_r16) Decode(r *uper.UperReader) error {
	var err error
	var hrnn_List_r16Present bool
	if hrnn_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if hrnn_List_r16Present {
		ie.hrnn_List_r16 = new(HRNN_List_r16)
		if err = ie.hrnn_List_r16.Decode(r); err != nil {
			return utils.WrapError("Decode hrnn_List_r16", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
