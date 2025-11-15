package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MPE_Resource_r17 struct {
	mpe_ResourceId_r17      MPE_ResourceId_r17                       `madatory`
	cell_r17                *ServCellIndex                           `optional`
	additionalPCI_r17       *AdditionalPCIIndex_r17                  `optional`
	mpe_ReferenceSignal_r17 MPE_Resource_r17_mpe_ReferenceSignal_r17 `madatory`
}

func (ie *MPE_Resource_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cell_r17 != nil, ie.additionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.mpe_ResourceId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mpe_ResourceId_r17", err)
	}
	if ie.cell_r17 != nil {
		if err = ie.cell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cell_r17", err)
		}
	}
	if ie.additionalPCI_r17 != nil {
		if err = ie.additionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode additionalPCI_r17", err)
		}
	}
	if err = ie.mpe_ReferenceSignal_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mpe_ReferenceSignal_r17", err)
	}
	return nil
}

func (ie *MPE_Resource_r17) Decode(r *uper.UperReader) error {
	var err error
	var cell_r17Present bool
	if cell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalPCI_r17Present bool
	if additionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.mpe_ResourceId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mpe_ResourceId_r17", err)
	}
	if cell_r17Present {
		ie.cell_r17 = new(ServCellIndex)
		if err = ie.cell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cell_r17", err)
		}
	}
	if additionalPCI_r17Present {
		ie.additionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.additionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode additionalPCI_r17", err)
		}
	}
	if err = ie.mpe_ReferenceSignal_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mpe_ReferenceSignal_r17", err)
	}
	return nil
}
