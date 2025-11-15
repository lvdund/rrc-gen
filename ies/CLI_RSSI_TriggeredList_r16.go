package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CLI_RSSI_TriggeredList_r16 struct {
	Value []RSSI_ResourceId_r16 `lb:1,ub:maxNrofCLI_RSSI_Resources_r16,madatory`
}

func (ie *CLI_RSSI_TriggeredList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*RSSI_ResourceId_r16]([]*RSSI_ResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCLI_RSSI_Resources_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CLI_RSSI_TriggeredList_r16", err)
	}
	return nil
}

func (ie *CLI_RSSI_TriggeredList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*RSSI_ResourceId_r16]([]*RSSI_ResourceId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCLI_RSSI_Resources_r16}, false)
	fn := func() *RSSI_ResourceId_r16 {
		return new(RSSI_ResourceId_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CLI_RSSI_TriggeredList_r16", err)
	}
	ie.Value = []RSSI_ResourceId_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
