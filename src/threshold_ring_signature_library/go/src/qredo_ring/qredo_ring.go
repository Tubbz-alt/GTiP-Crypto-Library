//Copyright (c) 2018 Global Telecom Inter-Payments Alliance Ltd

package qredoring

// #include <stdio.h>
// #include <stdlib.h>
import "C"
// import "unsafe"

//------------------------------------------

//!
//! \brief The system parameters
//!
type parameters struct {
	// code length
	u uint;
	// code dimension
	k uint;
	// minimum distance
	d uint
	// error correcting ability of underpinning code
	w uint;
	// security parameter
	l uint;
	// number of participants
	numberOfParticipants uint;
	// threshold number of participants
	threshold uint;
}

//------------------------------------------

//!
//! \brief Creates the public key and private key
//!
//! \param parameters input  : the system parameters
//! \param public_key output : the public key created
//! \param secret_key output : the private key created
//!
// output public, private keyes
func keygen(p parameters) (string, string){
	return "public_key", "public_key"
}

//------------------------------------------

//!
//! \brief Takes a message creates signature out of it
//!
//! \param parameters  input  : the system parameters
//! \param message     input  : the message to be signed,
//! \param message_len input  : the the length of the message to be signed
//! \param secret_key  input  : the private key of the participant
//! \param signature   output : the result signature buffer
//!
func participantSign(p parameters, message string, secretKey string) string{
	return "signature"
}

//------------------------------------------

//!
//! \brief Takes an threshold-concatenated participant signatures and creates a ring leader signature
//!
//! \param parameters             input  : the system parameters
//! \param message                input  : the message to be signed,
//! \param participant_signatures input  : threshold-concatenated participants signatures
//! \param signature              output : the ring leader signature
//!
func leaderSign(p  parameters, message string , participantSignatures string) string{
	return "signature"
}

//------------------------------------------

//!
//! \brief Verifies that the message was correctly signed
//!
//! \param parameters        input  : the system parameters
//! \param message           input  : the buffer to be signed,
//! \param ring_signature    input  : the ring signature,
//! \param ring_public_key   input  : the ring public key
//! \param result            output : true if given message was verified, false otherwise
//!
func verify(p parameters, message string , ringSignature string, ringPublicKey string) bool{
	return true;
}

//------------------------------------------

