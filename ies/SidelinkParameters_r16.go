package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkParameters_r16 struct {
	sidelinkParametersNR_r16    *SidelinkParametersNR_r16    `optional`
	sidelinkParametersEUTRA_r16 *SidelinkParametersEUTRA_r16 `optional`
}

func (ie *SidelinkParameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sidelinkParametersNR_r16 != nil, ie.sidelinkParametersEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sidelinkParametersNR_r16 != nil {
		if err = ie.sidelinkParametersNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sidelinkParametersNR_r16", err)
		}
	}
	if ie.sidelinkParametersEUTRA_r16 != nil {
		if err = ie.sidelinkParametersEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sidelinkParametersEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *SidelinkParameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var sidelinkParametersNR_r16Present bool
	if sidelinkParametersNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sidelinkParametersEUTRA_r16Present bool
	if sidelinkParametersEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sidelinkParametersNR_r16Present {
		ie.sidelinkParametersNR_r16 = new(SidelinkParametersNR_r16)
		if err = ie.sidelinkParametersNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sidelinkParametersNR_r16", err)
		}
	}
	if sidelinkParametersEUTRA_r16Present {
		ie.sidelinkParametersEUTRA_r16 = new(SidelinkParametersEUTRA_r16)
		if err = ie.sidelinkParametersEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sidelinkParametersEUTRA_r16", err)
		}
	}
	return nil
}
