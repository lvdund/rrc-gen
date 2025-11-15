package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResume_v1610_IEs struct {
	idleModeMeasurementReq_r16  *RRCResume_v1610_IEs_idleModeMeasurementReq_r16  `optional`
	restoreMCG_SCells_r16       *RRCResume_v1610_IEs_restoreMCG_SCells_r16       `optional`
	restoreSCG_r16              *RRCResume_v1610_IEs_restoreSCG_r16              `optional`
	mrdc_SecondaryCellGroup_r16 *RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16 `optional`
	needForGapsConfigNR_r16     *NeedForGapsConfigNR_r16                         `optional,setuprelease`
	nonCriticalExtension        *RRCResume_v1700_IEs                             `optional`
}

func (ie *RRCResume_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.idleModeMeasurementReq_r16 != nil, ie.restoreMCG_SCells_r16 != nil, ie.restoreSCG_r16 != nil, ie.mrdc_SecondaryCellGroup_r16 != nil, ie.needForGapsConfigNR_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.idleModeMeasurementReq_r16 != nil {
		if err = ie.idleModeMeasurementReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleModeMeasurementReq_r16", err)
		}
	}
	if ie.restoreMCG_SCells_r16 != nil {
		if err = ie.restoreMCG_SCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode restoreMCG_SCells_r16", err)
		}
	}
	if ie.restoreSCG_r16 != nil {
		if err = ie.restoreSCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode restoreSCG_r16", err)
		}
	}
	if ie.mrdc_SecondaryCellGroup_r16 != nil {
		if err = ie.mrdc_SecondaryCellGroup_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_SecondaryCellGroup_r16", err)
		}
	}
	if ie.needForGapsConfigNR_r16 != nil {
		tmp_needForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{
			Setup: ie.needForGapsConfigNR_r16,
		}
		if err = tmp_needForGapsConfigNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapsConfigNR_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCResume_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var idleModeMeasurementReq_r16Present bool
	if idleModeMeasurementReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var restoreMCG_SCells_r16Present bool
	if restoreMCG_SCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var restoreSCG_r16Present bool
	if restoreSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_SecondaryCellGroup_r16Present bool
	if mrdc_SecondaryCellGroup_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapsConfigNR_r16Present bool
	if needForGapsConfigNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if idleModeMeasurementReq_r16Present {
		ie.idleModeMeasurementReq_r16 = new(RRCResume_v1610_IEs_idleModeMeasurementReq_r16)
		if err = ie.idleModeMeasurementReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleModeMeasurementReq_r16", err)
		}
	}
	if restoreMCG_SCells_r16Present {
		ie.restoreMCG_SCells_r16 = new(RRCResume_v1610_IEs_restoreMCG_SCells_r16)
		if err = ie.restoreMCG_SCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode restoreMCG_SCells_r16", err)
		}
	}
	if restoreSCG_r16Present {
		ie.restoreSCG_r16 = new(RRCResume_v1610_IEs_restoreSCG_r16)
		if err = ie.restoreSCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode restoreSCG_r16", err)
		}
	}
	if mrdc_SecondaryCellGroup_r16Present {
		ie.mrdc_SecondaryCellGroup_r16 = new(RRCResume_v1610_IEs_mrdc_SecondaryCellGroup_r16)
		if err = ie.mrdc_SecondaryCellGroup_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_SecondaryCellGroup_r16", err)
		}
	}
	if needForGapsConfigNR_r16Present {
		tmp_needForGapsConfigNR_r16 := utils.SetupRelease[*NeedForGapsConfigNR_r16]{}
		if err = tmp_needForGapsConfigNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapsConfigNR_r16", err)
		}
		ie.needForGapsConfigNR_r16 = tmp_needForGapsConfigNR_r16.Setup
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCResume_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
