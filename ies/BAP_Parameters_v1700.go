package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BAP_Parameters_v1700 struct {
	bapHeaderRewriting_Rerouting_r17 *BAP_Parameters_v1700_bapHeaderRewriting_Rerouting_r17 `optional`
	bapHeaderRewriting_Routing_r17   *BAP_Parameters_v1700_bapHeaderRewriting_Routing_r17   `optional`
}

func (ie *BAP_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bapHeaderRewriting_Rerouting_r17 != nil, ie.bapHeaderRewriting_Routing_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bapHeaderRewriting_Rerouting_r17 != nil {
		if err = ie.bapHeaderRewriting_Rerouting_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bapHeaderRewriting_Rerouting_r17", err)
		}
	}
	if ie.bapHeaderRewriting_Routing_r17 != nil {
		if err = ie.bapHeaderRewriting_Routing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode bapHeaderRewriting_Routing_r17", err)
		}
	}
	return nil
}

func (ie *BAP_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var bapHeaderRewriting_Rerouting_r17Present bool
	if bapHeaderRewriting_Rerouting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bapHeaderRewriting_Routing_r17Present bool
	if bapHeaderRewriting_Routing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bapHeaderRewriting_Rerouting_r17Present {
		ie.bapHeaderRewriting_Rerouting_r17 = new(BAP_Parameters_v1700_bapHeaderRewriting_Rerouting_r17)
		if err = ie.bapHeaderRewriting_Rerouting_r17.Decode(r); err != nil {
			return utils.WrapError("Decode bapHeaderRewriting_Rerouting_r17", err)
		}
	}
	if bapHeaderRewriting_Routing_r17Present {
		ie.bapHeaderRewriting_Routing_r17 = new(BAP_Parameters_v1700_bapHeaderRewriting_Routing_r17)
		if err = ie.bapHeaderRewriting_Routing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode bapHeaderRewriting_Routing_r17", err)
		}
	}
	return nil
}
