package config

import "time"

const CertValidity = time.Hour * 24 * 365

// Lowering this will make previously created notes with
// names longer than this inaccessible
const NoteNameMaxLength uint64 = 512

// A note's content is read in one go form the filesystem (as to make it
// impossible for malicious clients to keep the filesystem locked
// almost indefinetely by lowering their download speed to the bare
// minimum), and then the read content is sent.
// Rising this to too high of a value may dramatically
// increase RAM usage.
const NoteContentsMaxLength uint64 = 1024 * 1024 * 2

// We cannot receive as many bytes as the current actual password
// since this would allow an attacker to deduce the password length
const PasswordMaxLength uint64 = 32

// If this is too big we allow an attacker to fill our RAM fake
// connections (even tho SWAP should help in that case) BUT it
// still might cost some extra little bit of CPU
const RecvActionDeadline = time.Minute * 5
