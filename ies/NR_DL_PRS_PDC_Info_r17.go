package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_DL_PRS_PDC_Info_r17 struct {
	nr_DL_PRS_PDC_ResourceSet_r17 *NR_DL_PRS_PDC_ResourceSet_r17 `optional`
}

func (ie *NR_DL_PRS_PDC_Info_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nr_DL_PRS_PDC_ResourceSet_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nr_DL_PRS_PDC_ResourceSet_r17 != nil {
		if err = ie.nr_DL_PRS_PDC_ResourceSet_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nr_DL_PRS_PDC_ResourceSet_r17", err)
		}
	}
	return nil
}

func (ie *NR_DL_PRS_PDC_Info_r17) Decode(r *uper.UperReader) error {
	var err error
	var nr_DL_PRS_PDC_ResourceSet_r17Present bool
	if nr_DL_PRS_PDC_ResourceSet_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if nr_DL_PRS_PDC_ResourceSet_r17Present {
		ie.nr_DL_PRS_PDC_ResourceSet_r17 = new(NR_DL_PRS_PDC_ResourceSet_r17)
		if err = ie.nr_DL_PRS_PDC_ResourceSet_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nr_DL_PRS_PDC_ResourceSet_r17", err)
		}
	}
	return nil
}
