package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_CellMobility struct {
	cellId                       PhysCellId                               `madatory`
	csi_rs_MeasurementBW         CSI_RS_CellMobility_csi_rs_MeasurementBW `lb:0,ub:2169,madatory`
	density                      *CSI_RS_CellMobility_density             `optional`
	csi_rs_ResourceList_Mobility []CSI_RS_Resource_Mobility               `lb:1,ub:maxNrofCSI_RS_ResourcesRRM,madatory`
}

func (ie *CSI_RS_CellMobility) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.density != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.cellId.Encode(w); err != nil {
		return utils.WrapError("Encode cellId", err)
	}
	if err = ie.csi_rs_MeasurementBW.Encode(w); err != nil {
		return utils.WrapError("Encode csi_rs_MeasurementBW", err)
	}
	if ie.density != nil {
		if err = ie.density.Encode(w); err != nil {
			return utils.WrapError("Encode density", err)
		}
	}
	tmp_csi_rs_ResourceList_Mobility := utils.NewSequence[*CSI_RS_Resource_Mobility]([]*CSI_RS_Resource_Mobility{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesRRM}, false)
	for _, i := range ie.csi_rs_ResourceList_Mobility {
		tmp_csi_rs_ResourceList_Mobility.Value = append(tmp_csi_rs_ResourceList_Mobility.Value, &i)
	}
	if err = tmp_csi_rs_ResourceList_Mobility.Encode(w); err != nil {
		return utils.WrapError("Encode csi_rs_ResourceList_Mobility", err)
	}
	return nil
}

func (ie *CSI_RS_CellMobility) Decode(r *uper.UperReader) error {
	var err error
	var densityPresent bool
	if densityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.cellId.Decode(r); err != nil {
		return utils.WrapError("Decode cellId", err)
	}
	if err = ie.csi_rs_MeasurementBW.Decode(r); err != nil {
		return utils.WrapError("Decode csi_rs_MeasurementBW", err)
	}
	if densityPresent {
		ie.density = new(CSI_RS_CellMobility_density)
		if err = ie.density.Decode(r); err != nil {
			return utils.WrapError("Decode density", err)
		}
	}
	tmp_csi_rs_ResourceList_Mobility := utils.NewSequence[*CSI_RS_Resource_Mobility]([]*CSI_RS_Resource_Mobility{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_ResourcesRRM}, false)
	fn_csi_rs_ResourceList_Mobility := func() *CSI_RS_Resource_Mobility {
		return new(CSI_RS_Resource_Mobility)
	}
	if err = tmp_csi_rs_ResourceList_Mobility.Decode(r, fn_csi_rs_ResourceList_Mobility); err != nil {
		return utils.WrapError("Decode csi_rs_ResourceList_Mobility", err)
	}
	ie.csi_rs_ResourceList_Mobility = []CSI_RS_Resource_Mobility{}
	for _, i := range tmp_csi_rs_ResourceList_Mobility.Value {
		ie.csi_rs_ResourceList_Mobility = append(ie.csi_rs_ResourceList_Mobility, *i)
	}
	return nil
}
