package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NZP_CSI_RS_Resource struct {
	nzp_CSI_RS_ResourceId  NZP_CSI_RS_ResourceId                     `madatory`
	resourceMapping        CSI_RS_ResourceMapping                    `madatory`
	powerControlOffset     int64                                     `lb:-8,ub:15,madatory`
	powerControlOffsetSS   *NZP_CSI_RS_Resource_powerControlOffsetSS `optional`
	scramblingID           ScramblingId                              `madatory`
	periodicityAndOffset   *CSI_ResourcePeriodicityAndOffset         `optional`
	qcl_InfoPeriodicCSI_RS *TCI_StateId                              `optional`
}

func (ie *NZP_CSI_RS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.powerControlOffsetSS != nil, ie.periodicityAndOffset != nil, ie.qcl_InfoPeriodicCSI_RS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.nzp_CSI_RS_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode nzp_CSI_RS_ResourceId", err)
	}
	if err = ie.resourceMapping.Encode(w); err != nil {
		return utils.WrapError("Encode resourceMapping", err)
	}
	if err = w.WriteInteger(ie.powerControlOffset, &uper.Constraint{Lb: -8, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger powerControlOffset", err)
	}
	if ie.powerControlOffsetSS != nil {
		if err = ie.powerControlOffsetSS.Encode(w); err != nil {
			return utils.WrapError("Encode powerControlOffsetSS", err)
		}
	}
	if err = ie.scramblingID.Encode(w); err != nil {
		return utils.WrapError("Encode scramblingID", err)
	}
	if ie.periodicityAndOffset != nil {
		if err = ie.periodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndOffset", err)
		}
	}
	if ie.qcl_InfoPeriodicCSI_RS != nil {
		if err = ie.qcl_InfoPeriodicCSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode qcl_InfoPeriodicCSI_RS", err)
		}
	}
	return nil
}

func (ie *NZP_CSI_RS_Resource) Decode(r *uper.UperReader) error {
	var err error
	var powerControlOffsetSSPresent bool
	if powerControlOffsetSSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var periodicityAndOffsetPresent bool
	if periodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var qcl_InfoPeriodicCSI_RSPresent bool
	if qcl_InfoPeriodicCSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.nzp_CSI_RS_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode nzp_CSI_RS_ResourceId", err)
	}
	if err = ie.resourceMapping.Decode(r); err != nil {
		return utils.WrapError("Decode resourceMapping", err)
	}
	var tmp_int_powerControlOffset int64
	if tmp_int_powerControlOffset, err = r.ReadInteger(&uper.Constraint{Lb: -8, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger powerControlOffset", err)
	}
	ie.powerControlOffset = tmp_int_powerControlOffset
	if powerControlOffsetSSPresent {
		ie.powerControlOffsetSS = new(NZP_CSI_RS_Resource_powerControlOffsetSS)
		if err = ie.powerControlOffsetSS.Decode(r); err != nil {
			return utils.WrapError("Decode powerControlOffsetSS", err)
		}
	}
	if err = ie.scramblingID.Decode(r); err != nil {
		return utils.WrapError("Decode scramblingID", err)
	}
	if periodicityAndOffsetPresent {
		ie.periodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
		if err = ie.periodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndOffset", err)
		}
	}
	if qcl_InfoPeriodicCSI_RSPresent {
		ie.qcl_InfoPeriodicCSI_RS = new(TCI_StateId)
		if err = ie.qcl_InfoPeriodicCSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode qcl_InfoPeriodicCSI_RS", err)
		}
	}
	return nil
}
