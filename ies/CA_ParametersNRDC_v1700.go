package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1700 struct {
	simultaneousRxTx_IAB_MultipleParents_r17 *CA_ParametersNRDC_v1700_simultaneousRxTx_IAB_MultipleParents_r17 `optional`
	condPSCellAdditionNRDC_r17               *CA_ParametersNRDC_v1700_condPSCellAdditionNRDC_r17               `optional`
	scg_ActivationDeactivationNRDC_r17       *CA_ParametersNRDC_v1700_scg_ActivationDeactivationNRDC_r17       `optional`
	scg_ActivationDeactivationResumeNRDC_r17 *CA_ParametersNRDC_v1700_scg_ActivationDeactivationResumeNRDC_r17 `optional`
	beamManagementType_CBM_r17               *CA_ParametersNRDC_v1700_beamManagementType_CBM_r17               `optional`
}

func (ie *CA_ParametersNRDC_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.simultaneousRxTx_IAB_MultipleParents_r17 != nil, ie.condPSCellAdditionNRDC_r17 != nil, ie.scg_ActivationDeactivationNRDC_r17 != nil, ie.scg_ActivationDeactivationResumeNRDC_r17 != nil, ie.beamManagementType_CBM_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.simultaneousRxTx_IAB_MultipleParents_r17 != nil {
		if err = ie.simultaneousRxTx_IAB_MultipleParents_r17.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousRxTx_IAB_MultipleParents_r17", err)
		}
	}
	if ie.condPSCellAdditionNRDC_r17 != nil {
		if err = ie.condPSCellAdditionNRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode condPSCellAdditionNRDC_r17", err)
		}
	}
	if ie.scg_ActivationDeactivationNRDC_r17 != nil {
		if err = ie.scg_ActivationDeactivationNRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scg_ActivationDeactivationNRDC_r17", err)
		}
	}
	if ie.scg_ActivationDeactivationResumeNRDC_r17 != nil {
		if err = ie.scg_ActivationDeactivationResumeNRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scg_ActivationDeactivationResumeNRDC_r17", err)
		}
	}
	if ie.beamManagementType_CBM_r17 != nil {
		if err = ie.beamManagementType_CBM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode beamManagementType_CBM_r17", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1700) Decode(r *uper.UperReader) error {
	var err error
	var simultaneousRxTx_IAB_MultipleParents_r17Present bool
	if simultaneousRxTx_IAB_MultipleParents_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var condPSCellAdditionNRDC_r17Present bool
	if condPSCellAdditionNRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_ActivationDeactivationNRDC_r17Present bool
	if scg_ActivationDeactivationNRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_ActivationDeactivationResumeNRDC_r17Present bool
	if scg_ActivationDeactivationResumeNRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var beamManagementType_CBM_r17Present bool
	if beamManagementType_CBM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if simultaneousRxTx_IAB_MultipleParents_r17Present {
		ie.simultaneousRxTx_IAB_MultipleParents_r17 = new(CA_ParametersNRDC_v1700_simultaneousRxTx_IAB_MultipleParents_r17)
		if err = ie.simultaneousRxTx_IAB_MultipleParents_r17.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousRxTx_IAB_MultipleParents_r17", err)
		}
	}
	if condPSCellAdditionNRDC_r17Present {
		ie.condPSCellAdditionNRDC_r17 = new(CA_ParametersNRDC_v1700_condPSCellAdditionNRDC_r17)
		if err = ie.condPSCellAdditionNRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condPSCellAdditionNRDC_r17", err)
		}
	}
	if scg_ActivationDeactivationNRDC_r17Present {
		ie.scg_ActivationDeactivationNRDC_r17 = new(CA_ParametersNRDC_v1700_scg_ActivationDeactivationNRDC_r17)
		if err = ie.scg_ActivationDeactivationNRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scg_ActivationDeactivationNRDC_r17", err)
		}
	}
	if scg_ActivationDeactivationResumeNRDC_r17Present {
		ie.scg_ActivationDeactivationResumeNRDC_r17 = new(CA_ParametersNRDC_v1700_scg_ActivationDeactivationResumeNRDC_r17)
		if err = ie.scg_ActivationDeactivationResumeNRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scg_ActivationDeactivationResumeNRDC_r17", err)
		}
	}
	if beamManagementType_CBM_r17Present {
		ie.beamManagementType_CBM_r17 = new(CA_ParametersNRDC_v1700_beamManagementType_CBM_r17)
		if err = ie.beamManagementType_CBM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode beamManagementType_CBM_r17", err)
		}
	}
	return nil
}
