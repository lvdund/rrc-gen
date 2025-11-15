package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RedirectedCarrierInfo_EUTRA struct {
	eutraFrequency ARFCN_ValueEUTRA                    `madatory`
	cnType         *RedirectedCarrierInfo_EUTRA_cnType `optional`
}

func (ie *RedirectedCarrierInfo_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cnType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.eutraFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode eutraFrequency", err)
	}
	if ie.cnType != nil {
		if err = ie.cnType.Encode(w); err != nil {
			return utils.WrapError("Encode cnType", err)
		}
	}
	return nil
}

func (ie *RedirectedCarrierInfo_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var cnTypePresent bool
	if cnTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.eutraFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode eutraFrequency", err)
	}
	if cnTypePresent {
		ie.cnType = new(RedirectedCarrierInfo_EUTRA_cnType)
		if err = ie.cnType.Decode(r); err != nil {
			return utils.WrapError("Decode cnType", err)
		}
	}
	return nil
}
