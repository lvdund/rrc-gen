package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyPathlossReferenceRS_v1710 struct {
	pusch_PathlossReferenceRS_Id_r17 PUSCH_PathlossReferenceRS_Id_r17 `madatory`
	additionalPCI_r17                *AdditionalPCIIndex_r17          `optional`
}

func (ie *DummyPathlossReferenceRS_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pusch_PathlossReferenceRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pusch_PathlossReferenceRS_Id_r17", err)
	}
	if ie.additionalPCI_r17 != nil {
		if err = ie.additionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPCI_r17", err)
		}
	}
	return nil
}

func (ie *DummyPathlossReferenceRS_v1710) Decode(r *uper.UperReader) error {
	var err error
	var additionalPCI_r17Present bool
	if additionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pusch_PathlossReferenceRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pusch_PathlossReferenceRS_Id_r17", err)
	}
	if additionalPCI_r17Present {
		ie.additionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.additionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode additionalPCI_r17", err)
		}
	}
	return nil
}
