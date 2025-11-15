package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PathlossReferenceRS_r17 struct {
	pathlossReferenceRS_Id_r17 PathlossReferenceRS_Id_r17                  `madatory`
	referenceSignal_r17        PathlossReferenceRS_r17_referenceSignal_r17 `madatory`
	additionalPCI_r17          *AdditionalPCIIndex_r17                     `optional`
}

func (ie *PathlossReferenceRS_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pathlossReferenceRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pathlossReferenceRS_Id_r17", err)
	}
	if err = ie.referenceSignal_r17.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal_r17", err)
	}
	if ie.additionalPCI_r17 != nil {
		if err = ie.additionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPCI_r17", err)
		}
	}
	return nil
}

func (ie *PathlossReferenceRS_r17) Decode(r *uper.UperReader) error {
	var err error
	var additionalPCI_r17Present bool
	if additionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pathlossReferenceRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pathlossReferenceRS_Id_r17", err)
	}
	if err = ie.referenceSignal_r17.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal_r17", err)
	}
	if additionalPCI_r17Present {
		ie.additionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.additionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode additionalPCI_r17", err)
		}
	}
	return nil
}
