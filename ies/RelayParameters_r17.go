package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RelayParameters_r17 struct {
	relayUE_Operation_L2_r17                   *RelayParameters_r17_relayUE_Operation_L2_r17                   `optional`
	remoteUE_Operation_L2_r17                  *RelayParameters_r17_remoteUE_Operation_L2_r17                  `optional`
	remoteUE_PathSwitchToIdleInactiveRelay_r17 *RelayParameters_r17_remoteUE_PathSwitchToIdleInactiveRelay_r17 `optional`
}

func (ie *RelayParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.relayUE_Operation_L2_r17 != nil, ie.remoteUE_Operation_L2_r17 != nil, ie.remoteUE_PathSwitchToIdleInactiveRelay_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.relayUE_Operation_L2_r17 != nil {
		if err = ie.relayUE_Operation_L2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode relayUE_Operation_L2_r17", err)
		}
	}
	if ie.remoteUE_Operation_L2_r17 != nil {
		if err = ie.remoteUE_Operation_L2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode remoteUE_Operation_L2_r17", err)
		}
	}
	if ie.remoteUE_PathSwitchToIdleInactiveRelay_r17 != nil {
		if err = ie.remoteUE_PathSwitchToIdleInactiveRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode remoteUE_PathSwitchToIdleInactiveRelay_r17", err)
		}
	}
	return nil
}

func (ie *RelayParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var relayUE_Operation_L2_r17Present bool
	if relayUE_Operation_L2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var remoteUE_Operation_L2_r17Present bool
	if remoteUE_Operation_L2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var remoteUE_PathSwitchToIdleInactiveRelay_r17Present bool
	if remoteUE_PathSwitchToIdleInactiveRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if relayUE_Operation_L2_r17Present {
		ie.relayUE_Operation_L2_r17 = new(RelayParameters_r17_relayUE_Operation_L2_r17)
		if err = ie.relayUE_Operation_L2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode relayUE_Operation_L2_r17", err)
		}
	}
	if remoteUE_Operation_L2_r17Present {
		ie.remoteUE_Operation_L2_r17 = new(RelayParameters_r17_remoteUE_Operation_L2_r17)
		if err = ie.remoteUE_Operation_L2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode remoteUE_Operation_L2_r17", err)
		}
	}
	if remoteUE_PathSwitchToIdleInactiveRelay_r17Present {
		ie.remoteUE_PathSwitchToIdleInactiveRelay_r17 = new(RelayParameters_r17_remoteUE_PathSwitchToIdleInactiveRelay_r17)
		if err = ie.remoteUE_PathSwitchToIdleInactiveRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode remoteUE_PathSwitchToIdleInactiveRelay_r17", err)
		}
	}
	return nil
}
