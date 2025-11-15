package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16 struct {
	supportPDCCH_ToPDSCH_r16    *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDCCH_ToPDSCH_r16    `optional`
	supportPDSCH_ToHARQ_ACK_r16 *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDSCH_ToHARQ_ACK_r16 `optional`
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportPDCCH_ToPDSCH_r16 != nil, ie.supportPDSCH_ToHARQ_ACK_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportPDCCH_ToPDSCH_r16 != nil {
		if err = ie.supportPDCCH_ToPDSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportPDCCH_ToPDSCH_r16", err)
		}
	}
	if ie.supportPDSCH_ToHARQ_ACK_r16 != nil {
		if err = ie.supportPDSCH_ToHARQ_ACK_r16.Encode(w); err != nil {
			return utils.WrapError("Encode supportPDSCH_ToHARQ_ACK_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16) Decode(r *uper.UperReader) error {
	var err error
	var supportPDCCH_ToPDSCH_r16Present bool
	if supportPDCCH_ToPDSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportPDSCH_ToHARQ_ACK_r16Present bool
	if supportPDSCH_ToHARQ_ACK_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if supportPDCCH_ToPDSCH_r16Present {
		ie.supportPDCCH_ToPDSCH_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDCCH_ToPDSCH_r16)
		if err = ie.supportPDCCH_ToPDSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportPDCCH_ToPDSCH_r16", err)
		}
	}
	if supportPDSCH_ToHARQ_ACK_r16Present {
		ie.supportPDSCH_ToHARQ_ACK_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16_supportPDSCH_ToHARQ_ACK_r16)
		if err = ie.supportPDSCH_ToHARQ_ACK_r16.Decode(r); err != nil {
			return utils.WrapError("Decode supportPDSCH_ToHARQ_ACK_r16", err)
		}
	}
	return nil
}
