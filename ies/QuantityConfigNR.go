package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfigNR struct {
	quantityConfigCell     QuantityConfigRS  `madatory`
	quantityConfigRS_Index *QuantityConfigRS `optional`
}

func (ie *QuantityConfigNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.quantityConfigRS_Index != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.quantityConfigCell.Encode(w); err != nil {
		return utils.WrapError("Encode quantityConfigCell", err)
	}
	if ie.quantityConfigRS_Index != nil {
		if err = ie.quantityConfigRS_Index.Encode(w); err != nil {
			return utils.WrapError("Encode quantityConfigRS_Index", err)
		}
	}
	return nil
}

func (ie *QuantityConfigNR) Decode(r *uper.UperReader) error {
	var err error
	var quantityConfigRS_IndexPresent bool
	if quantityConfigRS_IndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.quantityConfigCell.Decode(r); err != nil {
		return utils.WrapError("Decode quantityConfigCell", err)
	}
	if quantityConfigRS_IndexPresent {
		ie.quantityConfigRS_Index = new(QuantityConfigRS)
		if err = ie.quantityConfigRS_Index.Decode(r); err != nil {
			return utils.WrapError("Decode quantityConfigRS_Index", err)
		}
	}
	return nil
}
