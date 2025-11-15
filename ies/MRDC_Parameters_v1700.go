package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1700 struct {
	condPSCellAdditionENDC_r17               *MRDC_Parameters_v1700_condPSCellAdditionENDC_r17               `optional`
	scg_ActivationDeactivationENDC_r17       *MRDC_Parameters_v1700_scg_ActivationDeactivationENDC_r17       `optional`
	scg_ActivationDeactivationResumeENDC_r17 *MRDC_Parameters_v1700_scg_ActivationDeactivationResumeENDC_r17 `optional`
}

func (ie *MRDC_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.condPSCellAdditionENDC_r17 != nil, ie.scg_ActivationDeactivationENDC_r17 != nil, ie.scg_ActivationDeactivationResumeENDC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.condPSCellAdditionENDC_r17 != nil {
		if err = ie.condPSCellAdditionENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode condPSCellAdditionENDC_r17", err)
		}
	}
	if ie.scg_ActivationDeactivationENDC_r17 != nil {
		if err = ie.scg_ActivationDeactivationENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scg_ActivationDeactivationENDC_r17", err)
		}
	}
	if ie.scg_ActivationDeactivationResumeENDC_r17 != nil {
		if err = ie.scg_ActivationDeactivationResumeENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scg_ActivationDeactivationResumeENDC_r17", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var condPSCellAdditionENDC_r17Present bool
	if condPSCellAdditionENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_ActivationDeactivationENDC_r17Present bool
	if scg_ActivationDeactivationENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_ActivationDeactivationResumeENDC_r17Present bool
	if scg_ActivationDeactivationResumeENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if condPSCellAdditionENDC_r17Present {
		ie.condPSCellAdditionENDC_r17 = new(MRDC_Parameters_v1700_condPSCellAdditionENDC_r17)
		if err = ie.condPSCellAdditionENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condPSCellAdditionENDC_r17", err)
		}
	}
	if scg_ActivationDeactivationENDC_r17Present {
		ie.scg_ActivationDeactivationENDC_r17 = new(MRDC_Parameters_v1700_scg_ActivationDeactivationENDC_r17)
		if err = ie.scg_ActivationDeactivationENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scg_ActivationDeactivationENDC_r17", err)
		}
	}
	if scg_ActivationDeactivationResumeENDC_r17Present {
		ie.scg_ActivationDeactivationResumeENDC_r17 = new(MRDC_Parameters_v1700_scg_ActivationDeactivationResumeENDC_r17)
		if err = ie.scg_ActivationDeactivationResumeENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scg_ActivationDeactivationResumeENDC_r17", err)
		}
	}
	return nil
}
