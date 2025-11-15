package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_v1630 struct {
	dl_ReceptionIntraCellGuardband_r16 *SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionIntraCellGuardband_r16 `optional`
	dl_ReceptionLBT_subsetRB_r16       *SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionLBT_subsetRB_r16       `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1630) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_ReceptionIntraCellGuardband_r16 != nil, ie.dl_ReceptionLBT_subsetRB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_ReceptionIntraCellGuardband_r16 != nil {
		if err = ie.dl_ReceptionIntraCellGuardband_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_ReceptionIntraCellGuardband_r16", err)
		}
	}
	if ie.dl_ReceptionLBT_subsetRB_r16 != nil {
		if err = ie.dl_ReceptionLBT_subsetRB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_ReceptionLBT_subsetRB_r16", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1630) Decode(r *uper.UperReader) error {
	var err error
	var dl_ReceptionIntraCellGuardband_r16Present bool
	if dl_ReceptionIntraCellGuardband_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_ReceptionLBT_subsetRB_r16Present bool
	if dl_ReceptionLBT_subsetRB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_ReceptionIntraCellGuardband_r16Present {
		ie.dl_ReceptionIntraCellGuardband_r16 = new(SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionIntraCellGuardband_r16)
		if err = ie.dl_ReceptionIntraCellGuardband_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_ReceptionIntraCellGuardband_r16", err)
		}
	}
	if dl_ReceptionLBT_subsetRB_r16Present {
		ie.dl_ReceptionLBT_subsetRB_r16 = new(SharedSpectrumChAccessParamsPerBand_v1630_dl_ReceptionLBT_subsetRB_r16)
		if err = ie.dl_ReceptionLBT_subsetRB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_ReceptionLBT_subsetRB_r16", err)
		}
	}
	return nil
}
