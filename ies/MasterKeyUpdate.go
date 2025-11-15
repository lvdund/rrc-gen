package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MasterKeyUpdate struct {
	keySetChangeIndicator bool                 `madatory`
	nextHopChainingCount  NextHopChainingCount `madatory`
	nas_Container         *[]byte              `optional`
}

func (ie *MasterKeyUpdate) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nas_Container != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBoolean(ie.keySetChangeIndicator); err != nil {
		return utils.WrapError("WriteBoolean keySetChangeIndicator", err)
	}
	if err = ie.nextHopChainingCount.Encode(w); err != nil {
		return utils.WrapError("Encode nextHopChainingCount", err)
	}
	if ie.nas_Container != nil {
		if err = w.WriteOctetString(*ie.nas_Container, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode nas_Container", err)
		}
	}
	return nil
}

func (ie *MasterKeyUpdate) Decode(r *uper.UperReader) error {
	var err error
	var nas_ContainerPresent bool
	if nas_ContainerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bool_keySetChangeIndicator bool
	if tmp_bool_keySetChangeIndicator, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean keySetChangeIndicator", err)
	}
	ie.keySetChangeIndicator = tmp_bool_keySetChangeIndicator
	if err = ie.nextHopChainingCount.Decode(r); err != nil {
		return utils.WrapError("Decode nextHopChainingCount", err)
	}
	if nas_ContainerPresent {
		var tmp_os_nas_Container []byte
		if tmp_os_nas_Container, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode nas_Container", err)
		}
		ie.nas_Container = &tmp_os_nas_Container
	}
	return nil
}
