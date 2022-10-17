package sifr

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"sync"

	"github.com/Indra-Labs/indra/pkg/schnorr"
)

type Nonce [aes.BlockSize]byte

// GetNonce reads from a cryptographically secure random number source
func GetNonce() (nonce *Nonce) {
	nonce = &Nonce{}
	if _, e := rand.Read(nonce[:]); log.E.Chk(e) {
	}
	return
}

// Dialog is a data structure for tracking keys used in a message exchange.
type Dialog struct {
	sync.Mutex
	// LastIn is the newest pubkey seen in a received message from the
	// correspondent.
	LastIn *schnorr.Pubkey
	// LastOut is the newest privkey used in an outbound message.
	LastOut *schnorr.Privkey
	// Seen are the keys that have been seen since the last new message sent
	// out to the correspondent.
	Seen []*schnorr.Pubkey
	// Used are the recently used keys that have not been invalidated by the
	// counterparty sending them in the Expires field.
	Used []*schnorr.Privkey
	// UsedFingerprints are 1:1 mapped to Used private keys for fast
	// recognition.
	UsedFingerprints []schnorr.Fingerprint
}

// NewDialog creates a new Dialog for tracking a conversation between two nodes.
// For the initiator, the pubkey is the current one advertised by the
// correspondent, and for a correspondent, this pubkey is from the first one
// appearing in the initial message.
func NewDialog(pub *schnorr.Pubkey) (d *Dialog) {
	d = &Dialog{LastIn: pub}
	return
}

// WireFrame is the data format that goes on the wire. This message is wrapped
// inside a WireMessage and the payload is also inside a WireMessage.
type WireFrame struct {
	// To is the fingerprint of the pubkey used in the ECDH key exchange.
	To *schnorr.Fingerprint
	// From is the pubkey corresponding to the private key used in the ECDH
	// key exchange.
	From *schnorr.PubkeyBytes
	// Expires are the fingerprints of public keys that the correspondent
	// can now discard as they will not be used again.
	Expires []schnorr.Fingerprint
	// Data is the payload of the message, which is wrapped in a
	// WireMessage.
	Data []byte
}

// WireMessage is simply a wrapper that provides tamper-proofing and
// authentication based on schnorr signatures.
type WireMessage struct {
	Payload   []byte
	Signature schnorr.SignatureBytes
}

// EncryptedMessage is a form of WireMessage that is encrypted. The cipher must
// be conveyed by other means, mainly being ECDH.
type EncryptedMessage struct {
	*Nonce
	Message []byte
}

func (em *EncryptedMessage) Serialize() (z []byte) {
	z = append((*em.Nonce)[:], em.Message...)
	return
}

func EncryptMessage(secret schnorr.Hash, message []byte,
	signingKey *schnorr.Privkey) (em *EncryptedMessage, e error) {

	if e = secret.Valid(); log.E.Chk(e) {
		return
	}
	var sig *schnorr.Signature
	if sig, e = signingKey.Sign(schnorr.SHA256(message)); log.E.Chk(e) {
		return
	}
	var block cipher.Block
	if block, e = aes.NewCipher(secret); log.E.Chk(e) {
		return
	}
	nonce := GetNonce()
	stream := cipher.NewCTR(block, nonce[:])
	msgLen := len(message) + schnorr.SigLen
	msg := make([]byte, msgLen)
	copy(msg, append(message, sig.Serialize()[:]...))
	stream.XORKeyStream(msg, msg)
	em = &EncryptedMessage{Nonce: nonce, Message: msg}
	return
}

func DecryptMessage(secret schnorr.Hash, message []byte,
	pub *schnorr.Pubkey) (cleartext []byte, e error) {

	if e = secret.Valid(); log.E.Chk(e) {
		return
	}
	nonce := &Nonce{}
	copy(nonce[:], message[:aes.BlockSize])
	msg := message[aes.BlockSize:]
	em := &EncryptedMessage{Nonce: nonce, Message: msg}
	var block cipher.Block
	if block, e = aes.NewCipher(secret); log.E.Chk(e) {
		return
	}
	stream := cipher.NewCTR(block, nonce[:])
	stream.XORKeyStream(em.Message, em.Message)
	sigStart := len(msg) - schnorr.SigLen
	m, s := em.Message[:sigStart], em.Message[sigStart:]
	var sig *schnorr.Signature
	if sig, e = schnorr.ParseSignature(s); log.E.Chk(e) {
		return
	}
	if !sig.Verify(schnorr.SHA256(m), pub) {
		e = errors.New("message signature verification failed")
		return
	}
	cleartext = m
	return
}