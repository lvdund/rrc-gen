package utils

import (
	"fmt"
	"github.com/lvdund/asn1go/uper"
)

type GenericSetuprelease interface {
	Encode(w *uper.UperWriter) error
	Decode(r *uper.UperReader) error
}

type SetupRelease[T GenericSetuprelease] struct {
	Release interface{}
	Setup   T
}

func (ie *SetupRelease[T]) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(2, 2, false); err != nil {
		return err
	}
	if err = ie.Setup.Encode(w); err != nil {
		return WrapError("Encode Setuprelease", err)
	}
	return nil
}

func (ie *SetupRelease[T]) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadChoice(2, false); err != nil {
		return WrapError("Decode choice Setuprelease", err)
	} else if v == 2 {
		var tmp T
		if err = tmp.Decode(r); err != nil {
			return WrapError("Decode Setuprelease", err)
		}
		ie.Setup = tmp
	} else if v != 2 {
		return fmt.Errorf("invalid choice: %d", v)
	}
	return nil
}
