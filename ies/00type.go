package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"

	"github.com/lvdund/rrc/utils"
)

type AI_RNTI_r16 struct {
	RNTI_Value
}

func (ie *AI_RNTI_r16) Encode(w *uper.UperWriter) error {
	if err := ie.RNTI_Value.Encode(w); err != nil {
		return utils.WrapError("Encode AI_RNTI_r16", err)
	}
	return nil
}

func (ie *AI_RNTI_r16) Decode(r *uper.UperReader) error {
	if err := ie.RNTI_Value.Decode(r); err != nil {
		return utils.WrapError("Decode AI_RNTI_r16", err)
	}
	return nil
}

type MobilityHistoryReport_r16 struct {
	VisitedCellInfoList_r16
}

func (ie *MobilityHistoryReport_r16) Encode(w *uper.UperWriter) error {
	if err := ie.VisitedCellInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MobilityHistoryReport_r16", err)
	}
	return nil
}

func (ie *MobilityHistoryReport_r16) Decode(r *uper.UperReader) error {
	if err := ie.VisitedCellInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MobilityHistoryReport_r16", err)
	}
	return nil
}

type VarMobilityHistoryReport_r16 struct {
	VisitedCellInfoList_r16
}

func (ie *VarMobilityHistoryReport_r16) Encode(w *uper.UperWriter) error {
	if err := ie.VisitedCellInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode VarMobilityHistoryReport_r16", err)
	}
	return nil
}

func (ie *VarMobilityHistoryReport_r16) Decode(r *uper.UperReader) error {
	if err := ie.VisitedCellInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode VarMobilityHistoryReport_r16", err)
	}
	return nil
}

type RangeToBestCell struct {
	Q_OffsetRange
}

func (ie *RangeToBestCell) Encode(w *uper.UperWriter) error {
	if err := ie.Q_OffsetRange.Encode(w); err != nil {
		return utils.WrapError("Encode RangeToBestCell", err)
	}
	return nil
}

func (ie *RangeToBestCell) Decode(r *uper.UperReader) error {
	if err := ie.Q_OffsetRange.Decode(r); err != nil {
		return utils.WrapError("Decode RangeToBestCell", err)
	}
	return nil
}

const (
	PDSCH_HARQ_ACK_CodebookList_r16_Enum_semiStatic uper.Enumerated = 0
	PDSCH_HARQ_ACK_CodebookList_r16_Enum_dynamic    uper.Enumerated = 1
)

type PDSCH_HARQ_ACK_CodebookList_r16 struct {
	Value []uper.Enumerated `lb:1,ub:2,madatory`
}

func (ie *PDSCH_HARQ_ACK_CodebookList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*utils.ENUMERATED]([]*utils.ENUMERATED{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp_ie := utils.NewENUMERATED(int64(i), uper.Constraint{Lb: 0, Ub: 1}, false)
		tmp.Value = append(tmp.Value, &tmp_ie)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PDSCH_HARQ_ACK_CodebookList_r16", err)
	}
	return nil
}

func (ie *PDSCH_HARQ_ACK_CodebookList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*utils.ENUMERATED]([]*utils.ENUMERATED{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	fn := func() *utils.ENUMERATED {
		ie := utils.NewENUMERATED(0, uper.Constraint{Lb: 0, Ub: 1}, false)
		return &ie
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PDSCH_HARQ_ACK_CodebookList_r16", err)
	}
	ie.Value = []uper.Enumerated{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, i.Value)
	}
	return err
}

type MeasurementTimingConfiguration_v1610_IEs struct {
	csi_RS_Config_r16    MeasurementTimingConfiguration_v1610_IEs_csi_RS_Config_r16 `madatory`
	nonCriticalExtension interface{}                                                `optional`
}

func (ie *MeasurementTimingConfiguration_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.csi_RS_Config_r16.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_Config_r16", err)
	}
	return err
}

func (ie *MeasurementTimingConfiguration_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	if _, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.csi_RS_Config_r16.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS_Config_r16", err)
	}
	return nil
}

type PDCP_Config_moreThanTwoRLC_DRB_r16 struct {
	splitSecondaryPath_r16 *LogicalChannelIdentity `optional`
	duplicationState_r16   []bool                  `lb:3,ub:3,optional`
}

func (ie *PDCP_Config_moreThanTwoRLC_DRB_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.splitSecondaryPath_r16 != nil, len(ie.duplicationState_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.splitSecondaryPath_r16 != nil {
		if err = ie.splitSecondaryPath_r16.Encode(w); err != nil {
			return utils.WrapError("Encode splitSecondaryPath_r16", err)
		}
	}
	if len(ie.duplicationState_r16) > 0 {
		tmp_duplicationState_r16 := utils.NewSequence[*utils.BOOLEAN]([]*utils.BOOLEAN{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		for _, i := range ie.duplicationState_r16 {
			tmp_duplicationState_r16.Value = append(tmp_duplicationState_r16.Value, &utils.BOOLEAN{Value: i})
		}
		if err = tmp_duplicationState_r16.Encode(w); err != nil {
			return utils.WrapError("Encode duplicationState_r16", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_moreThanTwoRLC_DRB_r16) Decode(r *uper.UperReader) error {
	var err error
	var splitSecondaryPath_r16Present bool
	if splitSecondaryPath_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var duplicationState_r16Present bool
	if duplicationState_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if splitSecondaryPath_r16Present {
		ie.splitSecondaryPath_r16 = new(LogicalChannelIdentity)
		if err = ie.splitSecondaryPath_r16.Decode(r); err != nil {
			return utils.WrapError("Decode splitSecondaryPath_r16", err)
		}
	}
	if duplicationState_r16Present {
		tmp_duplicationState_r16 := utils.NewSequence[*utils.BOOLEAN]([]*utils.BOOLEAN{}, uper.Constraint{Lb: 3, Ub: 3}, false)
		fn_duplicationState_r16 := func() *utils.BOOLEAN {
			return new(utils.BOOLEAN)
		}
		if err = tmp_duplicationState_r16.Decode(r, fn_duplicationState_r16); err != nil {
			return utils.WrapError("Decode duplicationState_r16", err)
		}
		ie.duplicationState_r16 = []bool{}
		for _, i := range tmp_duplicationState_r16.Value {
			ie.duplicationState_r16 = append(ie.duplicationState_r16, i.Value)
		}
	}
	return nil
}

type RLF_TimersAndConstants struct {
	t310 RLF_TimersAndConstants_t310  `madatory`
	n310 RLF_TimersAndConstants_n310  `madatory`
	n311 RLF_TimersAndConstants_n311  `madatory`
	t311 *RLF_TimersAndConstants_t311 `optional,ext-1`
}

func (ie *RLF_TimersAndConstants) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.t311 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.t310.Encode(w); err != nil {
		return utils.WrapError("Encode t310", err)
	}
	if err = ie.n310.Encode(w); err != nil {
		return utils.WrapError("Encode n310", err)
	}
	if err = ie.n311.Encode(w); err != nil {
		return utils.WrapError("Encode n311", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.t311 != nil}
		for _, bit := range extBitmap {
			if err = w.WriteBool(bit); err != nil {
				return utils.WrapError("WriteExtBitMap RLF_TimersAndConstants", err)
			}
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// encode t311
			if err = ie.t311.Encode(extWriter); err != nil {
				return utils.WrapError("Encode t311", err)
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *RLF_TimersAndConstants) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.t310.Decode(r); err != nil {
		return utils.WrapError("Decode t310", err)
	}
	if err = ie.n310.Decode(r); err != nil {
		return utils.WrapError("Decode n310", err)
	}
	if err = ie.n311.Decode(r); err != nil {
		return utils.WrapError("Decode n311", err)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap := make([]bool, 1)
		for i := range extBitmap {
			if extBitmap[i], err = r.ReadBool(); err != nil {
				return err
			}
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			// decode t311
			if err = ie.t311.Decode(extReader); err != nil {
				return utils.WrapError("Decode t311", err)
			}
		}
	}
	return nil
}
