package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1550 struct {
	dummy *CA_ParametersNR_v1550_dummy `optional`
}

func (ie *CA_ParametersNR_v1550) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dummy != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR_v1550) Decode(r *uper.UperReader) error {
	var err error
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dummyPresent {
		ie.dummy = new(CA_ParametersNR_v1550_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	return nil
}
