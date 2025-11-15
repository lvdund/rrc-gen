package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyG struct {
	maxNumberSSB_CSI_RS_ResourceOneTx DummyG_maxNumberSSB_CSI_RS_ResourceOneTx `madatory`
	maxNumberSSB_CSI_RS_ResourceTwoTx DummyG_maxNumberSSB_CSI_RS_ResourceTwoTx `madatory`
	supportedCSI_RS_Density           DummyG_supportedCSI_RS_Density           `madatory`
}

func (ie *DummyG) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberSSB_CSI_RS_ResourceOneTx.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.maxNumberSSB_CSI_RS_ResourceTwoTx.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSSB_CSI_RS_ResourceTwoTx", err)
	}
	if err = ie.supportedCSI_RS_Density.Encode(w); err != nil {
		return utils.WrapError("Encode supportedCSI_RS_Density", err)
	}
	return nil
}

func (ie *DummyG) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberSSB_CSI_RS_ResourceOneTx.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSSB_CSI_RS_ResourceOneTx", err)
	}
	if err = ie.maxNumberSSB_CSI_RS_ResourceTwoTx.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSSB_CSI_RS_ResourceTwoTx", err)
	}
	if err = ie.supportedCSI_RS_Density.Decode(r); err != nil {
		return utils.WrapError("Decode supportedCSI_RS_Density", err)
	}
	return nil
}
