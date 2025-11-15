package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1690 struct {
	ca_ParametersNR_v1690 *CA_ParametersNR_v1690 `optional`
}

func (ie *BandCombination_v1690) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ca_ParametersNR_v1690 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ca_ParametersNR_v1690 != nil {
		if err = ie.ca_ParametersNR_v1690.Encode(w); err != nil {
			return utils.WrapError("Encode ca_ParametersNR_v1690", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1690) Decode(r *uper.UperReader) error {
	var err error
	var ca_ParametersNR_v1690Present bool
	if ca_ParametersNR_v1690Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ca_ParametersNR_v1690Present {
		ie.ca_ParametersNR_v1690 = new(CA_ParametersNR_v1690)
		if err = ie.ca_ParametersNR_v1690.Decode(r); err != nil {
			return utils.WrapError("Decode ca_ParametersNR_v1690", err)
		}
	}
	return nil
}
