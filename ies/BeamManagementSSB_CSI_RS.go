package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamManagementSSB_CSI_RS struct {
	maxNumberSSB_CSI_RS_ResourceOneTx BeamManagementSSB_CSI_RS_maxNumberSSB_CSI_RS_ResourceOneTx `madatory`
	maxNumberCSI_RS_Resource          BeamManagementSSB_CSI_RS_maxNumberCSI_RS_Resource          `madatory`
	maxNumberCSI_RS_ResourceTwoTx     BeamManagementSSB_CSI_RS_maxNumberCSI_RS_ResourceTwoTx     `madatory`
	supportedCSI_RS_Density           *BeamManagementSSB_CSI_RS_supportedCSI_RS_Density          `optional`
	maxNumberAperiodicCSI_RS_Resource BeamManagementSSB_CSI_RS_maxNumberAperiodicCSI_RS_Resource `madatory`
}

func (ie *BeamManagementSSB_CSI_RS) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.supportedCSI_RS_Density != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.maxNumberSSB_CSI_RS_ResourceOneTx.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.maxNumberCSI_RS_Resource.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberCSI_RS_Resource", err)
	}
	if err = ie.maxNumberCSI_RS_ResourceTwoTx.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberCSI_RS_ResourceTwoTx", err)
	}
	if ie.supportedCSI_RS_Density != nil {
		if err = ie.supportedCSI_RS_Density.Encode(w); err != nil {
			return utils.WrapError("Encode supportedCSI_RS_Density", err)
		}
	}
	if err = ie.maxNumberAperiodicCSI_RS_Resource.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberAperiodicCSI_RS_Resource", err)
	}
	return nil
}

func (ie *BeamManagementSSB_CSI_RS) Decode(r *uper.UperReader) error {
	var err error
	var supportedCSI_RS_DensityPresent bool
	if supportedCSI_RS_DensityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.maxNumberSSB_CSI_RS_ResourceOneTx.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.maxNumberCSI_RS_Resource.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberCSI_RS_Resource", err)
	}
	if err = ie.maxNumberCSI_RS_ResourceTwoTx.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberCSI_RS_ResourceTwoTx", err)
	}
	if supportedCSI_RS_DensityPresent {
		ie.supportedCSI_RS_Density = new(BeamManagementSSB_CSI_RS_supportedCSI_RS_Density)
		if err = ie.supportedCSI_RS_Density.Decode(r); err != nil {
			return utils.WrapError("Decode supportedCSI_RS_Density", err)
		}
	}
	if err = ie.maxNumberAperiodicCSI_RS_Resource.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberAperiodicCSI_RS_Resource", err)
	}
	return nil
}
