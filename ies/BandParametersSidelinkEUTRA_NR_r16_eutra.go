package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelinkEUTRA_NR_r16_eutra struct {
	bandParametersSidelinkEUTRA1_r16 *[]byte `optional`
	bandParametersSidelinkEUTRA2_r16 *[]byte `optional`
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_eutra) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bandParametersSidelinkEUTRA1_r16 != nil, ie.bandParametersSidelinkEUTRA2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bandParametersSidelinkEUTRA1_r16 != nil {
		if err = w.WriteOctetString(*ie.bandParametersSidelinkEUTRA1_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode bandParametersSidelinkEUTRA1_r16", err)
		}
	}
	if ie.bandParametersSidelinkEUTRA2_r16 != nil {
		if err = w.WriteOctetString(*ie.bandParametersSidelinkEUTRA2_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode bandParametersSidelinkEUTRA2_r16", err)
		}
	}
	return nil
}

func (ie *BandParametersSidelinkEUTRA_NR_r16_eutra) Decode(r *uper.UperReader) error {
	var err error
	var bandParametersSidelinkEUTRA1_r16Present bool
	if bandParametersSidelinkEUTRA1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bandParametersSidelinkEUTRA2_r16Present bool
	if bandParametersSidelinkEUTRA2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bandParametersSidelinkEUTRA1_r16Present {
		var tmp_os_bandParametersSidelinkEUTRA1_r16 []byte
		if tmp_os_bandParametersSidelinkEUTRA1_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode bandParametersSidelinkEUTRA1_r16", err)
		}
		ie.bandParametersSidelinkEUTRA1_r16 = &tmp_os_bandParametersSidelinkEUTRA1_r16
	}
	if bandParametersSidelinkEUTRA2_r16Present {
		var tmp_os_bandParametersSidelinkEUTRA2_r16 []byte
		if tmp_os_bandParametersSidelinkEUTRA2_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode bandParametersSidelinkEUTRA2_r16", err)
		}
		ie.bandParametersSidelinkEUTRA2_r16 = &tmp_os_bandParametersSidelinkEUTRA2_r16
	}
	return nil
}
