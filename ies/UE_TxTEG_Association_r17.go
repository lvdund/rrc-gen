package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_TxTEG_Association_r17 struct {
	ue_TxTEG_ID_r17                     int64                   `lb:0,ub:maxNrOfTxTEG_ID_1_r17,madatory`
	nr_TimeStamp_r17                    NR_TimeStamp_r17        `madatory`
	associatedSRS_PosResourceIdList_r17 []SRS_PosResourceId_r16 `lb:1,ub:maxNrofSRS_PosResources_r16,madatory`
	servCellId_r17                      *ServCellIndex          `optional`
}

func (ie *UE_TxTEG_Association_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servCellId_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.ue_TxTEG_ID_r17, &uper.Constraint{Lb: 0, Ub: maxNrOfTxTEG_ID_1_r17}, false); err != nil {
		return utils.WrapError("WriteInteger ue_TxTEG_ID_r17", err)
	}
	if err = ie.nr_TimeStamp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nr_TimeStamp_r17", err)
	}
	tmp_associatedSRS_PosResourceIdList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
	for _, i := range ie.associatedSRS_PosResourceIdList_r17 {
		tmp_associatedSRS_PosResourceIdList_r17.Value = append(tmp_associatedSRS_PosResourceIdList_r17.Value, &i)
	}
	if err = tmp_associatedSRS_PosResourceIdList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode associatedSRS_PosResourceIdList_r17", err)
	}
	if ie.servCellId_r17 != nil {
		if err = ie.servCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode servCellId_r17", err)
		}
	}
	return nil
}

func (ie *UE_TxTEG_Association_r17) Decode(r *uper.UperReader) error {
	var err error
	var servCellId_r17Present bool
	if servCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_ue_TxTEG_ID_r17 int64
	if tmp_int_ue_TxTEG_ID_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrOfTxTEG_ID_1_r17}, false); err != nil {
		return utils.WrapError("ReadInteger ue_TxTEG_ID_r17", err)
	}
	ie.ue_TxTEG_ID_r17 = tmp_int_ue_TxTEG_ID_r17
	if err = ie.nr_TimeStamp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode nr_TimeStamp_r17", err)
	}
	tmp_associatedSRS_PosResourceIdList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
	fn_associatedSRS_PosResourceIdList_r17 := func() *SRS_PosResourceId_r16 {
		return new(SRS_PosResourceId_r16)
	}
	if err = tmp_associatedSRS_PosResourceIdList_r17.Decode(r, fn_associatedSRS_PosResourceIdList_r17); err != nil {
		return utils.WrapError("Decode associatedSRS_PosResourceIdList_r17", err)
	}
	ie.associatedSRS_PosResourceIdList_r17 = []SRS_PosResourceId_r16{}
	for _, i := range tmp_associatedSRS_PosResourceIdList_r17.Value {
		ie.associatedSRS_PosResourceIdList_r17 = append(ie.associatedSRS_PosResourceIdList_r17, *i)
	}
	if servCellId_r17Present {
		ie.servCellId_r17 = new(ServCellIndex)
		if err = ie.servCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode servCellId_r17", err)
		}
	}
	return nil
}
