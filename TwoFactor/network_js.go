// +build js

package TwoFactor

import (
	"github.com/brokenbydefault/Nanollet/TwoFactor/Ephemeral"
	"net"
	"encoding/binary"
	"github.com/brokenbydefault/Nanollet/Wallet"
	"github.com/jaracil/goco/chrome/tcpsockets"
	"errors"
)

func NewRequesterServer(sk *Ephemeral.SecretKey, allowedDevices []Wallet.PublicKey) (req Request, ch <-chan Envelope, err error) {
	// NO-OP
	return req, nil, errors.New("not supported")
}

func ReplyRequest(device *Wallet.SecretKey, token Token, request Request) error {
	sk := Ephemeral.NewEphemeral()

	envelope := NewEnvelope(device.PublicKey(), sk.PublicKey(), request.Receiver, token)
	envelope.Sign(device)
	envelope.Encrypt(&sk)

	tcp, err := tcpsockets.Create()
	if err != nil {
		return err
	}

	if err := tcp.Connect(net.IP(request.IP[:]).String(), int(request.Port)); err != nil {
		return err
	}

	if err := binary.Write(tcp, binary.BigEndian, &envelope); err != nil {
		return err
	}

	return nil
}
